package firmware

import (
	"github.com/chopnico/opnsense"
	oc "github.com/chopnico/opnsense/internal/cli"

	"github.com/urfave/cli/v2"
)

func showInfo(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:  "info",
		Usage: "show firmware info",
		Action: func(c *cli.Context) error {
			info, err := api.FirmwareInfo()
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				oc.PrintJson(info)
			default:
				var l []interface{}
				l = append(l, info)
				oc.PrintList(&l, c.String("properties"))
			}
			return nil
		},
	}
}

func showStatus(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:  "status",
		Usage: "show firmware status",
		Action: func(c *cli.Context) error {
			status, err := api.FirmwareStatus()
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				oc.PrintJson(status)
			default:
				var l []interface{}
				l = append(l, status)
				oc.PrintList(&l, c.String("properties"))
			}
			return nil
		},
	}
}

func showCommands(app *cli.App, api *opnsense.Api) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		showInfo(app, api),
		showStatus(app, api),
	)

	return commands
}
