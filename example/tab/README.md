# examples/tab

Exercises every tab operation available in go-juggler against the running
container browser. Each step runs independently; a failed step is printed
with its error and does not stop the rest.

## Run

```sh
go run ./tab -addr http://localhost:9377
```

## Flags

| Flag    | Default                 | Description                    |
| ------- | ----------------------- | ------------------------------ |
| `-addr` | `http://localhost:9377` | HTTP endpoint of the container |

## What it demonstrates

| Step             | Method      | What it does                             |
| ---------------- | ----------- | ---------------------------------------- |
| open             | `OpenTab`   | Opens a new tab in a session (startup)   |
| snapshot         | `Snapshot`  | ARIA accessibility tree with ref IDs     |
| links            | `Links`     | Lists links on the page with pagination  |
| stats            | `Stats`     | URL, visited URLs, ref count             |
| inject button    | `Evaluate`  | Injects a button via JavaScript          |
| click button     | `Click`     | Clicks the injected button by CSS `#btn` |
| navigate         | `Navigate`  | Loads example.org                        |
| back             | `Back`      | Goes back in history                     |
| forward          | `Forward`   | Goes forward in history                  |
| inject form      | `Evaluate`  | Injects a form with an input field       |
| type text        | `Type`      | Types into the input field `#q`          |
| press enter      | `Press`     | Presses Enter to submit                  |
| scroll           | `Scroll`    | Scrolls the page down 500px              |
| refresh          | `Refresh`   | Reloads the current page                 |
| screenshot       | `Screenshot`| Saves a PNG screenshot                   |
| close            | `CloseTab` + `CloseSession` | Tears everything down |

## Sample output

```text
open             -> https://example.com/
snapshot         -> 2234 chars, 5 refs
links            -> 1 total (first 1)
stats            -> url=https://example.com/ visited=1
inject button    -> injected
click button     -> clicked #btn
navigate         -> example.org
back             -> ok
forward          -> ok
inject form      -> injected
type text        -> typed into #q
press enter      -> pressed Enter
scroll           -> scrolled 500px
refresh          -> reloaded
screenshot       -> saved screenshot.png (36846 bytes)
session closed
```
## License

MIT, see [LICENSE](../LICENSE).
