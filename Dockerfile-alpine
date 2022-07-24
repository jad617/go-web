FROM golang:alpine AS builder
# FROM golang:1.18-bullseye AS builder

# RUN apt-get update && apt-get install -y ca-certificates

RUN mkdir -p /build

WORKDIR /build

COPY . .

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN go build -o /build/go-web .

FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY --from=builder /build/go-web .

EXPOSE 8080

CMD ["/app/go-web"]
