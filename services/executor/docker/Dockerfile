# 使用官方 Go 镜像作为构建阶段
FROM golang:1.23.8 AS builder


ENV GOPROXY=https://goproxy.cn,direct


RUN apt-get update && apt-get install -y wget build-essential


# 设置工作目录
WORKDIR /workspace

# 将整个工作区复制到容器中
COPY /projects/dora/code/dora .

# 切换到服务 A 的目录
WORKDIR /workspace/services/executor

# 使用工作区模式构建服务 A
RUN go mod tidy && go build -o executor_service

# 使用更小的基础镜像运行程序
FROM ubuntu:22.04
# 设置工作目录
WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /workspace/services/executor/executor_service .
# 安装python3环境
RUN apt-get update && apt-get install -y python3 && apt-get install -y python3-pip
#验证版本
RUN python3 --version
RUN ldd --version

# 暴露需要的端口（根据服务需求）
EXPOSE 50052

# 运行服务
CMD ["./executor_service"]
