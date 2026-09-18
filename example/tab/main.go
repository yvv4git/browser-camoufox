// Command tab exercises every tab operation from go-juggler against the
// browser running in the container: open, snapshot, links, stats,
// evaluate, click, navigate, back, forward, type, press, scroll, refresh,
// screenshot, close. Each step runs independently; a failure is printed
// for that step only, it is not fatal.
//
// Run from the example directory:
//
//	go run ./tab -addr http://localhost:9377
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yvv4git/go-juggler"
)

const (
	_injectButtonExpr = `document.body.innerHTML = '<button id="btn">Click me</button>'; "ok"`
	_injectFormExpr   = `document.body.innerHTML = '<form><input id="q" type="text" placeholder="Search..."><button>OK</button></form>'; "ok"`
)

const _screenshotName = "screenshot.png"

func main() {
	addr := flag.String("addr", "http://localhost:9377", "HTTP endpoint of the running camofox-browser")

	flag.Parse()

	if err := run(*addr); err != nil {
		log.Fatal(err)
	}
}

func run(addr string) error {
	d := newDemo(addr)

	if _, err := d.open(); err != nil {
		return err
	}

	defer func() { _ = d.c.CloseTab(d.ctx, d.tab.TabID, d.session) }()

	for _, s := range d.steps() {
		out, err := s.fn()
		if err != nil {
			fmt.Printf("%-14s -> %v\n", s.name, err)
		} else {
			fmt.Printf("%-14s -> %s\n", s.name, out)
		}
	}

	return d.close()
}

type step struct {
	name string
	fn   func() (string, error)
}

func (d *demo) steps() []step {
	return []step{
		{name: "snapshot", fn: d.snapshot},
		{name: "links", fn: d.links},
		{name: "stats", fn: d.stats},
		{name: "inject button", fn: d.injectButton},
		{name: "click button", fn: d.clickButton},
		{name: "navigate", fn: d.navigate},
		{name: "back", fn: d.back},
		{name: "forward", fn: d.forward},
		{name: "scroll", fn: d.scroll},
		{name: "refresh", fn: d.refresh},
		{name: "inject form", fn: d.injectForm},
		{name: "type text", fn: d.typeText},
		{name: "press enter", fn: d.pressEnter},
		{name: "screenshot", fn: d.screenshot},
	}
}

type demo struct {
	c       *juggler.Client
	ctx     context.Context
	session string
	tab     *juggler.TabResponse
}

func newDemo(addr string) *demo {
	return &demo{c: juggler.NewClient(addr), ctx: context.Background(), session: "tab-demo"}
}

func (d *demo) open() (string, error) {
	tab, err := d.c.OpenTab(d.ctx, d.session, "https://example.com")
	if err != nil {
		return "", fmt.Errorf("open tab: %w", err)
	}

	d.tab = tab

	return tab.URL, nil
}

func (d *demo) snapshot() (string, error) {
	snap, err := d.c.Snapshot(d.ctx, d.tab.TabID, d.session)
	if err != nil {
		return "", fmt.Errorf("snapshot: %w", err)
	}

	return fmt.Sprintf("%d chars, %d refs", len(snap.Snapshot), snap.RefsCount), nil
}

func (d *demo) links() (string, error) {
	links, err := d.c.Links(d.ctx, d.tab.TabID, d.session, 5, 0)
	if err != nil {
		return "", fmt.Errorf("links: %w", err)
	}

	return fmt.Sprintf("%d total (first %d)", links.Pagination.Total, len(links.Links)), nil
}

func (d *demo) stats() (string, error) {
	stats, err := d.c.Stats(d.ctx, d.tab.TabID, d.session)
	if err != nil {
		return "", fmt.Errorf("stats: %w", err)
	}

	return fmt.Sprintf("url=%s visited=%d", stats.URL, len(stats.VisitedURLs)), nil
}

func (d *demo) injectButton() (string, error) {
	return d.inject(_injectButtonExpr)
}

func (d *demo) injectForm() (string, error) {
	return d.inject(_injectFormExpr)
}

func (d *demo) inject(expr string) (string, error) {
	if _, err := d.c.Evaluate(d.ctx, d.tab.TabID, d.session, expr); err != nil {
		return "", fmt.Errorf("evaluate: %w", err)
	}

	return "injected", nil
}

func (d *demo) clickButton() (string, error) {
	if err := d.c.Click(d.ctx, d.tab.TabID, d.session, "", "#btn"); err != nil {
		return "", fmt.Errorf("click: %w", err)
	}

	return "clicked #btn", nil
}

func (d *demo) navigate() (string, error) {
	if err := d.c.Navigate(d.ctx, d.tab.TabID, d.session, "https://example.org"); err != nil {
		return "", fmt.Errorf("navigate: %w", err)
	}

	return "example.org", nil
}

func (d *demo) back() (string, error) {
	if err := d.c.Back(d.ctx, d.tab.TabID, d.session); err != nil {
		return "", fmt.Errorf("back: %w", err)
	}

	return "ok", nil
}

func (d *demo) forward() (string, error) {
	if err := d.c.Forward(d.ctx, d.tab.TabID, d.session); err != nil {
		return "", fmt.Errorf("forward: %w", err)
	}

	return "ok", nil
}

func (d *demo) typeText() (string, error) {
	if err := d.c.Type(d.ctx, d.tab.TabID, d.session, "", "#q", "Go lang"); err != nil {
		return "", fmt.Errorf("type: %w", err)
	}

	return "typed into #q", nil
}

func (d *demo) pressEnter() (string, error) {
	if err := d.c.Press(d.ctx, d.tab.TabID, d.session, "Enter"); err != nil {
		return "", fmt.Errorf("press: %w", err)
	}

	return "pressed Enter", nil
}

func (d *demo) scroll() (string, error) {
	if err := d.c.Scroll(d.ctx, d.tab.TabID, d.session, "down", 500); err != nil {
		return "", fmt.Errorf("scroll: %w", err)
	}

	return "scrolled 500px", nil
}

func (d *demo) refresh() (string, error) {
	if err := d.c.Refresh(d.ctx, d.tab.TabID, d.session); err != nil {
		return "", fmt.Errorf("refresh: %w", err)
	}

	return "reloaded", nil
}

func (d *demo) screenshot() (string, error) {
	png, err := d.c.Screenshot(d.ctx, d.tab.TabID, d.session, false)
	if err != nil {
		return "", fmt.Errorf("screenshot: %w", err)
	}

	if err := os.WriteFile(_screenshotName, png, 0o644); err != nil {
		return "", fmt.Errorf("write file: %w", err)
	}

	return fmt.Sprintf("saved %s (%d bytes)", _screenshotName, len(png)), nil
}

func (d *demo) close() error {
	if err := d.c.CloseTab(d.ctx, d.tab.TabID, d.session); err != nil {
		return fmt.Errorf("close tab: %w", err)
	}

	if err := d.c.CloseSession(d.ctx, d.session); err != nil {
		return fmt.Errorf("close session: %w", err)
	}

	fmt.Println("session closed")

	return nil
}
