package firmware

import (
	"github.com/chopnico/opnsense"
	oc "github.com/chopnico/opnsense/internal/cli"

	"github.com/urfave/cli/v2"
)

func info(app *cli.App, api *opnsense.Api) *cli.Command {
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
				err := oc.PrintJson(info)
				if err != nil {
					return err
				}
				return nil
			default:
				data := [][]string{
					[]string{info.ProductName, info.ProductVersion},
				}

				headers := []string{"ProductName", "ProductVersion"}
				oc.PrintTable(data, headers)
			}
			return nil
		},
	}
}

func status(app *cli.App, api *opnsense.Api) *cli.Command {
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
				err := oc.PrintJson(status)
				if err != nil {
					return err
				}
				return nil
			default:
				var update string
				if status.Updates != "0" {
					update = "true"
				} else {
					update = "false"
				}

				data := [][]string{
					[]string{status.Connection, status.OsVersion, status.ProductVersion, update},
				}

				headers := []string{"Connection", "OsVersion", "ProductVersion", "Update?"}
				oc.PrintTable(data, headers)
			}
			return nil
		},
	}
}

func showCommands(app *cli.App, api *opnsense.Api) []*cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		info(app, api),
		status(app, api),
	)

	return commands
}
