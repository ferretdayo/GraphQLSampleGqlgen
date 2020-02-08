FROM golang:1.12.5-alpine

RUN apk update \
    && apk add --no-cache git \
    && GO111MODULE=off go get -u github.com/codegangsta/gin

ENV GO111MODULE=on

WORKDIR /go/src/github.com/GraphQLSample

COPY . /go/src/github.com/GraphQLSample

RUN go build

CMD gin -i run