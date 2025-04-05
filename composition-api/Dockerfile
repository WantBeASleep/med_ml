FROM golang:alpine

WORKDIR /service
COPY . .

RUN go build -o bin/service cmd/service/main.go

EXPOSE 50055

CMD ["sh", "-c", "./bin/service"]
