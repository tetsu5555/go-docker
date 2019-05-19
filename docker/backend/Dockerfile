FROM golang:1.12
ENV GO111MODULE=on
WORKDIR /go/src

COPY backend/go.mod .
COPY backend/go.sum .

RUN go mod download

COPY . .
# COPY . .
EXPOSE 1323
CMD ["go", "run", "main.go"]