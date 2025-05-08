# Dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
ENV APP_PORT 8080:wq

CMD ["./server"]
