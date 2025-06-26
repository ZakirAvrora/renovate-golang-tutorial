FROM golang:1.23.3

WORKDIR /go/src

COPY go .
RUN GOOS=linux GOARCH=amd64 go build -o /go/main main.go
ENTRYPOINT ["/go/main"]
