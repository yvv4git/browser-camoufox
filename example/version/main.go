// Command version asks the Camoufox browser running in the container to
// report its own fingerprint: user agent, app version, platform, vendor
// and hardware concurrency. The data comes from the browser itself via
// JavaScript (navigator.*), not from files inside the container.
//
// Run from the example directory:
//
//	go run ./version -addr http://localhost:9377
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/yvv4git/go-juggler"
)

const _fingerprintExpr = `JSON.stringify({
  userAgent: navigator.userAgent,
  appVersion: navigator.appVersion,
  platform: navigator.platform,
  vendor: navigator.vendor,
  hardwareConcurrency: navigator.hardwareConcurrency
})`

func main() {
	addr := flag.String("addr", "http://localhost:9377", "HTTP endpoint of the running camofox-browser")

	flag.Parse()

	if err := run(*addr); err != nil {
		log.Fatal(err)
	}
}

func run(addr string) error {
	b := newBrowser(addr)

	info, err := b.fingerprint(context.Background())
	if err != nil {
		return err
	}

	keys := []string{"userAgent", "appVersion", "platform", "vendor", "hardwareConcurrency"}

	for _, key := range keys {
		fmt.Printf("%-22s %v\n", key, info[key])
	}

	return nil
}

type browser struct {
	c *juggler.Client
}

func newBrowser(addr string) *browser {
	return &browser{c: juggler.NewClient(addr)}
}

func (b *browser) fingerprint(ctx context.Context) (map[string]any, error) {
	const session = "version-demo"

	tab, err := b.c.OpenTab(ctx, session, "https://example.com")
	if err != nil {
		return nil, fmt.Errorf("open tab: %w", err)
	}

	defer func() { _ = b.c.CloseTab(ctx, tab.TabID, session) }()

	res, err := b.c.Evaluate(ctx, tab.TabID, session, _fingerprintExpr)
	if err != nil {
		return nil, fmt.Errorf("evaluate: %w", err)
	}

	raw, ok := res.Result.(string)
	if !ok {
		return nil, fmt.Errorf("unexpected result type: %T", res.Result)
	}

	var info map[string]any

	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return info, nil
}
