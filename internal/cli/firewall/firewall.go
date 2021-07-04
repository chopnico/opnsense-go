package firewall

import (
	"github.com/chopnico/opnsense-go"

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
					Name:    "alias",
					Usage:   "list firewall stuff",
					Aliases: []string{"l"},
					Subcommands: []*cli.Command{
						{
							Name:        "list",
							Usage:       "firewall alias",
							Aliases:     []string{"a"},
							Subcommands: listFirewallAliasCommands(app, api),
						},
					},
				},
			},
		},
	)
}

func listFirewallAliasCommands(app *cli.App, api *opnsense.Api) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		firewallAlias(app, api),
	)

	return commands
}
