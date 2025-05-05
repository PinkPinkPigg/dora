# 使用官方 Go 镜像作为构建阶段
FROM golang:1.23.8 AS builder

# 设置工作目录
WORKDIR /workspace

ENV GOPROXY=https://goproxy.cn,direct

# 将整个工作区复制到容器中
COPY /projects/dora/code/dora .

# 切换到服务 A 的目录
WORKDIR /workspace/services/executor

# 使用工作区模式构建服务 A
RUN go mod tidy && go build -o executor_service

# 使用更小的基础镜像运行程序
FROM debian:bullseye-slim

# 设置工作目录
WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /workspace/services/executor/executor_service .

# 暴露需要的端口（根据服务需求）
EXPOSE 8080

# 运行服务
CMD ["./executor_service"]