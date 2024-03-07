ARG CORE_VERSION=v1.2.4

FROM golang:1.21-alpine as build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN  go build  ./cmd/firebtc

#######

FROM ghcr.io/streamingfast/firehose-core:$CORE_VERSION as core

COPY --from=build /app/firebtc /app/firebtc