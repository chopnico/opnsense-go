package diagnostics

import (
	"github.com/chopnico/opnsense"

	"github.com/urfave/cli/v2"
)

func showCommands(app *cli.App, api *opnsense.Api) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		showInterfaceArp(app, api),
		showInterfaceBpfStatistics(app, api),
		showInterfaces(app, api),
	)

	return commands
}
