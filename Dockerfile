# Build stage
FROM golang:1.21.2-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN go build \
    -ldflags="-X 'git.logntnu.no/tekkom/web/beehive/admin-api.version=${IMAGE_VERSION}'" \
    -o bin/main \
    cmd/main.go

# Run stage
FROM scratch

WORKDIR /app

COPY --from=builder /app/bin/main .

CMD ["./main"]
