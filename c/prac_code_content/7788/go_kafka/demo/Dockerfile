FROM golang:1.18  as builder

WORKDIR /build
COPY go.mod .
COPY go.sum .
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \ 
    GOPROXY="https://goproxy.cn,direct"

RUN go mod download
COPY ./c/prac_code_content ./c/prac_code_content
RUN go build -o web_demo ./c/prac_code_content/7788/go_kafka/demo

FROM ccr.ccs.tencentyun.com/back_end/ubuntu:20.04 as publish
WORKDIR /bin
COPY --from=builder /build/web_demo .
ENTRYPOINT [ "./web_demo" ]

