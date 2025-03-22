# Etapa de construcción
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init

RUN go mod tidy
RUN go build -o /app/leal

# Etapa final
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/leal /app/leal
COPY --from=builder /app/docs /app/docs

RUN chmod +x /app/leal

EXPOSE 6060

CMD ["./leal"]
