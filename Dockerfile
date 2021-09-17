FROM golang:1.16.4-alpine AS builder

WORKDIR /build
ENV CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./main.go


FROM alpine:3.13.6

RUN apk add ca-certificates
WORKDIR /app

COPY --from=builder /build/main ./