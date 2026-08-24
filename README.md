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
- [Makefile](#makefile)
- [Environment Variables](#environment-variables)
- [Ports](#ports)
- [License](#license)
- [Support](#support)

## Features

- Camoufox (Firefox-based) with headless mode via Xvfb
- Built-in yt-dlp for video downloading
- Proxy support
- Log size limits
- Multi-architecture: amd64 / arm64

## Quick Start

```bash
# Build image
make image_build

# Start container
make compose_up

# Stop container
make compose_down
```

## Makefile

| Command        | Description                       |
| -------------- | --------------------------------- |
| `image_build`  | Build Docker image                |
| `image_push`   | Push image to Docker Hub          |
| `image_pull`   | Pull image from Docker Hub        |
| `image_remove` | Remove local Docker image         |
| `compose_up`   | Start container (docker compose)  |
| `compose_down` | Stop container                    |

## Environment Variables

| Variable             | Default               | Description                 |
| -------------------- | --------------------- | --------------------------- |
| `PROXY_HOST`         | `host.docker.internal`| Proxy host                  |
| `PROXY_PORT`         | `8138`                | Proxy port                  |
| `PROXY_SCHEME`       | `http`                | Proxy scheme                |
| `MAX_OLD_SPACE_SIZE` | `128`                 | Node.js memory limit (MB)   |

## Ports

| Port | Description      |
| ---- | ---------------- |
| 9377 | HTTP API server  |

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
