# Examples

Go programs that drive the Camoufox browser **running in a container**
through the HTTP API on `http://localhost:9377`, using the
[go-juggler](https://github.com/yvv4git/go-juggler) client.

> All examples connect to an already running container. None of them
> launch a browser on the host machine.

## Prerequisites

Start the container once (see the root [README](../README.md)):

```sh
make compose_up
```

Verify it is up:

```sh
curl -s http://localhost:9377/health
```

## Examples

Every example is a standalone `package main` in its own directory with a
README that explains how to run it and what it demonstrates:

| Example                                | Description                                             |
| -------------------------------------- | ------------------------------------------------------- |
| [basic](./basic/README.md)             | Full tab lifecycle: health, open, snapshot, close       |
| [version](./version/README.md)         | The browser reports its own version and fingerprint     |
| [html](./html/README.md)               | Page HTML before and after dynamic content loads        |
| [headers](./headers/README.md)         | Custom HTTP headers sent with navigation                |
| [inject](./inject/README.md)           | Evaluating JavaScript inside the page                   |
| [requests](./requests/README.md)       | Network requests captured from the Performance API      |
| [screen](./screen/README.md)           | Screenshot of the page as a PNG file                    |
| [tab](./tab/README.md)                 | Every tab operation: click, type, back, refresh, ...    |
| [tabs](./tabs/README.md)               | Multiple tabs: open, list, close                        |

## Run

From this directory:

```sh
go run ./basic -addr http://localhost:9377
go run ./version -addr http://localhost:9377
go run ./html https://www.wikipedia.org
```
## License

MIT, see [LICENSE](../LICENSE).
