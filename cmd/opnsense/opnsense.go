package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/chopnico/opnsense-go"
	CLI "github.com/chopnico/opnsense-go/internal/cli"

	"github.com/urfave/cli/v2"
)

// some application and default variables
var (
	AppName  string = "flare"
	AppUsage string = "a cloudflare cli/tui tool"
	// ldflags will be used to set this. check Makefile
	AppVersion string

	DefaultLoggingLevel = "info"
	DefaultPrintFormat  = "table"
	DefaultTimeOut      = 60
)

func main() {
	app := cli.NewApp()
	app.Name = "opnsense"
	app.Usage = "opnsense CLI"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "key",
			Usage:       "account `KEY`",
			EnvVars:     []string{"OPNSENSE_KEY"},
			DefaultText: "none",
		},
		&cli.StringFlag{
			Name:        "secret",
			Usage:       "account `SECRET`",
			EnvVars:     []string{"OPNSENSE_SECRET"},
			DefaultText: "none",
		},
		&cli.StringFlag{
			Name:        "host",
			Usage:       "firewall `HOST` (firewall.example.local)",
			EnvVars:     []string{"OPNSENSE_HOST"},
			DefaultText: "none",
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
		var err error
		var api opnsense.Api

		if c.String("key") == "" {
			cli.ShowAppHelp(c)
			return errors.New(opnsense.ErrorEmptyKey)
		} else if c.String("secret") == "" {
			cli.ShowAppHelp(c)
			return errors.New(opnsense.ErrorEmptySecret)
		} else if c.String("host") == "" {
			cli.ShowAppHelp(c)
			return errors.New(opnsense.ErrorEmptyHost)
		}

		// create new api client with basic auth
		api, err = opnsense.NewApiBasicAuth(
			c.String("key"),
			c.String("secret"),
			c.String("host"),
		)
		if err != nil {
			return err
		}

		// set options
		api.Timeout(c.Int("timeout")).
			LoggingLevel(c.String("logging")).
			Proxy(c.String("proxy"))

		// should we ignore ssl errors?
		if c.Bool("ignore-ssl") {
			api.IgnoreSslErrors()
		}

		// add the api client to the cli context so that it can be used
		// throughout the application. call c.Context.Value("api").(opnsense.Api)
		// in an action to retrieve the api client
		ctx := context.WithValue(c.Context, "api", api)
		c.Context = ctx

		return nil
	}

	// create cli commands
	CLI.NewCommands(app)

	// run the application
	err := app.Run(os.Args)
	if err != nil {
		if err.Error() != "debugging" {
			log.Fatal(err)
		}
	}

	os.Exit(0)
}
