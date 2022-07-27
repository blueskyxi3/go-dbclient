FROM golang:alpine AS builder
LABEL author=vincentzou \
      version=0.0.1
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 app
# RUN go build -o app .

# Build
RUN go mod download && \
	go build -o mysqlctl mysql/main.go  && \
	go build -o oraclectl oracle/main.go




###################
# 接下来创建一个小镜像
###################
FROM alpine as mysqlctl
# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/mysqlctl /app
# 需要运行的命令
ENTRYPOINT ["/app"]

FROM alpine as oraclectl
# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/oraclectl /app
# 需要运行的命令
ENTRYPOINT ["/app"]