FROM golang:1.20.4
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o main
EXPOSE 6379 # Redis port, see setting/config.yaml
EXPOSE 3306 # MySQL port, see setting/config.yaml
CMD ["./main"]
