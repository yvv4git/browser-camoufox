// Command inject executes a JavaScript expression inside the page loaded
// in the container browser and verifies the side effect by reading the
// mutated value back.
//
// Run from the example directory:
//
//	go run ./inject -addr http://localhost:9377
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/yvv4git/go-juggler"
)

const _injectExpr = `(document.title = "Hi, friends!") && document.title`

func main() {
	addr := flag.String("addr", "http://localhost:9377", "HTTP endpoint of the running camofox-browser")

	flag.Parse()

	if err := run(*addr); err != nil {
		log.Fatal(err)
	}
}

func run(addr string) error {
	d := newDemo(addr)

	if err := d.health(); err != nil {
		return err
	}

	tab, err := d.open()
	if err != nil {
		return err
	}

	defer func() { _ = d.closeTab(tab.TabID) }()

	res, err := d.evaluate(tab.TabID, _injectExpr)
	if err != nil {
		return err
	}

	fmt.Printf("result: %v\n", res)

	return nil
}

type demo struct {
	c       *juggler.Client
	ctx     context.Context
	session string
}

func newDemo(addr string) *demo {
	return &demo{c: juggler.NewClient(addr), ctx: context.Background(), session: "inject-demo"}
}

func (d *demo) health() error {
	health, err := d.c.Health(d.ctx)
	if err != nil {
		return fmt.Errorf("health: %w", err)
	}

	fmt.Printf("engine: %s browser: %v\n", health.Engine, health.BrowserConnected)

	return nil
}

func (d *demo) open() (*juggler.TabResponse, error) {
	tab, err := d.c.OpenTab(d.ctx, d.session, "https://example.com")
	if err != nil {
		return nil, fmt.Errorf("open tab: %w", err)
	}

	return tab, nil
}

func (d *demo) evaluate(tabID, expr string) (any, error) {
	res, err := d.c.Evaluate(d.ctx, tabID, d.session, expr)
	if err != nil {
		return nil, fmt.Errorf("evaluate: %w", err)
	}

	return res.Result, nil
}

func (d *demo) closeTab(tabID string) error {
	if err := d.c.CloseTab(d.ctx, tabID, d.session); err != nil {
		return fmt.Errorf("close tab: %w", err)
	}

	fmt.Println("tab closed")

	return nil
}
