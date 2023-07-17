FROM golang:1.20.4
WORKDIR /app
COPY . /app

# remove the following two line if you are not in China
RUN export GO111MODULE=on
RUN export GOPROXY=https://goproxy.cn

RUN go mod download
RUN go build -o main

# Redis port, see setting/config.yaml
EXPOSE 6379

# MySQL port, see setting/config.yaml
EXPOSE 3306

CMD ["./main"]
