ARG CORE_VERSION=v1.18.0

FROM golang:1.26-alpine as build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN  go build  ./cmd/firebtc

#######

FROM ghcr.io/streamingfast/firehose-core:$CORE_VERSION as core

COPY --from=build /app/firebtc /app/firebtc