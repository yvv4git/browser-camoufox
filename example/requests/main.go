// Command requests lists the network requests made by a page: the main
// document plus CSS, scripts, images and trackers. It reads the
// Performance API with polling, so late-loaded resources are captured
// too.
//
// Run from the example directory:
//
//	go run ./requests https://www.wikipedia.org
//	go run ./requests -wait 20s https://rutube.ru
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/yvv4git/go-juggler"
)

func main() {
	addr := flag.String("addr", "http://localhost:9377", "HTTP endpoint of the running camofox-browser")
	url := flag.String("url", "https://www.wikipedia.org", "page to load (or a positional arg)")
	session := flag.String("session", "requests-demo", "session key")
	wait := flag.Duration("wait", 15*time.Second, "how long to poll for late requests")

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

	entries, err := d.c.PollNetworkRequests(d.ctx, tab.TabID, session, wait, time.Second)
	if err != nil {
		return fmt.Errorf("poll requests: %w", err)
	}

	fmt.Printf("total requests: %d\n\n", len(entries))

	for i, e := range entries {
		name := e.Name

		if len(name) > 100 {
			name = name[:97] + "..."
		}

		fmt.Printf("%2d. [%-12s] %s\n", i+1, strings.ToUpper(e.Type), name)
	}

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
