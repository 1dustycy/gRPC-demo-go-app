# Build image
FROM golang:1.14-alpine AS builder

ENV GOFLAGS="-mod=readonly"

RUN apk add --update --no-cache ca-certificates make git

RUN mkdir -p /app
WORKDIR /app

ARG GOPROXY=https://goproxy.cn,direct

COPY ./ /app

RUN go mod download

RUN set -xe && make build-release


# Final image
FROM alpine:3.11

RUN apk add --update --no-cache ca-certificates tzdata bash

SHELL ["/bin/bash", "-c"]

COPY --from=builder /app/release/* /usr/local/bin/

EXPOSE 8080
CMD serve
