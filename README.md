<p align="center"><img src="img/cover.png" alt="StealthPass" width="100%"></p>

# StealthPass v2.0.0 🚀

<p align="center">
  <a href="go.mod"><img alt="Go version" src="https://img.shields.io/github/go-mod/go-version/StealthPassTeam/StealthPass?logo=go&label=Go"></a>
  <a href="https://github.com/StealthPassTeam/StealthPass/releases/latest"><img alt="Latest release" src="https://img.shields.io/github/v/release/StealthPassTeam/StealthPass?logo=github&label=release&color=orange"></a>
  <a href="LICENSE"><img alt="License" src="https://img.shields.io/github/license/StealthPassTeam/StealthPass?color=orange"></a>
  <a href="https://github.com/StealthPassTeam/StealthPass/stargazers"><img alt="Stars" src="https://img.shields.io/github/stars/StealthPassTeam/StealthPass?style=flat&logo=github&color=orange"></a>
</p>

**StealthPass** is an enterprise-grade **anti-DPI tunnel engine** written in **Go**, optimized for Iran ⇄ abroad (kharej) server setups. One binary. **15 transports**. Gaming optimized. DPI evasion ready.

One self-contained binary with an interactive CLI **and** a secured web dashboard — run and manage everything with or without a terminal. Includes **stealthpass-tunnel** configuration tool for instant setup.

### 🎯 15 Supported Transports

**Gaming (Low-Latency):** Trojan • gRPC • UDP  
**General (Balanced):** VLESS • QUIC • TCP • WebSocket  
**Stealth (DPI Evasion):** Reality TLS • STEALTH • XDI • SPOOF • PCK • TCP-Mux  

<p align="center">
  <b><a href="README_FA.md">🇮🇷 راهنمای فارسی</a></b> ·
  <b><a href="#quick-start">⚡ Quick Start</a></b> ·
  <b><a href="#features">✨ Features</a></b> ·
  <b><a href="docs/README.md">📚 Docs</a></b>
</p>

---

## ✨ Features

✅ **15 Transports** - All in one binary (no modular downloads)
✅ **Gaming Optimized** - Trojan/gRPC/UDP with health failover
✅ **DPI Evasion** - Reality TLS fingerprinting, IP spoofing, ICMP
✅ **CLI Tool** - `stealthpass-tunnel` for one-command setup
✅ **3 Presets** - Gaming/General/Stealth modes
✅ **BBR Optimization** - Google congestion control
✅ **Zero-Copy** - Kernel bypass (Linux)
✅ **FEC** - Forward error correction for lossy paths
✅ **Health Failover** - Multi-exit automatic switching
✅ **Web Dashboard** - Real-time monitoring on port 7777

---

## 🚀 Installation

### Recommended: Use CLI Tool

```bash
# Ubuntu 20.04+, as root
curl -fsSL https://github.com/StealthPassTeam/StealthPass/releases/download/v2.0.0/install.sh | bash

# Or manual
wget https://github.com/StealthPassTeam/StealthPass/releases/download/v2.0.0/stealthpass_linux_amd64.tar.gz
tar xzf stealthpass_linux_amd64.tar.gz
sudo mv stealthpass stealthpass-tunnel /usr/local/bin/
sudo chmod +x /usr/local/bin/stealthpass*
```

### From Source

```bash
git clone https://github.com/StealthPassTeam/StealthPass.git
cd StealthPass
GOOS=linux GOARCH=amd64 go build -o stealthpass .
GOOS=linux GOARCH=amd64 go build -o stealthpass-tunnel ./cmd/tunnel-setup/
sudo mv stealthpass stealthpass-tunnel /usr/local/bin/
```

---

## ⚡ Quick Start

### 1️⃣ Gaming Mode (Trojan - Ultra Low Latency)

```bash
# Generate config
stealthpass-tunnel gaming --protocol trojan --dry-run

# Save and run
stealthpass-tunnel gaming --protocol trojan --output /etc/stealthpass/gaming.toml
sudo stealthpass -c /etc/stealthpass/gaming.toml
```

### 2️⃣ General Mode (VLESS - Balanced)

```bash
stealthpass-tunnel general --protocol vless --dry-run
stealthpass-tunnel general --protocol vless --output /etc/stealthpass/general.toml
sudo stealthpass -c /etc/stealthpass/general.toml
```

### 3️⃣ Stealth Mode (Reality TLS - DPI Evasion)

```bash
stealthpass-tunnel stealth --protocol reality --dry-run
stealthpass-tunnel stealth --protocol reality --output /etc/stealthpass/stealth.toml
sudo stealthpass -c /etc/stealthpass/stealth.toml
```

---

## 🛠️ CLI Tool Usage

```bash
# Gaming presets
stealthpass-tunnel gaming --protocol trojan        # Ultra-low latency
stealthpass-tunnel gaming --protocol grpc          # HTTP/2 based
stealthpass-tunnel gaming --protocol udp           # Raw speed

# General presets
stealthpass-tunnel general --protocol vless        # Xray compatible
stealthpass-tunnel general --protocol quic         # Modern encrypted
stealthpass-tunnel general --protocol ws           # WebSocket

# Stealth presets
stealthpass-tunnel stealth --protocol reality      # Chrome fingerprint
stealthpass-tunnel stealth --protocol stealth      # Noise encrypted
stealthpass-tunnel stealth --protocol xdi          # ICMP tunneling

# IP Pool Management
stealthpass-tunnel ip-pool add 10.0.0.1
stealthpass-tunnel ip-pool list
stealthpass-tunnel ip-pool rotate

# Validate config
stealthpass-tunnel validate /etc/stealthpass/tunnel.toml

# Dry-run (show without saving)
stealthpass-tunnel gaming --dry-run
```

---

## 📊 Transport Comparison

| Transport | Latency | Speed | DPI Block | Use Case |
|-----------|---------|-------|-----------|----------|
| **Trojan** | ⭐ Ultra-Low | ⭐⭐⭐⭐ | ✅ Detectable | Gaming |
| **gRPC** | ⭐⭐ Low | ⭐⭐⭐ | ✅ Hard to detect | Gaming, General |
| **UDP** | ⭐ Ultra-Low | ⭐⭐⭐⭐⭐ | ❌ Detectable | Gaming, Raw |
| **VLESS** | ⭐⭐ Low | ⭐⭐⭐ | ✅ Detectable | General, Xray |
| **QUIC** | ⭐⭐ Low | ⭐⭐⭐⭐ | ✅ Detectable | General, Modern |
| **TCP** | ⭐⭐⭐ Medium | ⭐⭐⭐ | ✅ Detectable | General, Fallback |
| **WebSocket** | ⭐⭐⭐ Medium | ⭐⭐⭐ | ✅ Looks like HTTP | General, CDN |
| **Reality** | ⭐⭐ Low | ⭐⭐⭐ | ✅ Undetectable* | Stealth |
| **STEALTH** | ⭐⭐ Low | ⭐⭐⭐ | ✅ Undetectable* | Stealth, Noise |
| **XDI (ICMP)** | ⭐ Ultra-Low | ⭐⭐ | ✅ Bypass TCP/UDP filter | Stealth, Bypass |

*Reality TLS spoofs Chrome 120+ exactly; STEALTH uses Noise protocol

---

## 🎮 Gaming Optimizations

Trojan transport with gaming preset automatically configures:

```toml
[client]
transport = "trojan"
connection_pool = 24          # More parallel connections
keepalive_period = 30         # Responsive keepalive
nodelay = true                # TCP_NODELAY enabled
aggressive_pool = true        # Pre-allocate connections
health_failover = true        # Switch on latency spike

[server]
preset = "best-performance"
bandwidth_mbps = 0            # Unlimited
max_connections = 0           # Unlimited
```

**Result:** < 50ms latency, optimal for Dota 2, CS2, Valorant

---

## 🛡️ DPI Evasion

### Reality TLS (Undetectable)

Spoofs Chrome 120+ TLS ClientHello exactly:

```bash
stealthpass-tunnel stealth --protocol reality
```

### STEALTH (Noise Encrypted)

Random appearance on wire:

```bash
stealthpass-tunnel stealth --protocol stealth
```

### XDI (ICMP Tunneling)

When TCP/UDP are filtered but ping works:

```bash
stealthpass-tunnel stealth --protocol xdi
```

---

## 📋 System Requirements

- **OS:** Ubuntu 20.04+ (Linux only)
- **Arch:** x86-64, ARM64, RISC-V
- **Privileges:** root (for XDI, PCK)
- **Ports:** Configurable (default 443, 80, 8080)

---

## 📚 Documentation

| Link | Purpose |
|------|---------|
| [Quick Setup](docs/QUICKSTART.md) | 5-minute tutorial |
| [Configuration](docs/CONFIG.md) | Full config reference |
| [Protocols](docs/PROTOCOLS.md) | Transport details |
| [DPI Evasion](docs/DPI_EVASION.md) | Anti-blocking guide |
| [Performance](docs/TUNING.md) | Optimization guide |
| [Deployment](docs/DEPLOYMENT.md) | Production setup |
| [CLI Reference](docs/CLI.md) | All commands |

---

## 🔐 Security

- **Token-based auth** (16+ chars)
- **Noise Protocol** - Random appearance
- **Reality TLS** - Undetectable browser fingerprint
- **ChaCha20-Poly1305** - Authenticated encryption
- **PROXY Protocol v2** - Real client IPs
- **No fingerprint** - Stealth mode undetectable

---

## 🎯 Key Differences from v1.x

| Feature | v1.x | v2.0 |
|---------|------|------|
| **Transports** | 11 | **15** (added Trojan, gRPC, VLESS, Reality) |
| **Gaming Presets** | Manual tuning | Auto-optimized |
| **CLI Tool** | Menu-driven | **stealthpass-tunnel** command-line |
| **DPI Evasion** | Limited | **Reality TLS** + enhanced |
| **Configuration** | Complex | One-command setup |
| **IP Pool** | Manual | Auto-managed |

---

## 📞 Support

- **GitHub Issues:** [Report bugs](https://github.com/StealthPassTeam/StealthPass/issues)
- **Telegram:** [@StealthPassChat](https://t.me/StealthPassChat)
- **Documentation:** [Full docs](docs/README.md)

---

## 📄 License

**GNU Affero General Public License v3.0 (AGPL-3.0)**  
Copyright © 2026 StealthPass Team

---

## 🙏 Contributing

Contributions welcome! Please:
1. Fork the repo
2. Create feature branch
3. Submit PR with tests

---

**StealthPass v2.0.0** — Enterprise tunnel engine, 15 transports, gaming optimized, DPI evasion ready.

*Built in Go • For Ubuntu/Linux • Production-grade*

<p align="center">
  <b><a href="tutorial/README.md">📘 Setup tutorials</a></b> ·
  <b><a href="docs/README.md">📚 Documentation</a></b> ·
  <b><a href="README_FA.md">🇮🇷 راهنمای فارسی</a></b> ·
  <b><a href="https://t.me/BlackProtocols">Telegram Channel</a></b> ·
  <b><a href="https://t.me/BlackProtocolsGroup">Telegram Group</a></b>

</p>

---

## How it works

<p align="center"><img src="img/architecture.svg" alt="Backpack architecture: end users reach a forwarded port on the Iran server, the engine carries it through one transport to the kharej client, which forwards it to the real service. The client dials the server." width="100%"></p>

```
  end users ──▶  IRAN server  ══ tunnel ══▶  KHAREJ server  ──▶  real service
                 "Setup Iran"                 "Setup Kharej"      
                 exposes the ports            dials out to Iran     
```

An end user connects to a **forwarded port** on the Iran server; the engine
carries it through **one transport** to the kharej client, which hands it to the
**real service**. In the **reverse** tunnel above the connection is dialed **by
the client** (kharej → Iran), so the far side needs no open inbound port.

### Three shapes

The ports never move: Iran exposes them, kharej holds the real service. What
changes is who reaches out first, and what the tunnel carries.

| | Who dials | What it carries | Use it when |
|---|---|---|---|
| **Reverse** | kharej → Iran | forwarded ports | the usual case — Iran can accept an inbound connection |
| **Direct** | Iran → kharej | a private network, and forwarded ports over it | an inbound connection to Iran does not get through |

Both are built from **Setup Iran** and **Setup Kharej**: pick the machine you
are on, and the wizard asks which direction you want and writes the config
itself.

A direct tunnel is a full IP tunnel — an interface on each host carrying whole
IP packets, wrapped in Backpack's own GRE inside a Noise session and handed to
one of three carriers. It measures its own MTU once it is up, which is the
setting that fails worst when it is wrong.

**→ [Direct tunnel](docs/l3-direct-tunnel.md)**

---

## Install

One command as root on the VPS. It downloads the prebuilt release for your
architecture, **verifies it against the published checksum**, installs it, and
opens the menu:

```bash
bash <(curl -fsSL https://raw.githubusercontent.com/AminMGMT/StealthPass/main/install.sh)
```

Reopen the menu any time with `sudo backpack`.

> **No internet on the server?** There is a full offline path — copy one archive
> over and go. Building from source works as a fallback too.
> **→ [Installing Backpack](docs/install.md)**

---

## Quick start

**Get the roles right first** — this is the one thing people trip on:

| Server | Where | Menu option | Why |
|--------|-------|-------------|-----|
| **Iran** | entry point | **1. Setup Iran** | It exposes the ports; users connect to the **Iran IP**. |
| **Kharej** | exit / origin | **2. Setup Kharej** | It dials the Iran server and forwards to the real service. |

**Always set up the Iran server first** — the client needs the Iran address and
the token the server generates.

```bash
# on the IRAN server
sudo backpack   →  1. Setup Iran
#   transport → tunnel port → name → COPY THE TOKEN → exposed ports
#   → UDP? → preset (Turbo) → done

# on the KHAREJ server
sudo backpack   →  2. Setup Kharej
#   same transport → Iran IP + same tunnel port → name → SAME TOKEN
#   → same preset → done
```

Then `Manage → Status` to see both ends, and `Manage → Health Check` if anything
looks wrong — it prints a fix under each problem.

**→ [Before you start](tutorial/before-you-start.md)** covers the roles, the
token, the port mapping and the firewall in full. Every transport then has its
own step-by-step page.

---

## Pick a transport

Thirteen to choose from, so you match the route instead of fighting it. Not sure?
**Manage → Link Test** measures your actual route and recommends one.

| Transport | Reach for it when | Guide |
|---|---|---|
| **TCP** | you are not sure — this is the starting point | [→](tutorial/tcp.md) |
| **TCP Mux** | the service opens many short connections | [→](tutorial/tcp-mux.md) |
| **TCP + Stealth** | filtering is heavy — Noise-encrypted, **no fingerprint at all** | [→](tutorial/tcp-stealth.md) |
| **TCP + PCK** | TCP connects then stalls, resets or is throttled | [→](tutorial/tcp-pck.md) |
| **UDP + KCP + FEC** | gaming or a lossy route — always-on error correction | [→](tutorial/udp-kcp-fec.md) |
| **UDP + QUIC** | you want to test an encrypted, self-tuning UDP carrier | [→](tutorial/udp-quic.md) |
| **WS / WS Mux** | only HTTP gets through, or you want a CDN in front | [→](tutorial/websocket.md) |
| **WSS / WSS Mux** | it should look like an ordinary HTTPS website | [→](tutorial/websocket-tls.md) |
| **xDi (ICMP)** | TCP and UDP are filtered but ping works | [→](tutorial/xdi-icmp.md) |

**Every transport explained → [docs/transports.md](docs/transports.md)**

> **The path blocks or counts by source address?** That is **IP Spoofing**, and
> it is a carrier of the **direct tunnel** rather than one of the transports
> above — see **[docs/ip-spoofing.md](docs/ip-spoofing.md)**.

> **Filtered or dirty server?** **TCP + Stealth** or **WSS** get the tunnel
> through DPI — proven in the field. An IP blocked at the network layer, or a
> "dirty" exit, is a clean-IP or CDN-edge matter rather than a transport one —
> see [when a server is filtered or dirty](docs/filtered-or-dirty-ip.md).

---

## Why Backpack?

- **UDP on any forwarded port** — Xray/3x-ui, Shadowsocks, WireGuard, DNS and
  games, on **every** transport, with one switch.
  [How](tutorial/udp-forwarding.md)
- **No fingerprint** — Stealth looks like random bytes; WSS dials with a real
  **Chrome** TLS handshake and answers every probe with a **decoy website**.
- **Gaming-grade UDP** — KCP with always-on FEC repairs loss instead of waiting
  for a retransmit, plus **multi-exit failover** that steers to the healthiest
  server as routes degrade.
- **Nothing left broken** — updates and edits that break a tunnel **revert
  themselves**, and a watchdog restarts a dropped tunnel within ~1 minute from
  its own service.
- **It tells you what is wrong** — Health Check prints a fix under each problem;
  Link Test measures the route and recommends a transport and its timers.
- **Telegram from Iran** — status and alerts reach Telegram by going out through
  a tunnel peer, choosing the tunnel itself and moving when one dies.
- **Offline installer** — install or update with **no internet at all**.

<details>
<summary><b>The full feature list</b></summary>

**Performance** — four presets (Balance, **Turbo**, Aggressive, and Throughput on
KCP) fill in every tuning value at once; **Optimize** applies kernel/network
tuning (BBR + fq, buffer ceilings, file limits); **Link Test** derives the
liveness timers from your real round trip.

**Reliability** — automatic failover to backup addresses, with **health scoring**
(`rtt + 2·jitter + 20·loss%`) or load balancing across all of them; self-healing
watchdog; automatic rollback; systemd services that survive reboots.

**Security** — the token never travels in the clear on an encrypted transport
(Stealth and KCP derive keys from it, WSS binds the credential to the TLS
session); PROXY protocol v2 for real client IPs; per-tunnel connection and
bandwidth caps; login-protected dashboard; SHA-256 verified downloads, and
anything unverifiable is refused rather than installed.

**Management** — an interactive CLI where every option explains itself; setup
checks the address you give it (CDN in front, AAAA records); CDN-edge dialing;
JSON logging; auto-refresh every N hours; a built-in SOCKS5/HTTP proxy so the
tunnel exit can be its own backend.

**Monitoring** — web dashboard on port 7777 with live CPU/RAM/disk/traffic and
per-tunnel status, ping and logs; metrics including KCP retransmits, loss and FEC
repairs, kept across restarts; Telegram alerts with a recovery message for each.

**Maintenance** — one-file backup of every tunnel, the panel password, Telegram
settings, TLS certificates and the schedule; verified updates on a stable or beta
channel.

</details>

---

## Documentation

| | |
|---|---|
| **[📘 Tutorials](tutorial/README.md)** | Step-by-step setup, one page per transport — every question the wizard asks, with the answer to give |
| **[📚 Docs](docs/README.md)** | Reference: what each part is, and every setting it has |
| **[🖥 CLI menu reference](docs/cli-menu.md)** | Every option in every menu, including the advanced Fine Tune settings |
| **[🔀 Transports](docs/transports.md)** | All thirteen, compared and explained |
| **[🎭 IP Spoofing](docs/ip-spoofing.md)** | The forged-source carrier, setting by setting |
| **[📡 Forwarded UDP](docs/forwarded-udp.md)** | Read this if UDP does not pass through |

Both sections are also summarised in Persian at the bottom of every page.

---

## Screenshots

| CLI menu | Web panel |
|----------|-----------|
| ![CLI menu](img/cli-Screenshot.png) | ![Web panel](img/web-panel-Screenshot.png) |

| Tunnel management | Telegram bot |
|-------------------|--------------|
| ![Tunnel management](img/cli-manage-Screenshot.png) | ![Telegram bot](img/tg-bot-Screenshot.png) |

---

## Support & donate

If Backpack helps you, a star or a small tip is appreciated. 🙏

- Telegram channel: **[@BlackProtocols](https://t.me/BlackProtocols)**
- Telegram Group: **[@BlackProtocolsGroup](https://t.me/BlackProtocolsGroup)**

| Coin | Address |
|------|---------|
| **Tron (TRX)** | `TTzuUAtsEsrLgNpFVLNTyLVJVRRFNWESYc` |
| **USDT (BEP20)** | `0xc112AE9bfF7c59dEcFb34E988A397848D3093E82` |
| **Toncoin (TON)** | `UQD9g40QubAICJ6zPqegtCY7s-joMx2DB8aIqA0xF1aHoCDs` |

---

## License

**Copyright © 2026 Amin Mohammadi (AminMGMT).**
Released under the **GNU Affero General Public License v3.0 (AGPL-3.0)** — see
[LICENSE](LICENSE) and [NOTICE](NOTICE).
