FROM golang:1.12
WORKDIR /go/src
COPY . .
CMD ["go", "run", "hello.go"]