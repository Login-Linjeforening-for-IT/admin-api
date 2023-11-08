# Build stage
FROM golang:1.21.2-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN go build -o bin/main cmd/main.go

# Run stage
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/bin/main .

CMD ["./main"]
