# examples/inject

Executes a JavaScript expression inside the page of the running container
browser via `POST /tabs/:tabId/evaluate` and reads the mutated value back.

## Run

```sh
go run ./inject -addr http://localhost:9377
```

## Flags

| Flag    | Default                 | Description                    |
| ------- | ----------------------- | ------------------------------ |
| `-addr` | `http://localhost:9377` | HTTP endpoint of the container |

## What it demonstrates

Evaluating arbitrary JavaScript in the page context. The expression sets
`document.title` and returns it, so the side effect is verifiable:

```javascript
(document.title = "Hi, friends!") && document.title
```

## Sample output

```text
engine: camoufox browser: true
result: Hi, friends!
tab closed
```

Requires camofox-browser with the evaluate endpoint (>= 1.4.0).
## License

MIT, see [LICENSE](../LICENSE).
