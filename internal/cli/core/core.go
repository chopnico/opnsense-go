package core

import (
	"github.com/chopnico/opnsense-go"

	"github.com/urfave/cli/v2"
)

func NewCommand(app *cli.App, api *opnsense.Api) {
	app.Commands = append(app.Commands,
		&cli.Command{
			Name:        "firmware",
			Aliases:     []string{"f"},
			Usage:       "firewall firmware",
			Subcommands: firmwareCommands(app, api),
		},
	)
}

func firmwareCommands(app *cli.App, api *opnsense.Api) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		firmwareInfo(app, api),
		firmwareStatus(app, api),
		firmwareRunning(app, api),
	)

	return commands
}
