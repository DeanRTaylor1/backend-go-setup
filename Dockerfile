# Build stage
FROM golang:1.20.6-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY .env.development.local .
COPY .env.test.local .

EXPOSE 8080
CMD [ "/app/main" ]