package diagnostics

import (
	"github.com/chopnico/opnsense-go"

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
					Name:    "interface",
					Usage:   "show interface diagnostics information",
					Aliases: []string{"i"},
					Subcommands: []*cli.Command{
						{
							Name:        "list",
							Usage:       "list interface stuff",
							Aliases:     []string{"l"},
							Subcommands: listInterfaceCommands(app, api),
						},
						{
							Name:        "show",
							Usage:       "show interface stuff",
							Aliases:     []string{"s"},
							Subcommands: showInterfaceCommands(app, api),
						},
					},
				},
				{
					Name:        "firewall",
					Usage:       "show firewall diagnostics information",
					Aliases:     []string{"f"},
					Subcommands: firewallCommands(app, api),
				},
			},
		},
	)
}

func showInterfaceCommands(app *cli.App, api *opnsense.Api) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		interfaceStatistics(app, api),
	)

	return commands
}

func listInterfaceCommands(app *cli.App, api *opnsense.Api) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		interfaces(app, api),
		interfaceArp(app, api),
		interfaceNdp(app, api),
		interfaceRoutes(app, api),
	)

	return commands
}

func firewallCommands(app *cli.App, api *opnsense.Api) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		firewallLogs(app, api),
	)

	return commands
}
