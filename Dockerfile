FROM alpine:latest

# 安装 curl
RUN apk add --no-cache curl

# 下载二进制
RUN curl -L -o /usr/local/bin/fuck-u-code https://github.com/Done-0/fuck-u-code/releases/latest/download/fuck-u-code_linux_amd64 \
    && chmod +x /usr/local/bin/fuck-u-code

# 容器启动默认执行
ENTRYPOINT ["/usr/local/bin/fuck-u-code"]
