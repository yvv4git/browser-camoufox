# examples/screen

Captures a screenshot of a page in the running container browser and
saves it as a PNG file.

## Run

```sh
go run ./screen https://www.wikipedia.org
go run ./screen -out screenshot.png -full https://example.com
```

## Flags

| Flag       | Default                  | Description                             |
| ---------- | ------------------------ | --------------------------------------- |
| `-addr`    | `http://localhost:9377`  | HTTP endpoint of the container          |
| `-url`     | `https://www.wikipedia.org` | page to load (or a positional arg)   |
| `-session` | `screen-demo`            | session key                             |
| `-out`     | `screenshot.png`         | output PNG file                         |
| `-full`    | `false`                  | capture the whole page, not the viewport|

## What it demonstrates

Downloading raw screenshot bytes via `GET /tabs/:tabId/screenshot` and
writing them to disk.

## Sample output

```text
opened https://example.com/
saved screenshot.png (36846 bytes)
```
## License

MIT, see [LICENSE](../LICENSE).
