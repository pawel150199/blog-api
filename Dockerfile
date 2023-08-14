FROM golang:1.19
LABEL maintainer="Pawel Polski"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download