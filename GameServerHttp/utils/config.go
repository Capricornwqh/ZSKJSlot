package utils

import (
	"bytes"
	"context"
	"io"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

const (
	MD_TRACE_ID   = "trace-id"
	MD_USER_ID    = "user-id"
	MD_CLIENT_IP  = "x-forwarded-for"
	MD_GRPC_AGENT = "user-agent"

	MD_AGENT_APP = "slot-app"
	MD_AGENT_WEB = "slot-web"
)

var Conf *Config

type Config struct {
	Version     string
	BuildTime   string
	Branch      string
	CommitId    string
	Game        Game         `yaml:"game"`
	GeoDB       string       `yaml:"geodb"`
	Email       Email        `yaml:"email"`
	I18N        I18N         `yaml:"i18n"`
	PostgreSQL  PostgreSQL   `yaml:"postgresql"`
	Redis       Redis        `yaml:"redis"`
	Tracing     Tracing      `yaml:"tracing"`
	Logging     []LogrusHook `yaml:"logging"`
	Environment string       `yaml:"environment"`
	Server      Server       `yaml:"server"`
}

type I18N struct {
	Default string `yaml:"default"`
	Dir     string `yaml:"dir"`
}

type Game struct {
	MTCount int     `yaml:"mtCount"`
	MCCount uint64  `yaml:"mcCount"`
	MCPrec  float64 `yaml:"mcPrec"`
	DefMRTP float64 `yaml:"defMRTP"`
}

// 服务器配置
type Server struct {
	Type   string `yaml:"type"`
	Domain string `yaml:"domain"`
}

// 邮箱配置
type Email struct {
	Addr     string `yaml:"addr"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
	Port     int    `yaml:"port"`
}

// PostgreSQL配置
type PostgreSQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"sslmode"`
}

// Redis配置
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	Timeout  int    `yaml:"timeout"`
}

// 链路追踪配置
type Tracing struct {
	Enabled bool       `yaml:"enabled"`
	OTLP    OTLPConfig `yaml:"otlp"`
}

// OpenTelemetry OTLP配置
type OTLPConfig struct {
	ServiceName   string            `yaml:"serviceName"`   // 服务名称
	Endpoint      string            `yaml:"endpoint"`      // OTLP端点，如 "localhost:4317"
	Insecure      bool              `yaml:"insecure"`      // 是否使用不安全连接
	Headers       map[string]string `yaml:"headers"`       // 自定义头信息
	SamplingRatio float64           `yaml:"samplingRatio"` // 采样比例，1.0表示全采样
	Timeout       int               `yaml:"timeout"`       // 连接超时(秒)
	Disabled      bool              `yaml:"disabled"`      // 是否禁用
}

// 日志配置
type Log struct {
	Release bool   `yaml:"release"`
	Path    string `yaml:"path"`
	Port    int64  `yaml:"port"`
}

type LogrusHook struct {
	// The type of hook, currently only "file" is supported.
	Type string `yaml:"type"`

	// The level of the logs to produce. Will output only this level and above.
	Level string `yaml:"level"`

	// The parameters for this hook.
	Params map[string]any `yaml:"params"`
}

// 初始化
func SetupConfig(path string) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		logrus.Fatal(err)
	}
}

var PathIgnore []byte

var ReservedUsernames []byte

const (
	DefaultConfigFileName = "server.yaml"
)

var (
	ConfigFileDir  = "/conf/"
	UploadFilePath = "/uploads/"
	CacheDir       = "/cache/"
)

// jaeger链路追踪 (使用OpenTelemetry OTLP)
func SetupTracing() (io.Closer, error) {
	if !Conf.Tracing.Enabled || len(Conf.Tracing.OTLP.ServiceName) <= 0 {
		return io.NopCloser(bytes.NewReader([]byte{})), nil
	}

	ctx := context.Background()

	// 设置连接选项
	var opts []otlptracegrpc.Option
	opts = append(opts, otlptracegrpc.WithEndpoint(Conf.Tracing.OTLP.Endpoint))
	if Conf.Tracing.OTLP.Insecure {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}
	if len(Conf.Tracing.OTLP.Headers) > 0 {
		opts = append(opts, otlptracegrpc.WithHeaders(Conf.Tracing.OTLP.Headers))
	}

	// 设置连接超时
	if Conf.Tracing.OTLP.Timeout > 0 {
		opts = append(opts, otlptracegrpc.WithTimeout(time.Duration(Conf.Tracing.OTLP.Timeout)*time.Second))
	}
	exporter, err := otlptracegrpc.New(ctx, opts...)
	if err != nil {
		logrus.Fatalf("创建OTLP导出器失败: %v", err)
	}
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(Conf.Tracing.OTLP.ServiceName),
			semconv.ServiceVersion(Conf.Version),
		),
	)

	if err != nil {
		logrus.Fatalf("创建资源失败: %v", err)
	}

	// 配置采样率
	samplingRatio := Conf.Tracing.OTLP.SamplingRatio
	if samplingRatio <= 0 {
		samplingRatio = 1.0
	}

	// 创建并配置TracerProvider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(samplingRatio)),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))
	closer := &tracerProviderCloser{tp: tp}

	return closer, nil
}

// 实现io.Closer接口的包装器，用于优雅关闭TracerProvider
type tracerProviderCloser struct {
	tp *sdktrace.TracerProvider
}

func (t *tracerProviderCloser) Close() error {
	// 使用5秒超时关闭TracerProvider，确保数据发送完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return t.tp.Shutdown(ctx)
}
