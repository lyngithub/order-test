FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /build
COPY . .
RUN go build -o app .
FROM scratch
COPY --from=builder /build/app /
COPY --from=builder /build/config-debug.yaml /
ENTRYPOINT ["/app"]