package utils_middleware

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"

	// 替换opentracing为OpenTelemetry依赖
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// otelLogrusHook 将logrus日志与OpenTelemetry span关联
type otelLogrusHook struct {
	tracer  trace.Tracer
	span    trace.Span
	traceId string
	url     string
}

func (hook *otelLogrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *otelLogrusHook) Fire(entry *logrus.Entry) error {
	// 创建子span
	ctx := trace.ContextWithSpan(context.Background(), hook.span)
	_, childSpan := hook.tracer.Start(ctx, hook.url)
	defer childSpan.End()

	// 为日志添加trace-id
	entry.Data["trace-id"] = hook.traceId

	// 记录日志信息到span中
	attrs := []attribute.KeyValue{
		attribute.String("log.level", entry.Level.String()),
		attribute.String("log.time", entry.Time.Format(time.RFC3339)),
	}

	if entry.Caller != nil {
		attrs = append(attrs,
			attribute.String("log.file", fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)),
			attribute.String("log.function", entry.Caller.Function),
		)
	}

	attrs = append(attrs, attribute.String("log.message", entry.Message))

	childSpan.AddEvent("log", trace.WithAttributes(attrs...))

	return nil
}

type responseRecorder struct {
	gin.ResponseWriter
	statusCode      int
	body            *bytes.Buffer
	skipBodyCapture bool  // 添加标志指示是否跳过捕获响应体
	contentLength   int64 // 记录响应内容长度
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	if !r.skipBodyCapture {
		r.body.Write(b)
	} else {
		r.contentLength += int64(len(b))
	}
	return r.ResponseWriter.Write(b)
}

// TracerMiddleware 返回一个Gin中间件，使用OpenTelemetry记录请求
func TracerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := ""
		tracer := otel.Tracer("gin-http")
		propagator := otel.GetTextMapPropagator()
		ctx := propagator.Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

		// 创建span名称
		spanName := fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)
		ctx, span := tracer.Start(ctx, spanName,
			trace.WithSpanKind(trace.SpanKindServer),
		)

		// 设置span属性
		span.SetAttributes(
			attribute.String("http.method", c.Request.Method),
			attribute.String("http.url", c.Request.URL.String()),
			attribute.String("http.host", c.Request.Host),
			attribute.String("http.user_agent", c.Request.UserAgent()),
			attribute.Int64("http.request_size", c.Request.ContentLength),
			attribute.String("component", "gin-http"),
		)

		// 提取并设置trace-id（取消注释）
		spanContext := span.SpanContext()
		if spanContext.TraceID().IsValid() {
			traceId = spanContext.TraceID().String()
		} else {
			traceId = fmt.Sprintf("%s.%d", strings.ReplaceAll(uuid.New().String(), "-", ""), time.Now().UnixNano())
		}
		span.SetAttributes(attribute.String("trace-id", traceId))

		// 记录请求头
		headers := make(map[string]string)
		for k, v := range c.Request.Header {
			if len(v) > 0 {
				headers[k] = v[0]
			}
		}
		span.AddEvent("http.headers", trace.WithAttributes(attribute.String("headers", fmt.Sprintf("%v", headers))))

		// 记录URL参数
		queryParams := c.Request.URL.Query()
		if len(queryParams) > 0 {
			params := make(map[string]string)
			for k, v := range queryParams {
				if len(v) > 0 {
					params[k] = v[0]
				}
			}
			span.AddEvent("http.query_params", trace.WithAttributes(attribute.String("params", fmt.Sprintf("%v", params))))
		}

		// 将当前context注入响应头，以便客户端可以继续跟踪
		propagator.Inject(ctx, propagation.HeaderCarrier(c.Writer.Header()))

		// 添加logrus hook，使日志也能被span跟踪
		logrus.AddHook(&otelLogrusHook{tracer: tracer, span: span, traceId: traceId, url: c.Request.URL.Path})

		// 读取并记录请求体
		if c.Request.Body != nil && c.Request.Method != http.MethodOptions {
			contentType := c.GetHeader("Content-Type")
			if strings.Contains(contentType, "multipart/form-data") {
				span.AddEvent("http.request_body", trace.WithAttributes(
					attribute.String("content_type", contentType),
					attribute.Int64("size", c.Request.ContentLength),
				))
				logrus.Debugf("Request URL: %s, Method: %s, BodyType: multipart/form-data, Size: %d",
					c.Request.URL, c.Request.Method, c.Request.ContentLength)
			} else {
				body, err := io.ReadAll(c.Request.Body)
				if err == nil {
					c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
					if strings.Contains(contentType, "application/json") {
						var jsonData map[string]any
						if err := jsoniter.Unmarshal(body, &jsonData); err == nil {
							span.AddEvent("http.request_body", trace.WithAttributes(
								attribute.String("body", fmt.Sprintf("%v", jsonData)),
							))
						} else {
							span.AddEvent("http.request_body", trace.WithAttributes(
								attribute.String("body", string(body)),
							))
						}
						logrus.Debugf("Request URL: %s, Method: %s, Body: %s",
							c.Request.URL, c.Request.Method, string(body))
					} else {
						span.AddEvent("http.request_body", trace.WithAttributes(
							attribute.String("content_type", contentType),
							attribute.String("body", string(body)),
						))
						logrus.Debugf("Request URL: %s, Method: %s, BodyLen: %d",
							c.Request.URL, c.Request.Method, len(body))
					}
				}
			}
		}

		// 包装ResponseWriter以捕获响应
		response := &responseRecorder{
			ResponseWriter:  c.Writer,
			body:            bytes.NewBufferString(""),
			statusCode:      http.StatusOK,
			skipBodyCapture: false,
			contentLength:   0,
		}

		// 检查是否是文件下载请求路径
		if strings.Contains(c.Request.URL.Path, "/files/download") ||
			strings.Contains(c.Request.URL.Path, "/files/preview") {
			response.skipBodyCapture = true
		}

		c.Writer = response

		// 设置请求上下文
		c.Request = c.Request.WithContext(ctx)

		// 处理请求
		c.Next()

		// 记录状态码
		statusCode := response.statusCode
		span.SetAttributes(attribute.Int("http.status_code", statusCode))

		// 处理错误情况
		if statusCode >= 400 {
			span.SetStatus(codes.Error, fmt.Sprintf("HTTP status code: %d", statusCode))
			span.SetAttributes(attribute.String("error.type", fmt.Sprintf("http_status_%d", statusCode)))
		} else {
			span.SetStatus(codes.Ok, "")
		}

		// 记录响应体
		if c.Request.Method != http.MethodOptions {
			contentType := response.Header().Get("Content-Type")

			if response.skipBodyCapture {
				contentDisposition := response.Header().Get("Content-Disposition")
				span.AddEvent("http.file_download", trace.WithAttributes(
					attribute.String("content_type", contentType),
					attribute.String("content_disposition", contentDisposition),
					attribute.Int64("size", response.contentLength),
				))
				logrus.Debugf(
					"File Response URL: %s, Method: %s, Status: %d, ContentType: %s, Size: %d",
					c.Request.URL, c.Request.Method, statusCode, contentType, response.contentLength,
				)
			} else {
				responseBody := response.body.Bytes()
				if strings.Contains(contentType, "application/json") {
					var jsonData map[string]any
					if err := jsoniter.Unmarshal(responseBody, &jsonData); err == nil {
						span.AddEvent("http.response_body", trace.WithAttributes(
							attribute.String("body", fmt.Sprintf("%v", jsonData)),
						))
					} else {
						span.AddEvent("http.response_body", trace.WithAttributes(
							attribute.String("body", string(responseBody)),
						))
					}
					logrus.Debugf("Response URL: %s, Method: %s, Status: %d, Body: %s",
						c.Request.URL, c.Request.Method, statusCode, string(responseBody))
				} else {
					span.AddEvent("http.response_body", trace.WithAttributes(
						attribute.String("content_type", contentType),
						attribute.Int("size", len(responseBody)),
					))
					logrus.Debugf("Response URL: %s, Method: %s, Status: %d, BodyLen: %d",
						c.Request.URL, c.Request.Method, statusCode, len(responseBody))
				}
			}
		}

		// 记录错误信息
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				span.AddEvent("error", trace.WithAttributes(
					attribute.String("error.message", err.Error()),
				))
			}
		}

		// 结束span
		span.End()

		// 移除logrus钩子
		logrus.StandardLogger().ReplaceHooks(make(logrus.LevelHooks))
	}
}
