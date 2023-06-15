FROM golang:alpine AS builder

RUN mkdir -p /go/src/github.com/DmitryLogunov/trading-platform
WORKDIR /go/src/github.com/DmitryLogunov/trading-platform
COPY . /go/src/github.com/DmitryLogunov/trading-platform

RUN apk --no-cache add git

RUN go build -o /go/src/github.com/DmitryLogunov/trading-platform/trading-platform-api app/server.go

FROM alpine

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY --from=builder /go/src/github.com/DmitryLogunov/trading-platform/trading-platform-api .

CMD ["./trading-platform"]