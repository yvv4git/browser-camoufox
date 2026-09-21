# Environment Variables

All variables are set in `.env` and passed into the container via
`docker-compose.yaml`. Restart the container after changes:

```bash
docker compose down && docker compose up -d
```

## Server and Security

| Variable                         | Default               | Description                                                        |
| -------------------------------- | --------------------- | ------------------------------------------------------------------ |
| `CAMOFOX_PORT`                   | `9377`                | HTTP API server port                                               |
| `PORT`                           | `9377`                | Port fallback (Fly.io, Railway, etc.)                              |
| `CAMOFOX_BIND_HOST`              | -                     | Bind host: `127.0.0.1` (loopback only), `0.0.0.0` (all IPv4)       |
| `NODE_ENV`                       | `development`         | Run environment (`production` hides error details)                 |
| `CAMOFOX_API_KEY`                | -                     | Key for cookie import endpoint (`POST /sessions/:userId/cookies`)  |
| `CAMOFOX_ADMIN_KEY`              | -                     | Key for `POST /stop`                                               |
| `CAMOFOX_ACCESS_KEY`             | -                     | Bearer token for all routes except `/health`                       |
| `CAMOFOX_EVALUATE_MAX_BODY_SIZE` | `1mb`                 | Max JSON body size for `POST /tabs/:tabId/evaluate`                |
| `MAX_OLD_SPACE_SIZE`             | `128`                 | Node.js V8 heap limit, MB                                          |

## VNC

| Variable         | Default | Description                              |
| ---------------- | ------- | ---------------------------------------- |
| `ENABLE_VNC`     | `0`     | Enable VNC (`1` enable, `0` disable)     |
| `VNC_PASSWORD`   | -       | VNC access password                      |
| `VNC_PORT`       | -       | VNC server port                          |
| `NOVNC_PORT`     | `6080`  | noVNC web UI port                        |
| `VNC_BIND`       | -       | VNC bind address (e.g. `0.0.0.0`)        |
| `VNC_RESOLUTION` | -       | Screen resolution (e.g. `1920x1080`)     |
| `VIEW_ONLY`      | -       | View-only mode, no input allowed         |

## Proxy

Proxy scheme is always `http`. Proxy is disabled when `PROXY_HOST` is empty.

| Variable                         | Default       | Description                                              |
| -------------------------------- | ------------- | -------------------------------------------------------- |
| `PROXY_HOST`                     | -             | Proxy hostname or IP (simple mode)                       |
| `PROXY_PORT`                     | -             | Proxy port (simple mode)                                 |
| `PROXY_PORTS`                    | -             | Port list `"10001,10002"` or range `"10001-10010"`       |
| `PROXY_USERNAME`                 | -             | Proxy auth username                                      |
| `PROXY_PASSWORD`                 | -             | Proxy auth password                                      |
| `PROXY_STRATEGY`                 | `round_robin` | Mode: `backconnect` (rotating sessions) or empty         |
| `PROXY_PROVIDER`                 | `decodo`      | Provider: `decodo` or `generic`                          |
| `PROXY_BACKCONNECT_HOST`         | -             | Backconnect gateway hostname                             |
| `PROXY_BACKCONNECT_PORT`         | `7000`        | Backconnect gateway port                                 |
| `PROXY_COUNTRY`                  | -             | Geo-targeting country                                    |
| `PROXY_STATE`                    | -             | Geo-targeting state/region                               |
| `PROXY_CITY`                     | -             | Geo-targeting city                                       |
| `PROXY_ZIP`                      | -             | Geo-targeting postal code                                |
| `PROXY_SESSION_DURATION_MINUTES` | `10`          | Sticky session duration, minutes                         |

With a proxy configured, browser locale, timezone, and geolocation are
derived from the proxy IP automatically. Without a proxy, defaults are
`en-US`, `America/Los_Angeles`, San Francisco.

## Session and Tab Limits

| Variable                           | Default   | Description                                            |
| ---------------------------------- | --------- | ------------------------------------------------------ |
| `MAX_SESSIONS`                     | `50`      | Max concurrent sessions                                |
| `MAX_TABS_PER_SESSION`             | `10`      | Max tabs per session                                   |
| `MAX_TABS_GLOBAL`                  | `50`      | Max tabs in total                                      |
| `MAX_CONCURRENT_PER_USER`          | `3`       | Max concurrent requests per user                       |
| `SESSION_TIMEOUT_MS`               | `1800000` | Inactive session timeout (30 min)                      |
| `TAB_INACTIVITY_MS`                | `300000`  | Close tabs idle longer than this (5 min)               |
| `BROWSER_IDLE_TIMEOUT_MS`          | `300000`  | Kill browser when idle (0 = never)                     |
| `HANDLER_TIMEOUT_MS`               | `30000`   | Max time for any handler (30 s)                        |
| `NAVIGATE_TIMEOUT_MS`              | `25000`   | Navigation timeout (25 s)                              |
| `BUILDREFS_TIMEOUT_MS`             | `12000`   | Snapshot refs build timeout (12 s)                     |
| `NATIVE_MEM_RESTART_THRESHOLD_MB`  | `300`     | Browser restart threshold by native memory             |
| `BROWSER_RSS_RESTART_THRESHOLD_MB` | `1500`    | Browser restart threshold by RSS                       |

## Storage Directories

Useful to mount as volumes to persist state across container recreation.

| Variable                     | Default               | Description                                        |
| ---------------------------- | --------------------- | -------------------------------------------------- |
| `CAMOFOX_COOKIES_DIR`        | `~/.camofox/cookies`  | Directory for cookie files                         |
| `CAMOFOX_UPLOADS_DIR`        | `~/.camofox/uploads`  | Directory for file uploads                         |
| `CAMOFOX_PROFILE_DIR`        | `~/.camofox/profiles` | Directory for persisted sessions                   |
| `CAMOFOX_TRACES_DIR`         | `~/.camofox/traces`   | Directory for trace archives (zip)                 |
| `CAMOFOX_TRACES_MAX_BYTES`   | `52428800` (50 MB)    | Max trace size, larger files removed on startup    |
| `CAMOFOX_TRACES_TTL_HOURS`   | `24`                  | Trace age (hours), older ones swept on startup     |

## Behavior

| Variable                          | Default | Description                                                                                              |
| --------------------------------- | ------- | -------------------------------------------------------------------------------------------------------- |
| `CAMOFOX_INTERACTIVE`             | `off`   | Mode: `desktop` (visible window), `novnc`, `off`                                                         |
| `CAMOFOX_DISABLE_DEFAULT_ADDONS`  | `0`     | `1`/`true` — skip default uBlock Origin (UBO) addon                                                      |
| `PROMETHEUS_ENABLED`              | -       | `1`/`true` — enable Prometheus metrics                                                                   |
| `CAMOUFOX_EXECUTABLE`             | -       | Path to external Camoufox executable (bundle siblings: `properties.json`, `version.json`, `fontconfig/`) |
| `CAMOUFOX_EXECUTABLE_PATH`        | -       | Compat alias for `CAMOUFOX_EXECUTABLE`                                                                   |
| `CAMOFOX_EXECUTABLE_PATH`         | -       | Compat alias for `CAMOUFOX_EXECUTABLE`                                                                   |

## Telemetry

| Variable                          | Default                    | Description                                |
| --------------------------------- | -------------------------- | ------------------------------------------ |
| `CAMOFOX_CRASH_REPORT_ENABLED`    | `true`                     | `false` disables anonymous crash telemetry |
| `CAMOFOX_CRASH_REPORT_URL`        | (Cloudflare Worker)        | Custom telemetry endpoint                  |
| `CAMOFOX_CRASH_REPORT_REPO`       | `jo-inc/camofox-browser`   | Repo for telemetry issues                  |
| `CAMOFOX_CRASH_REPORT_RATE_LIMIT` | `10`                       | Reports per hour limit                     |
| `SENTRY_DSN`                      | -                          | DSN to send errors to Sentry               |

## Internal (usually not required)

| Variable           | Default     | Description                                            |
| ------------------ | ----------- | ------------------------------------------------------ |
| `FLY_MACHINE_ID`   | -           | Fly.io machine ID (horizontal scaling)                 |
| `FLY_APP_NAME`     | -           | Fly.io app name                                        |
| `FLY_API_TOKEN`    | -           | Fly.io API token                                       |
| `XDG_CACHE_HOME`   | `~/.cache`  | Cache directory (`<...>/camoufox` for browser binary)  |
