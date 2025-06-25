# -------- Build stage ---------
FROM golang:1.24-alpine AS builder
LABEL authors="emmanuel-kpendo"
LABEL lab="week4_lab1"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main .

# ------- Run stage -------
FROM alpine:latest

WORKDIR /root/

# Set Gin to release mode
ENV GIN_MODE=release

COPY --from=builder /app/main .

EXPOSE 5000

ENTRYPOINT ["./main"]
