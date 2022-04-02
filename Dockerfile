## Build

FROM golang:1.18-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GO111MODULE=on \
    GOARCH=amd64 \
    GOPATH=/go

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build

## Deploy

FROM alpine

WORKDIR /app
COPY --from=builder /app ./
RUN chmod +x /app/go-rest

ENTRYPOINT ["/app/go-rest"]
