FROM golang:1.14-alpine AS build

WORKDIR /src

COPY go.mod /src/
COPY main.go /src/

RUN go mod download

CMD go run main.go
