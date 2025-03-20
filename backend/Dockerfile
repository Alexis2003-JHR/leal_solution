FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY . .

WORKDIR /app/cmd/

RUN go mod tidy
RUN go build -o /app/leal

# Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/leal /app/leal

RUN chmod +x /app/leal

CMD ["./leal"]