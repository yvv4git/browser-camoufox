// Command tabs opens several tabs in one session, lists them, and closes
// them all.
//
// Run from the example directory:
//
//	go run ./tabs
//	go run ./tabs -count 5
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
	session := flag.String("session", "tabs-demo", "session key")
	count := flag.Int("count", 3, "number of tabs to open")

	flag.Parse()

	if err := run(*addr, *session, *count); err != nil {
		log.Fatal(err)
	}
}

func run(addr, session string, count int) error {
	d := newDemo(addr)
	urls := d.urls()

	for i := 0; i < count; i++ {
		if err := d.openTab(session, urls[i%len(urls)]); err != nil {
			return err
		}
	}

	if err := d.list(session); err != nil {
		return err
	}

	return d.closeAll(session)
}

type demo struct {
	c   *juggler.Client
	ctx context.Context
}

func newDemo(addr string) *demo {
	return &demo{c: juggler.NewClient(addr), ctx: context.Background()}
}

func (d *demo) urls() []string {
	return []string{
		"https://example.com",
		"https://example.org",
		"https://example.net",
		"https://example.edu",
	}
}

func (d *demo) openTab(session, url string) error {
	tab, err := d.c.OpenTab(d.ctx, session, url)
	if err != nil {
		return fmt.Errorf("open tab: %w", err)
	}

	fmt.Printf("opened: %s\n", tab.URL)

	return nil
}

func (d *demo) list(session string) error {
	tabs, err := d.c.ListTabs(d.ctx, session)
	if err != nil {
		return fmt.Errorf("list tabs: %w", err)
	}

	fmt.Printf("total tabs: %d\n", len(tabs.Tabs))

	for _, t := range tabs.Tabs {
		fmt.Printf("  [%s] %s\n", t.Title, t.URL)
	}

	return nil
}

func (d *demo) closeAll(session string) error {
	tabs, err := d.c.ListTabs(d.ctx, session)
	if err != nil {
		return fmt.Errorf("list tabs: %w", err)
	}

	for _, t := range tabs.Tabs {
		if err := d.c.CloseTab(d.ctx, t.TabID, session); err != nil {
			return fmt.Errorf("close tab: %w", err)
		}

		fmt.Printf("closed %s\n", t.TabID[:8])
	}

	return nil
}
