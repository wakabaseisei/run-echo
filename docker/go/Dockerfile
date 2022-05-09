# 開発用イメージ
FROM golang:1.17.2-alpine3.14 as dev
ENV CGO_ENABLED=0

COPY ./src/go.mod /go/src/go.mod
COPY ./src/go.sum /go/src/go.sum

RUN apk update \
  && apk add --no-cache git

WORKDIR /go/src
RUN go mod download

RUN go install -v github.com/ramya-rao-a/go-outline@latest
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install -v golang.org/x/tools/gopls@latest

COPY src/ /go/src/

# ビルド用イメージ
FROM golang:1.17.2-alpine3.14 as builder

ENV CGO_ENABLED=0

COPY src/go.mod /go/src/go.mod
COPY src/go.sum /go/src/go.sum

RUN apk update \
  && apk add --no-cache git

WORKDIR /go/src
RUN go mod download

COPY src/ /go/src/
RUN GOOS=linux go build -o ./binary

# 本番用イメージ
FROM scratch as prod

COPY --from=builder /go/src/binary /go/src/

EXPOSE 8080
CMD ["/go/src/binary"]