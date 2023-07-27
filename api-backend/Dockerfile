FROM golang:1.20-alpine AS builder 

WORKDIR /go/src/app

COPY . .

RUN GOOS=linux GOARCH=amd64 GOARM=7 CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o main cmd/*.go