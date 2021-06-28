package firewall

import (
	"github.com/chopnico/opnsense"

	"github.com/urfave/cli/v2"
)

func NewCommand(app *cli.App, api *opnsense.Api) {
	app.Commands = append(app.Commands,
		&cli.Command{
			Name:    "firewall",
			Aliases: []string{"fw"},
			Usage:   "firewall",
			Subcommands: []*cli.Command{
				{
					Name:        "alias",
					Usage:       "firewall alias",
					Aliases:     []string{"a"},
					Subcommands: firewallAliasCommands(app, api),
				},
			},
		},
	)
}

func firewallAliasCommands(app *cli.App, api *opnsense.Api) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		firewallAlias(app, api),
	)

	return commands
}
