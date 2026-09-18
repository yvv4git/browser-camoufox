// Command headers shows how to send custom HTTP headers with the initial
// navigation. The page (httpbin.org/headers) echoes the received headers
// into its body, so the snapshot shows exactly what the container browser
// sent to the server.
//
// Run from the example directory:
//
//	go run ./headers -addr http://localhost:9377
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

	flag.Parse()

	if err := run(*addr); err != nil {
		log.Fatal(err)
	}
}

func run(addr string) error {
	d := newDemo(addr)

	tab, err := d.open(d.customHeaders())
	if err != nil {
		return err
	}

	defer func() { _ = d.closeTab(tab.TabID) }()

	snap, err := d.snapshot(tab.TabID)
	if err != nil {
		return err
	}

	fmt.Printf("snapshot (%d chars):\n%s\n", len(snap.Snapshot), snap.Snapshot)

	return nil
}

type demo struct {
	c       *juggler.Client
	ctx     context.Context
	session string
}

func newDemo(addr string) *demo {
	return &demo{c: juggler.NewClient(addr), ctx: context.Background(), session: "headers-demo"}
}

func (d *demo) customHeaders() []juggler.Header {
	return []juggler.Header{
		{Name: "Authorization", Value: "Bearer my-secret-token"},
		{Name: "X-Custom-Header", Value: "custom-value"},
		{Name: "Accept-Language", Value: "en-US,en;q=0.9"},
	}
}

func (d *demo) open(headers []juggler.Header) (*juggler.TabResponse, error) {
	tab, err := d.c.OpenTab(d.ctx, d.session, "https://httpbin.org/headers", headers...)
	if err != nil {
		return nil, fmt.Errorf("open tab: %w", err)
	}

	fmt.Printf("tab opened: %s\n", tab.TabID)

	return tab, nil
}

func (d *demo) snapshot(tabID string) (*juggler.SnapshotResponse, error) {
	snap, err := d.c.Snapshot(d.ctx, tabID, d.session)
	if err != nil {
		return nil, fmt.Errorf("snapshot: %w", err)
	}

	return snap, nil
}

func (d *demo) closeTab(tabID string) error {
	if err := d.c.CloseTab(d.ctx, tabID, d.session); err != nil {
		return fmt.Errorf("close tab: %w", err)
	}

	return nil
}
