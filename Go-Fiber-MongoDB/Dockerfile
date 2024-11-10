# Go uygulaması için yapılandırma
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o go-fiber-hrms cmd/hrms/main.go

# Uygulamanın çalıştırılması için yapılandırma
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/go-fiber-hrms .
COPY .env .
CMD ["./go-fiber-hrms"]
