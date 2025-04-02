# 使用官方的 Alpine Linux 镜像作为运行环境
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY . .


# 暴露应用程序端口（如果需要）
# EXPOSE 8080

# 运行应用程序
CMD ["./go-i18n-grpc"]