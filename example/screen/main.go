// Command screen captures a screenshot of a page in the container browser
// and saves it as a PNG file.
//
// Run from the example directory:
//
//	go run ./screen https://www.wikipedia.org
//	go run ./screen -out screenshot.png -full https://example.com
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yvv4git/go-juggler"
)

func main() {
	addr := flag.String("addr", "http://localhost:9377", "HTTP endpoint of the running camofox-browser")
	url := flag.String("url", "https://www.wikipedia.org", "page to load (or a positional arg)")
	session := flag.String("session", "screen-demo", "session key")
	out := flag.String("out", "screenshot.png", "output PNG file")
	fullPage := flag.Bool("full", false, "capture the whole page, not just the viewport")

	flag.Parse()

	if flag.NArg() > 0 {
		*url = flag.Arg(0)
	}

	if err := run(*addr, *session, *url, *out, *fullPage); err != nil {
		log.Fatal(err)
	}
}

func run(
	addr string,
	session string,
	url string,
	out string,
	fullPage bool,
) error {
	d := newDemo(addr)

	tab, err := d.open(session, url)
	if err != nil {
		return err
	}

	defer func() { _ = d.closeTab(session, tab.TabID) }()

	png, err := d.c.Screenshot(d.ctx, tab.TabID, session, fullPage)
	if err != nil {
		return fmt.Errorf("screenshot: %w", err)
	}

	if err := os.WriteFile(out, png, 0o644); err != nil {
		return fmt.Errorf("write file %s: %w", out, err)
	}

	fmt.Printf("saved %s (%d bytes)\n", out, len(png))

	return nil
}

type demo struct {
	c   *juggler.Client
	ctx context.Context
}

func newDemo(addr string) *demo {
	return &demo{c: juggler.NewClient(addr), ctx: context.Background()}
}

func (d *demo) open(session, url string) (*juggler.TabResponse, error) {
	tab, err := d.c.OpenTab(d.ctx, session, url)
	if err != nil {
		return nil, fmt.Errorf("open tab: %w", err)
	}

	fmt.Printf("opened %s\n", tab.URL)

	return tab, nil
}

func (d *demo) closeTab(session, tabID string) error {
	if err := d.c.CloseTab(d.ctx, tabID, session); err != nil {
		return fmt.Errorf("close tab: %w", err)
	}

	return nil
}
