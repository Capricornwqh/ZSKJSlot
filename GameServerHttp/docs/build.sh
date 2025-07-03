#!/bin/bash

# 设置版本信息
VERSION="v1.0.0"
BUILD_TIME=$(date +%Y-%m-%dT%H:%M:%S)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
COMMIT_ID=$(git rev-parse --short=10 HEAD)

# 创建输出目录
mkdir -p ../bin

# 为Linux平台构建
GOOS=linux GOARCH=amd64 go build -a -ldflags "-extldflags '-static' -X 'main.version=$VERSION' -X 'main.buildTime=$BUILD_TIME' -X 'main.branch=$BRANCH' -X 'main.commitId=$COMMIT_ID'" -o ../bin/file_manager_linux ../cmd

# 为Darwin平台构建
GOOS=darwin GOARCH=amd64 go build -a -ldflags "-extldflags '-static' -X 'main.version=$VERSION' -X 'main.buildTime=$BUILD_TIME' -X 'main.branch=$BRANCH' -X 'main.commitId=$COMMIT_ID'" -o ../bin/file_manager_darwin ../cmd

echo "构建完成: Linux 和 Darwin 平台的可执行文件已生成在 ../bin 目录中。"