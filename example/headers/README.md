# examples/headers

Sends custom HTTP headers with the initial navigation and verifies them
on the page that echoes the received headers
(https://httpbin.org/headers).

## Run

```sh
go run ./headers -addr http://localhost:9377
```

## Flags

| Flag    | Default                 | Description                    |
| ------- | ----------------------- | ------------------------------ |
| `-addr` | `http://localhost:9377` | HTTP endpoint of the container |

## What it demonstrates

- Opening a tab with custom headers passed to `OpenTab`.
- Reading the echoed headers back from the accessibility snapshot.

The example sends:

| Header            | Value                   |
| ----------------- | ----------------------- |
| `Authorization`   | `Bearer my-secret-token` |
| `X-Custom-Header` | `custom-value`          |
| `Accept-Language` | `en-US,en;q=0.9`        |

Custom headers are useful for authentication (Bearer/API keys), language
preferences, request tracking and other per-request settings.
`go-juggler` also supports per-navigation headers via `Navigate`.

## Sample output

```text
tab opened: 590e5f78-85a9-46d3-b73d-3c1967e3b58f
snapshot (2325 chars):
- heading "Request Headers" [e1]:
- code:
  - "Accept": "text/html,application/xhtml+xml,..."
  - "Accept-Language": "en-US,en;q=0.9"
  - "Authorization": "Bearer my-secret-token"
  - "X-Custom-Header": "custom-value"
```
## License

MIT, see [LICENSE](../LICENSE).
