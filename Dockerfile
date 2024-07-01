########-- Build stage --########

FROM golang:1.22-alpine AS builder
LABEL authors="MarkRaiter"

WORKDIR /opt

COPY . /opt

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o ./runner ./cmd

########--- Deploy stage ---########

FROM alpine:3.18
LABEL authors="MarkRaiter"

WORKDIR /opt 

COPY --from=builder /opt/runner /opt/
COPY .env /opt/.env

EXPOSE 8000

ENTRYPOINT /opt/runner