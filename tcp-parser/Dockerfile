FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache bash


COPY --from=builder /app/main .

CMD [ "./main" ]
