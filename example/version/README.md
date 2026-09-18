# examples/version

Asks the Camoufox browser **running in the container** to report its own
version and fingerprint. The values are read from `navigator.*` via
`POST /tabs/:tabId/evaluate` -- i.e. the browser itself answers, not
files on disk.

## Run

```sh
go run ./version -addr http://localhost:9377
```

## Flags

| Flag    | Default                 | Description                    |
| ------- | ----------------------- | ------------------------------ |
| `-addr` | `http://localhost:9377` | HTTP endpoint of the container |

## What it demonstrates

Opening a tab and evaluating JavaScript inside the real browser context.
The evaluated expression collects fingerprint data:

```javascript
JSON.stringify({
  userAgent: navigator.userAgent,
  appVersion: navigator.appVersion,
  platform: navigator.platform,
  vendor: navigator.vendor,
  hardwareConcurrency: navigator.hardwareConcurrency
})
```

## Sample output

```text
userAgent              Mozilla/5.0 (X11; Linux x86_64; rv:152.0) Gecko/20100101 Firefox/152.0
appVersion             5.0 (X11)
platform               Linux x86_64
vendor
hardwareConcurrency    8
```

The user agent reports the Firefox engine (152.0) bundled into Camoufox.
The `vendor` field is empty in modern Camoufox.
## License

MIT, see [LICENSE](../LICENSE).
