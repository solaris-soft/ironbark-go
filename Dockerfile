FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o main cmd/http/main.go

FROM alpine:latest

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["./main"]