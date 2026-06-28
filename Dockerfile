FROM golang:1.26.3 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM debian:stable-slim
WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]