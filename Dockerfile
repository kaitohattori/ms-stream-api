FROM golang:1.16.8-alpine3.14 AS build

WORKDIR /go/src/ms-stream-api

COPY . .

ENV GO111MODULE=on

RUN go mod download
RUN go get -u github.com/swaggo/swag/cmd/swag

RUN swag init
RUN go build main.go

ENTRYPOINT ./main
