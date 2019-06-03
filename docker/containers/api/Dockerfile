FROM golang:1.12-alpine

ENV GO111MODULE=on

WORKDIR /go/src/app

RUN apk add --no-cache \
        alpine-sdk \
        git

COPY go.mod go.sum ./

RUN go mod download

RUN go get -u github.com/pilu/fresh

COPY . .

EXPOSE 8080

CMD ["fresh"]
