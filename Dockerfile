FROM golang:1.19.1-alpine as builder
ENV GOPROXY=https://goproxy.cn
WORKDIR /build
COPY . .
RUN go mod tidy
RUN go build -o helloGin main.go

FROM  alpine:latest
RUN mkdir -p /cmd
WORKDIR  /cmd
COPY  --from=builder /build  .
EXPOSE 8888
CMD ["./helloGin"]