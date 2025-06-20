FROM golang:1.24-alpine
LABEL authors="emmanuel-kpendo"

WORKDIR /app

COPY . .

EXPOSE 5000

RUN go mod tidy