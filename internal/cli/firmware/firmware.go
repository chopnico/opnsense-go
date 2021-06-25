package firmware

import (
	"github.com/chopnico/opnsense"

	"github.com/urfave/cli/v2"
)

func NewCommand(app *cli.App, api *opnsense.Api) {
	app.Commands = append(app.Commands,
		&cli.Command{
			Name:    "firmware",
			Aliases: []string{"f"},
			Usage:   "firewall firmware",
			Subcommands: []*cli.Command{
				{
					Name:        "show",
					Usage:       "show firewall information",
					Aliases:     []string{"s"},
					Subcommands: showCommands(app, api),
				},
			},
		},
	)
}
