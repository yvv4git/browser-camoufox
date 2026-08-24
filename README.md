# Browser Camoufox

Docker image with [Camoufox][camoufox] browser —
a Firefox fork obfuscated for browser automation
with minimal fingerprinting.

[camoufox]: https://github.com/nicegram/nicegram-web/blob/main/nicegram-web.md

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
