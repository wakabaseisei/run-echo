FROM golang:1.17.2-alpine3.14

COPY ./src/go.mod /go/src/go.mod
COPY ./src/go.sum /go/src/go.sum

RUN apk update \
  && apk add --no-cache git

WORKDIR /go/src
RUN go mod download

COPY src/ /go/src/