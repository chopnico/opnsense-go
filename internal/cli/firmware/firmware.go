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
					Name:    "show",
					Usage:   "show firewall information",
					Aliases: []string{"s"},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "properties",
							Aliases:  []string{"p"},
							Usage:    "Show only these`PROPERTIES` (comma separated)",
							Required: false,
						},
					},
					Subcommands: showCommands(app, api),
				},
			},
		},
	)
}
