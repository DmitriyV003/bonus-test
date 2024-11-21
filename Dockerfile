FROM golang:1.23-alpine AS builder

RUN apk add --no-cache build-base sqlite sqlite-dev git bash

RUN go install github.com/gobuffalo/pop/soda@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/main.go

FROM alpine:3.18

RUN apk add --no-cache git bash

COPY --from=builder /app/server /app/server

COPY --from=builder /go/bin/soda /usr/local/bin/soda

COPY migrations /app/migrations
COPY database.yml /app/database.yml

WORKDIR /app

EXPOSE 8080

COPY docker-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

ENTRYPOINT ["docker-entrypoint.sh"]
