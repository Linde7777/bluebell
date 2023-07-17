FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o bluebell_app .

FROM debian:stretch-slim

COPY ./settings /settings

COPY --from=builder /build/bluebell_app /

EXPOSE 9091
EXPOSE 3306
EXPOSE 6379

ENTRYPOINT ["/bluebell_app", "conf/config.yaml"]
