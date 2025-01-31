FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o payment_service main.go
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/payment_service .
COPY config.yaml .
EXPOSE 8080
CMD ["./payment_service", "-config", "config.yaml"]
