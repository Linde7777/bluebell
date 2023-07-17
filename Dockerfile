FROM golang:1.20.4 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./

RUN export GO111MODULE=on

# TODO: remove this if you are not in China
RUN export GOPROXY=https://goproxy.cn

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /bluebell

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /bluebell /bluebell

# Redis port, see setting/config.yaml
EXPOSE 6379

# MySQL port, see setting/config.yaml
EXPOSE 3306

USER nonroot:nonroot

CMD ["/bluebell"]
