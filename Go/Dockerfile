FROM golang:1.8 as goimage
ENV SRC=/go/src/
RUN mkdir -p /go/src/
WORKDIR /go/src/go_docker
RUN  go build -o bin/go_docker