FROM golang:1.19
LABEL maintainer="Pawel Polski"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./ 
COPY .env ./
CMD ["go", "run", "main.go"]

