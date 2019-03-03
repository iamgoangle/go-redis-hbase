FROM golang:1.11.1-alpine3.8

ENV CGO_ENABLED=0

# install git, curl
RUN apk add --no-cache git curl
