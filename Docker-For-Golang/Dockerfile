FROM golang:1.22-alpine3.18 as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]

