# examples/tabs

Working with multiple tabs in one session: open several, list them, and
close them all.

## Run

```sh
go run ./tabs
go run ./tabs -count 5
```

## Flags

| Flag      | Default                 | Description                  |
| --------- | ----------------------- | ---------------------------- |
| `-addr`   | `http://localhost:9377` | HTTP endpoint of the container |
| `-session`| `tabs-demo`             | session key                  |
| `-count`  | `3`                     | number of tabs to open       |

## What it demonstrates

- Opening multiple tabs in one session (`POST /tabs/open`).
- Listing the session tabs (`GET /tabs?userId=...`).
- Closing each tab (`DELETE /tabs/:tabId`).

## Sample output

```text
opened: https://example.com/
opened: https://example.org/
opened: https://example.net/
total tabs: 3
  [Example Domain] https://example.com/
  [Example Domain] https://example.org/
  [Example Domain] https://example.net/
closed acd74c0c
closed 8c0951ac
closed 57adc6dd
```
## License

MIT, see [LICENSE](../LICENSE).
