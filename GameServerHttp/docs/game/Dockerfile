# 构建阶段
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache git ca-certificates tzdata

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 设置构建参数(可通过Jenkins传入)
ARG VERSION=dev
ARG BUILD_TIME=unknown
ARG COMMIT_SHA=unknown
ARG BRANCH=unknown

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.version=${VERSION} -X main.buildTime=${BUILD_TIME} -X main.branch=${BRANCH} -X main.commitId=${COMMIT_SHA}" -o GameService ./cmd/game

# 运行阶段
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 安装运行时依赖
RUN apk add --no-cache ca-certificates tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 创建非root用户
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN mkdir -p /app/conf/i18n && chown -R appuser:appgroup /app

# 从构建阶段复制二进制文件
COPY --from=builder --chown=appuser:appgroup /app/GameService /app/
COPY --chown=appuser:appgroup cmd/game/game.yaml /app/
COPY --chown=appuser:appgroup conf/i18n /app/conf/i18n

# 设置用户
USER appuser

# 暴露端口
EXPOSE 10010

# 健康检查
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD wget -qO- http://localhost:10010/health || exit 1

# 添加元数据标签
LABEL maintainer="ZSKJ Team"
LABEL app="slot-game-game-service"
LABEL description="Slot Game Game Service"

# 设置启动命令
ENTRYPOINT ["./GameService", "--config", "./game.yaml"]