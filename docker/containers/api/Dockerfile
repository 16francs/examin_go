FROM golang:1.12-alpine

ENV GO111MODULE=on

WORKDIR /go/src/app

RUN apk add --no-cache \
        alpine-sdk \
        git

RUN go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]
