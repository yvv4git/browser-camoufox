# examples/html

Fetches page HTML twice: right after navigation and again after dynamic
content has loaded. A growing delta means the page renders client-side.

## Run

```sh
go run ./html https://www.wikipedia.org
go run ./html -wait 5s https://rutube.ru
go run ./html -addr http://localhost:9377 -wait 20s https://example.com
```

## Flags

| Flag      | Default                   | Description                     |
| --------- | ------------------------- | ------------------------------- |
| `-addr`   | `http://localhost:9377`   | HTTP endpoint of the container  |
| `-url`    | `https://www.wikipedia.org` | page to load (or positional)  |
| `-session`| `html-demo`               | session key                     |
| `-wait`   | `10s`                     | wait for dynamic content        |

## What it demonstrates

- Evaluating JavaScript in the page (`document.documentElement.outerHTML`).
- Reading the page twice and comparing sizes.

## Sample output

```text
early HTML: 121000 bytes
late  HTML: 121051 bytes (delta 51)
```

`delta` is `len(late) - len(early)` and is negative when the page shrinks.
## License

MIT, see [LICENSE](../LICENSE).
