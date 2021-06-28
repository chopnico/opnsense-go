package main

import (
	"errors"
	"log"
	"os"

	"github.com/chopnico/opnsense-go"

	"github.com/chopnico/opnsense-go/internal/cli/core"
	"github.com/chopnico/opnsense-go/internal/cli/diagnostics"
	"github.com/chopnico/opnsense-go/internal/cli/firewall"

	"github.com/urfave/cli/v2"
)

func main() {
	var api opnsense.Api

	app := cli.NewApp()
	app.Name = "opnsense"
	app.Usage = "opnsense CLI"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "key",
			Usage:   "account `KEY`",
			EnvVars: []string{"OPNSENSE_KEY"},
		},
		&cli.StringFlag{
			Name:    "secret",
			Usage:   "account `SECRET`",
			EnvVars: []string{"OPNSENSE_SECRET"},
		},
		&cli.StringFlag{
			Name:    "host",
			Usage:   "firewall `HOST` (firewall.example.local)",
			EnvVars: []string{"OPNSENSE_HOST"},
		},
		&cli.BoolFlag{
			Name:  "ignore-ssl",
			Usage: "ignore ssl errors",
			Value: false,
		},
		&cli.IntFlag{
			Name:  "timeout",
			Usage: "http timeout",
			Value: 0,
		},
		&cli.StringFlag{
			Name:  "format",
			Usage: "printing format (json, list, table)",
			Value: "table",
		},
		&cli.StringFlag{
			Name:  "logging",
			Usage: "set logging level",
			Value: "info",
		},
		&cli.StringFlag{
			Name:  "proxy",
			Usage: "set http proxy",
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.String("key") == "" {
			return errors.New(opnsense.ErrorEmptyUsername)
		} else if c.String("secret") == "" {
			return errors.New(opnsense.ErrorEmptyPassword)
		} else if c.String("host") == "" {
			return errors.New(opnsense.ErrorEmptyHost)
		}

		options := opnsense.ApiOptions{
			IgnoreSslErrors: c.Bool("ignore-ssl"),
			TimeOut:         c.Int("timeout"),
			Logging:         c.String("logging"),
			Proxy:           c.String("proxy"),
			Print:           c.String("format"),
		}

		a, err := opnsense.NewApiBasicAuth(
			c.String("key"),
			c.String("secret"),
			c.String("host"),
			&options,
		)
		if err != nil {
			return err
		}

		api = (*a)

		return nil
	}

	core.NewCommand(app, &api)
	diagnostics.NewCommand(app, &api)
	firewall.NewCommand(app, &api)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
