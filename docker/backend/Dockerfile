FROM golang:1.12
ENV GO111MODULE=on

RUN apt-get update -qq && \
    apt-get install -y \
    vim \
    git \
    && go get -u github.com/pilu/fresh \
    && go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

WORKDIR /go/src/backend

COPY backend/go.mod .
COPY backend/go.sum .

RUN go mod download

COPY . .

CMD ["fresh"]