FROM golang:1.24-alpine AS builder
WORKDIR /app

ARG APP_UID
ARG APP_GID

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
    chmod -R 755 /app/logs

COPY --from=builder /app/main .
COPY --chown=1001:1001 ./config ./config

USER ${APP_UID}:${APP_GID}
EXPOSE 8080

CMD ["./main"]
