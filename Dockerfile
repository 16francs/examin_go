# マルチビルド
# 開発環境
FROM golang:1.12-alpine as build

WORKDIR /go/src/app

ENV GO111MODULE=on

COPY . .

RUN apk add --no-cache git
RUN go build -o app
# RUN go get github.com/oxequa/realize

# 本番環境
