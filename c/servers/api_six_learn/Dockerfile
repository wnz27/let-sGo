FROM golang:1.17 as builder

WORKDIR /build
COPY ./c/servers/api_six_learn .
COPY go.mod .

RUN go env -w CGO_ENABLED="0"

ENV http_proxy=socks5://172.2.0.230:7891
ENV https_proxy=socks5://172.2.0.230:7891
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN export GIN_MODE=release

# RUN GOOS=darwin GOARCH=arm64 go build
RUN GOARCH=arm64 go build

FROM alpine:latest as publish
WORKDIR /opt/www/api-six-test
COPY --from=builder /build/api_six_learn .

EXPOSE 8888

ENTRYPOINT ["./api_six_learn"]