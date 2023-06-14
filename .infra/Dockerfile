FROM golang:alpine AS builder

RUN mkdir -p /go/src/github.com/DmitryLogunov/golang-graphql
WORKDIR /go/src/github.com/DmitryLogunov/golang-graphql
COPY . /go/src/github.com/DmitryLogunov/golang-graphql

RUN apk --no-cache add git

RUN go build -o /go/src/github.com/DmitryLogunov/golang-graphql/golang-graphql

FROM alpine

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY --from=builder /go/src/github.com/DmitryLogunov/golang-graphql/golang-graphql .

CMD ["./golang-graphql"]