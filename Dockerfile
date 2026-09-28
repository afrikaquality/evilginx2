# ---------- Build stage ----------
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git make gcc musl-dev

WORKDIR /src
COPY . .

RUN go build -o /usr/local/bin/evilginx .

# ---------- Runtime stage ----------
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata bash

# Create the app user (avoid running as root when possible)
RUN addgroup -S evilginx && adduser -S evilginx -G evilginx

COPY --from=builder /usr/local/bin/evilginx /usr/local/bin/evilginx
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Evilginx writes phishlets/, sessions/, certs/ relative to WORKDIR
WORKDIR /app
RUN mkdir -p /app/phishlets /app/sessions /app/certs && \
    chown -R evilginx:evilginx /app

USER evilginx

EXPOSE 53/udp 80/tcp 443/tcp

ENTRYPOINT ["/entrypoint.sh"]
