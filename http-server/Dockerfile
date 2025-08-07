FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o main ./

FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache bash ca-certificates curl && \
    mkdir -p /app/logs && \
    chown -R 1001:1001 /app && \
    chmod -R 755 /app/logs/*

COPY --from=builder /app/main .
COPY --chown=1001:1001 ./config ./config

USER 1001
EXPOSE 8080

CMD ["./main"]
