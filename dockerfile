# syntax=docker/dockerfile:1

# ---- Builder stage ----
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Needed for some Go modules
# RUN apk add --no-cache git

# Copy go.mod files first (cache optimization)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary (main.go is in root)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o main

# ---- Final image ----
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
