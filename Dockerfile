FROM node:22-trixie-slim

ARG CAMOUFOX_VERSION=135.0.1
ARG CAMOUFOX_RELEASE=beta.24
ARG TARGETARCH

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    python3 \
    libgtk-3-0 \
    libdbus-glib-1-2 \
    libxt6 \
    libasound2 \
    libx11-xcb1 \
    libxcomposite1 \
    libxcursor1 \
    libxdamage1 \
    libxfixes3 \
    libxi6 \
    libxrandr2 \
    libxrender1 \
    libxss1 \
    libxtst6 \
    libegl1 \
    libgl1-mesa-dri \
    libgbm1 \
    xvfb \
    fonts-liberation \
    fonts-noto-color-emoji \
    fontconfig \
    ca-certificates \
    curl \
    unzip \
    python3-minimal \
    git \
    && rm -rf /var/lib/apt/lists/*

RUN set -eux; \
    case "${TARGETARCH}" in \
      amd64) CAMOUFOX_ARCH="x86_64"; YTDLP_SUFFIX="" ;; \
      arm64) CAMOUFOX_ARCH="arm64";   YTDLP_SUFFIX="_aarch64" ;; \
      *) echo "Unsupported arch: ${TARGETARCH}" && exit 1 ;; \
    esac; \
    mkdir -p /root/.cache/camoufox; \
    curl -fSL "https://github.com/daijro/camoufox/releases/download/v${CAMOUFOX_VERSION}-${CAMOUFOX_RELEASE}/camoufox-${CAMOUFOX_VERSION}-${CAMOUFOX_RELEASE}-lin.${CAMOUFOX_ARCH}.zip" \
      -o /tmp/camoufox.zip; \
    (unzip -q /tmp/camoufox.zip -d /root/.cache/camoufox || true); \
    chmod -R 755 /root/.cache/camoufox; \
    echo "{\"version\":\"${CAMOUFOX_VERSION}\",\"release\":\"${CAMOUFOX_RELEASE}\"}" > /root/.cache/camoufox/version.json; \
    test -f /root/.cache/camoufox/camoufox-bin && echo "Camoufox installed successfully"; \
    rm /tmp/camoufox.zip; \
    curl -fSL "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp_linux${YTDLP_SUFFIX}" \
      -o /usr/local/bin/yt-dlp; \
    chmod 755 /usr/local/bin/yt-dlp

WORKDIR /app

RUN git clone --depth 1 https://github.com/jo-inc/camofox-browser.git .

RUN PLAYWRIGHT_SKIP_BROWSER_DOWNLOAD=1 npm ci --production && \
    apt-get purge -y --auto-remove build-essential && \
    rm -rf /var/lib/apt/lists/*

RUN sh scripts/install-plugin-deps.sh

ENV NODE_ENV=production
ENV CAMOFOX_PORT=9377

EXPOSE 9377

CMD ["sh", "-c", "node --max-old-space-size=${MAX_OLD_SPACE_SIZE:-128} server.js"]
