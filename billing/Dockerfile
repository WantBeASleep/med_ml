FROM golang:alpine

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /service
COPY . .
RUN go build -o bin/service cmd/service/main.go

EXPOSE 50055

CMD ["sh", "-c", "goose -dir /service/db/migrations up && ./bin/service"]
