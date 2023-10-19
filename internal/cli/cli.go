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
		&cli.Command{
			Name:        "interfaces",
			Aliases:     []string{"i"},
			Usage:       "interact with interfaces",
			Subcommands: interfacesCommands(app),
		},
	)
}

func globalFlags(flags []cli.Flag) []cli.Flag {
	flags = append(flags,
		&cli.StringFlag{
			Name:     "properties",
			Aliases:  []string{"p"},
			Usage:    "`PROPERTIES` to print (comma seperated list)",
			Required: false,
		},
	)
	return flags
}

func addQuietFlag(flags []cli.Flag) []cli.Flag {
	flags = append(flags,
		&cli.BoolFlag{
			Name:     "quiet",
			Aliases:  []string{"q"},
			Usage:    "show ids only",
			Required: false,
		},
	)
	return flags
}

func firmwareCommands(app *cli.App) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		&cli.Command{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "show firmware information",
			Subcommands: []*cli.Command{
				firmwareShowInfo(app),
				firmwareShowStatus(app),
				firmwareShowRunning(app),
			},
		},
	)

	return commands
}

func routesCommands(app *cli.App) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		routesList(app),
		routesGet(app),
		routesSet(app),
		routesDelete(app),
	)

	return commands
}

func interfacesCommands(app *cli.App) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		&cli.Command{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "show interface information",
			Subcommands: []*cli.Command{
				interfacesShowRoutes(app),
				interfacesShowArp(app),
				interfacesShowBpfStatistics(app),
				interfacesShowStatistics(app),
			},
		},
		interfacesList(app),
	)

	return commands
}
