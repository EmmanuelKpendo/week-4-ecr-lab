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

# Copy built binary
COPY --from=builder /app/main .

# ✅ Copy templates folder for LoadHTMLGlob("templates/*") to work
COPY --from=builder /app/templates ./templates

# ✅ Optional: copy static files if using ServeStatic or similar
COPY --from=builder /app/static ./static

EXPOSE 5000

ENTRYPOINT ["./main"]
