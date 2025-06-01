FROM golang:alpine AS builder

WORKDIR /service
COPY . .

# Install goose for migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go build -o bin/chat_service cmd/chat_service/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /service

# Copy binary and config
COPY --from=builder /service/bin/chat_service .
COPY --from=builder /service/config/config.yaml ./config/
COPY --from=builder /service/migrations ./migrations/
COPY --from=builder /go/bin/goose /usr/local/bin/

EXPOSE 8080 9090

# Run migrations and start service
CMD ["sh", "-c", "goose -dir /service/migrations postgres \"$POSTGRES_DSN\" up && ./chat_service"] 