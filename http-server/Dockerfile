FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o main ./

FROM alpine:latest

WORKDIR /app
RUN apk add --no-cache \
    bash ca-certificates curl && \
    rm -rf /var/cache/apk/*

COPY --from=builder /app/main .
COPY ./config/* ./config/
COPY ./logs/* ./logs/

RUN mkdir -p logs

EXPOSE 8080

CMD ["./main"]
