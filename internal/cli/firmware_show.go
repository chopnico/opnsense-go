package cli

import (
	"fmt"
	"strings"

	"github.com/chopnico/opnsense-go"

	"github.com/chopnico/output"
	"github.com/urfave/cli/v2"
)

func firmwareShowInfo(app *cli.App) *cli.Command {
	flags := globalFlags(nil)

	return &cli.Command{
		Name:  "info",
		Usage: "show firmware info",
		Flags: flags,
		Action: func(c *cli.Context) error {
			api := c.Context.Value("api").(opnsense.Api)
			info, err := api.CoreFirmwareInfo()
			if err != nil {
				return err
			}

			switch c.String("format") {
			case "json":
				fmt.Print(output.FormatItemAsJson(info))
			default:
				fmt.Print(output.FormatItemAsList(info, []string{"ProductName", "ProductVersion"}))
			}
			return nil
		},
	}
}

func firmwareShowStatus(app *cli.App) *cli.Command {
	flags := globalFlags(nil)

	return &cli.Command{
		Name:  "status",
		Usage: "show firmware status",
		Flags: flags,
		Action: func(c *cli.Context) error {
			api := c.Context.Value("api").(opnsense.Api)
			status, err := api.CoreFirmwareStatus()
			if err != nil {
				return err
			}

			switch c.String("format") {
			case "json":
				output.FormatItemAsJson(status)
			default:
				if c.String("properties") == "" {
					fmt.Print(output.FormatItemAsList(status, nil))
				} else {
					p := strings.Split(c.String("properties"), ",")
					fmt.Print(output.FormatItemAsList(status, p))
				}
			}
			return nil
		},
	}
}

func firmwareShowRunning(app *cli.App) *cli.Command {
	flags := globalFlags(nil)

	return &cli.Command{
		Name:  "running",
		Usage: "show if firmware is ready",
		Flags: flags,
		Action: func(c *cli.Context) error {
			api := c.Context.Value("api").(opnsense.Api)
			running, err := api.CoreFirmwareRunning()
			if err != nil {
				return err
			}

			switch c.String("format") {
			case "json":
				fmt.Print(output.FormatItemAsJson(running))
			default:
				if c.String("properties") == "" {
					fmt.Print(output.FormatItemAsList(running, nil))
				} else {
					p := strings.Split(c.String("properties"), ",")
					fmt.Print(output.FormatItemAsList(running, p))
				}
			}
			return nil
		},
	}
}
