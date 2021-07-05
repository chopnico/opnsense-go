package cli

import (
	"github.com/urfave/cli/v2"
)

func NewCommands(app *cli.App) {
	app.Commands = append(app.Commands,
		&cli.Command{
			Name:    "system",
			Aliases: []string{"s"},
			Usage:   "interact with system related stuff",
			Subcommands: []*cli.Command{
				&cli.Command{
					Name:        "firmware",
					Aliases:     []string{"f"},
					Usage:       "interact with the system firmware",
					Subcommands: firmwareCommands(app),
				},
			},
		},
		&cli.Command{
			Name:        "routes",
			Aliases:     []string{"r"},
			Usage:       "interact with routes",
			Subcommands: routesCommands(app),
		},
	)
}

func globalFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "properties",
			Aliases:  []string{"p"},
			Usage:    "`PROPERTIES` to print (only relevant to list format)",
			Required: false,
		},
	}
}

func firmwareCommands(app *cli.App) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		firmwareInfo(app),
		firmwareStatus(app),
		firmwareRunning(app),
	)

	return commands
}

func routesCommands(app *cli.App) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		listRoutes(app),
		getRoute(app),
		setRoute(app),
		deleteRoute(app),
	)

	return commands
}
