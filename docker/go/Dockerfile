FROM golang:1.17.2-alpine3.14

COPY ./src/api/go.mod /go/src/api/go.mod
COPY ./src/api/go.sum /go/src/api/go.sum

RUN apk update \
  && apk add --no-cache git

WORKDIR /go/src/api
RUN go mod download

COPY src/api/ /go/src/api/