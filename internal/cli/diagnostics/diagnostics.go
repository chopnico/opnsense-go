package diagnostics

import (
	"github.com/chopnico/opnsense"

	"github.com/urfave/cli/v2"
)

func NewCommand(app *cli.App, api *opnsense.Api) {
	app.Commands = append(app.Commands,
		&cli.Command{
			Name:    "diagnostics",
			Aliases: []string{"d"},
			Usage:   "firewall diagnostics",
			Subcommands: []*cli.Command{
				{
					Name:        "show",
					Usage:       "show firewall diagnostics information",
					Aliases:     []string{"s"},
					Subcommands: showCommands(app, api),
				},
			},
		},
	)
}
