FROM golang:alpine

WORKDIR /mallus_api

ADD . .

RUN go mod download

ENTRYPOINT go build server.go && ./server