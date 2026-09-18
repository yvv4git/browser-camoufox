# examples/requests

Lists the network requests made by a page: the main document plus CSS,
scripts, images and trackers. Reads the Performance API with polling and
deduplication, so late-loaded resources are captured too.

## Run

```sh
go run ./requests https://www.wikipedia.org
go run ./requests -wait 20s https://rutube.ru
```

## Flags

| Flag      | Default                     | Description                             |
| --------- | --------------------------- | --------------------------------------- |
| `-addr`   | `http://localhost:9377`     | HTTP endpoint of the container          |
| `-url`    | `https://www.wikipedia.org` | page to load (or a positional arg)      |
| `-session`| `requests-demo`             | session key                             |
| `-wait`   | `15s`                       | how long to poll for late requests      |

## What it demonstrates

Requests in chronological order, starting with the main document:

```text
total requests: 6

 1. [NAVIGATION  ] https://www.wikipedia.org/
 2. [CSS         ] https://www.wikipedia.org/portal/wikipedia.org/assets/img/sprite-e49fbf32.svg
 3. [SCRIPT      ] https://www.wikipedia.org/portal/wikipedia.org/assets/js/index-34f340e24a.js
 4. [IMG         ] https://www.wikipedia.org/portal/wikipedia.org/assets/img/Wikipedia-logo-v2.png
```

## Notes

- The first entry is always the main document.
- `-wait` polls every 1s to catch late-loading trackers and ads.
- `performance.getEntriesByType("resource")` does not report failed or
  blocked requests, nor websockets, so the count may be lower than
  DevTools shows.
## License

MIT, see [LICENSE](../LICENSE).
