package core

import (
	"github.com/chopnico/opnsense-go"
	oc "github.com/chopnico/opnsense-go/internal/cli"

	"github.com/urfave/cli/v2"
)

func firmwareInfo(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:  "info",
		Usage: "show firmware info",
		Action: func(c *cli.Context) error {
			info, err := api.CoreFirmwareInfo()
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

func firmwareStatus(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:  "status",
		Usage: "show firmware status",
		Action: func(c *cli.Context) error {
			status, err := api.CoreFirmwareStatus()
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

func firmwareRunning(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:  "running",
		Usage: "show if firmware is ready",
		Action: func(c *cli.Context) error {
			running, err := api.CoreFirmwareRunning()
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				oc.PrintJson(running)
			default:
				var l []interface{}
				l = append(l, running)
				oc.PrintList(&l, c.String("properties"))
			}
			return nil
		},
	}
}
