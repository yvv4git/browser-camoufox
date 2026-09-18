# examples/basic

Full tab lifecycle against the Camoufox browser running in a container:
health check, open a page, snapshot the accessibility tree, close the tab.

## Run

```sh
go run ./basic -addr http://localhost:9377
```

## Flags

| Flag      | Default                 | Description                      |
| --------- | ----------------------- | -------------------------------- |
| `-addr`   | `http://localhost:9377` | HTTP endpoint of the container   |
| `-url`    | `https://example.com`   | page to load                     |

## What it demonstrates

1. **Health check** -- the container browser is alive (`GET /health`).
2. **Open tab** -- new tab in a session (`POST /tabs/open`).
3. **Snapshot** -- the ARIA accessibility tree with ref IDs
   (`GET /tabs/:tabId/snapshot`). Refs like `[e1]` are used to locate
   elements without CSS selectors.
4. **Close tab** -- tears down the tab (`DELETE /tabs/:tabId`).

## Sample output

```text
health: ok=true engine=camoufox sessions=0
tab opened: 25009a48-e2e0-440e-b369-7e04f4ba48dd
snapshot (2232 chars, 5 refs):
textbox "Search the World Wide Web" [e14]:
  ...
tab closed
```
## License

MIT, see [LICENSE](../LICENSE).
