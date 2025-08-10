FROM golang:1.24-alpine AS builder
WORKDIR /app

ARG APP_UID
ARG APP_GID

ENV APP_UID=${APP_UID}
ENV APP_GID=${APP_GID}

COPY go.mod ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest
WORKDIR /app

ARG APP_UID
ARG APP_GID

ENV APP_UID=${APP_UID}
ENV APP_GID=${APP_GID}

RUN apk add --no-cache bash ca-certificates curl && \
    addgroup -g ${APP_GID} testusers && \
    adduser -u ${APP_UID} -G testusers -s /bin/sh -D user-${APP_UID} && \
    mkdir -p /app/logs && \
    chown -R ${APP_UID}:${APP_GID} /app && \
    chmod -R 755 /app/logs

COPY --from=builder /app/main .
COPY --chown=${APP_UID}:${APP_GID} ./config ./config

EXPOSE 8080

CMD ["./main"]
