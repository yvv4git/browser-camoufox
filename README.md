# Browser Camoufox

![Browser Camoufox logo](./assets/browser-camoufox.jpeg)

[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Image-2496ED?style=flat-square&logo=docker&logoColor=white)](Dockerfile)
[![Camoufox](https://img.shields.io/badge/Camoufox-135.0.1-0098EA?style=flat-square)](https://camoufox.com)
[![Buy me a TON](https://img.shields.io/badge/Buy%20me%20a%20TON-0098EA?style=flat-square)](#support)

Docker image with [Camoufox][camoufox] browser —
a Firefox fork obfuscated for browser automation
with minimal fingerprinting.

[camoufox]: https://github.com/nicegram/nicegram-web/blob/main/nicegram-web.md

## Table of contents

- [Features](#features)
- [Quick Start](#quick-start)
- [VNC](#vnc)
- [Proxy](#proxy)
- [Makefile](#makefile)
- [Environment Variables](#environment-variables)
- [Ports](#ports)
- [License](#license)
- [Support](#support)

## Features

- Camoufox (Firefox-based) with headless mode via Xvfb
- Built-in yt-dlp for video downloading
- HTTP proxy support with authentication
- Interactive VNC access via noVNC web interface
- Log size limits (10 MB x 3 files)
- Multi-architecture: amd64 / arm64

## Quick Start

```bash
# Build image
make image_build

# Start container (without VNC)
make compose_up

# Start container (with VNC)
make compose_vnc

# Stop container
make compose_down
```

## VNC

VNC provides interactive access to the browser via a web interface.
Enable it in `.env`:

```bash
ENABLE_VNC=1
VNC_PASSWORD=your-secret-password
```

Then start the container:

```bash
docker compose up -d
```

Open `http://localhost:6080/vnc.html` in your browser and enter the VNC password.

To disable VNC, set `ENABLE_VNC=0` in `.env` and restart:

```bash
docker compose down && docker compose up -d
```

## Proxy

The browser supports HTTP proxy with optional authentication.
Configure it in `.env`:

```bash
PROXY_HOST=your-proxy-host
PROXY_PORT=8080
PROXY_SCHEME=http
PROXY_USERNAME=user
PROXY_PASSWORD=pass
```

If `PROXY_HOST` is empty or not set, the browser runs without proxy.

Restart the container after changes:

```bash
docker compose down && docker compose up -d
```

## Makefile

| Command        | Description                       |
| -------------- | --------------------------------- |
| `image_build`  | Build Docker image                |
| `image_push`   | Push image to Docker Hub          |
| `image_pull`   | Pull image from Docker Hub        |
| `image_remove` | Remove local Docker image         |
| `compose_up`   | Start container (docker compose)  |
| `compose_vnc`  | Start container with VNC enabled  |
| `compose_down` | Stop container                    |

## Environment Variables

### General

| Variable             | Default | Description                   |
| -------------------- | ------- | ----------------------------- |
| `MAX_OLD_SPACE_SIZE` | `128`   | Node.js V8 heap limit (MB)    |

### VNC Settings

| Variable        | Default | Description                           |
| --------------- | ------- | ------------------------------------- |
| `ENABLE_VNC`    | `0`     | Enable VNC (`1` to enable)            |
| `VNC_PASSWORD`  | -       | Password for VNC access               |

### Proxy Settings

| Variable         | Default | Description                              |
| ---------------- | ------- | ---------------------------------------- |
| `PROXY_HOST`     | -       | Proxy hostname or IP                     |
| `PROXY_PORT`     | -       | Proxy port                               |
| `PROXY_SCHEME`   | `http`  | Proxy scheme (`http` or `https`)         |
| `PROXY_USERNAME` | -       | Proxy auth username                      |
| `PROXY_PASSWORD` | -       | Proxy auth password                      |

## Ports

| Port | Description             |
| ---- | ----------------------- |
| 9377 | HTTP API server         |
| 6080 | noVNC web interface     |

## License

MIT, see [LICENSE](LICENSE). See [NOTICE](NOTICE) for attribution.

## Support

<p align="center">
  <a href="https://tonviewer.com/UQCcbp-mue-7HTjDNQ_ZrKtg-tUxIFu817APmItjXasiBGP3">
    <img src="https://img.shields.io/badge/Buy%20me%20a%20TON-0098EA?style=for-the-badge">
  </a>
</p>

<p align="center">
  If this tool helps you, consider buying me a coffee!
</p>
