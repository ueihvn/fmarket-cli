package cliapp

import (
	"context"

	"github.com/rivo/tview"
	"github.com/ueihvn/fmarket-cli/pkg/fmarketclient"
)

type CLIApp interface {
	Start() error
	Stop()
}

type cliApp struct {
	app *tview.Application
	fc  fmarketclient.FmarketClient
}

func MustNewCliApp() *cliApp {
	app, err := NewCliApp()
	if err != nil {
		panic(err)
	}
	return app
}

func NewCliApp() (*cliApp, error) {
	fc, err := fmarketclient.NewFmarketClient()
	if err != nil {
		return nil, err
	}
	return &cliApp{
		app: tview.NewApplication(),
		fc:  fc,
	}, nil
}

func (c *cliApp) Start() error {
	if err := c.setupApp(); err != nil {
		return err
	}
	return c.app.Run()
}

func (c *cliApp) Stop() {
	c.app.Stop()
}

func (c *cliApp) setupApp() error {
	ctx := context.Background()
	tableFundListing, err := c.TableFundListing(ctx)
	if err != nil {
		return err
	}
	c.app.SetRoot(tableFundListing, true).SetFocus(tableFundListing)
	return nil
}
