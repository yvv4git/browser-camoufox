// Command basic connects to a running Camoufox container through the
// camofox-browser HTTP wrapper and walks through the tab lifecycle:
// health check, open a page, snapshot its accessibility tree, close.
//
// Run from the example directory:
//
//	go run ./basic -addr http://localhost:9377
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/yvv4git/go-juggler"
)

func main() {
	addr := flag.String("addr", "http://localhost:9377", "HTTP endpoint of the running camofox-browser")
	url := flag.String("url", "https://example.com", "page to load")

	flag.Parse()

	if err := run(*addr, *url); err != nil {
		log.Fatal(err)
	}
}

func run(addr, url string) error {
	d := newDemo(addr)

	if err := d.health(); err != nil {
		return err
	}

	tab, err := d.open(url)
	if err != nil {
		return err
	}

	defer func() { _ = d.closeTab(tab.TabID) }()

	return d.snapshot(tab.TabID)
}

type demo struct {
	c       *juggler.Client
	ctx     context.Context
	session string
}

func newDemo(addr string) *demo {
	return &demo{c: juggler.NewClient(addr), ctx: context.Background(), session: "basic-demo"}
}

func (d *demo) health() error {
	health, err := d.c.Health(d.ctx)
	if err != nil {
		return fmt.Errorf("health: %w", err)
	}

	fmt.Printf("health: ok=%v engine=%s sessions=%d\n", health.OK, health.Engine, health.Sessions)

	return nil
}

func (d *demo) open(url string) (*juggler.TabResponse, error) {
	tab, err := d.c.OpenTab(d.ctx, d.session, url)
	if err != nil {
		return nil, fmt.Errorf("open tab: %w", err)
	}

	fmt.Printf("tab opened: %s\n", tab.TabID)

	return tab, nil
}

func (d *demo) snapshot(tabID string) error {
	snap, err := d.c.Snapshot(d.ctx, tabID, d.session)
	if err != nil {
		return fmt.Errorf("snapshot: %w", err)
	}

	fmt.Printf("snapshot (%d chars, %d refs):\n%s\n", len(snap.Snapshot), snap.RefsCount, snap.Snapshot)

	return nil
}

func (d *demo) closeTab(tabID string) error {
	if err := d.c.CloseTab(d.ctx, tabID, d.session); err != nil {
		return fmt.Errorf("close tab: %w", err)
	}

	fmt.Println("tab closed")

	return nil
}
