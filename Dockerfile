# Base
FROM golang:1.22.1-alpine AS builder
RUN apk add --no-cache git build-base
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o ./cmd/reveal ./cmd/reveal

# Release
FROM alpine:3.18.3
RUN apk -U upgrade --no-cache \
    && apk add --no-cache bind-tools ca-certificates
COPY --from=builder /app/cmd/reveal/reveal /usr/local/bin/

ENTRYPOINT ["reveal"]