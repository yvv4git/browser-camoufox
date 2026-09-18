// Command html fetches page HTML twice: right after navigation and again
// after dynamic content has loaded, so the difference shows whether the
// page renders client-side.
//
// Run from the example directory:
//
//	go run ./html https://www.wikipedia.org
//	go run ./html -wait 5s https://rutube.ru
//	go run ./html -addr http://localhost:9377 -wait 20s https://example.com
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/yvv4git/go-juggler"
)

const _htmlExpr = "document.documentElement.outerHTML"

func main() {
	addr := flag.String("addr", "http://localhost:9377", "HTTP endpoint of the running camofox-browser")
	url := flag.String("url", "https://www.wikipedia.org", "page to load (or a positional arg)")
	session := flag.String("session", "html-demo", "session key")
	wait := flag.Duration("wait", 10*time.Second, "wait for dynamic content")

	flag.Parse()

	if flag.NArg() > 0 {
		*url = flag.Arg(0)
	}

	if err := run(*addr, *session, *url, *wait); err != nil {
		log.Fatal(err)
	}
}

func run(addr, session, url string, wait time.Duration) error {
	d := newDemo(addr)

	tab, err := d.open(session, url)
	if err != nil {
		return err
	}

	defer func() { _ = d.closeTab(session, tab.TabID) }()

	early, err := d.html(session, tab.TabID)
	if err != nil {
		return err
	}

	fmt.Printf("early HTML: %d bytes\n", len(early))

	time.Sleep(wait)

	late, err := d.html(session, tab.TabID)
	if err != nil {
		return err
	}

	delta := len(late) - len(early)

	fmt.Printf("late  HTML: %d bytes (delta %d)\n", len(late), delta)

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

	return tab, nil
}

func (d *demo) closeTab(session, tabID string) error {
	if err := d.c.CloseTab(d.ctx, tabID, session); err != nil {
		return fmt.Errorf("close tab: %w", err)
	}

	return nil
}

func (d *demo) html(session, tabID string) (string, error) {
	res, err := d.c.Evaluate(d.ctx, tabID, session, _htmlExpr)
	if err != nil {
		return "", fmt.Errorf("evaluate: %w", err)
	}

	raw, ok := res.Result.(string)
	if !ok {
		return "", fmt.Errorf("unexpected result type: %T", res.Result)
	}

	return raw, nil
}
