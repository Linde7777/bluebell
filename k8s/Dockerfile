FROM golang:1.20.4
WORKDIR /app
COPY . /app

RUN export GO111MODULE=on

# remove this if you are not in China
RUN export GOPROXY=https://goproxy.cn

RUN go mod tidy
RUN go build -o main

# Redis port, see setting/config.yaml
EXPOSE 6379

# MySQL port, see setting/config.yaml
EXPOSE 3306

CMD ["./main"]
