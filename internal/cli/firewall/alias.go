package firewall

import (
	"github.com/chopnico/opnsense-go"
	oc "github.com/chopnico/opnsense-go/internal/cli"

	"github.com/urfave/cli/v2"
)

func firewallAlias(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:    "export",
		Usage:   "list aliases",
		Aliases: []string{"e"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "properties",
				Aliases: []string{"p"},
				Usage:   "show only these `PROPERTIES` (comma separated)",
			},
		},
		Action: func(c *cli.Context) error {
			aliases, err := api.FirewallAlias()
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				oc.PrintJson(aliases)
			case "list":
				var l []interface{}
				for _, i := range *aliases {
					l = append(l, i)
				}
				oc.PrintList(&l, c.String("properties"))
			default:
				data := [][]string{}
				for _, a := range *aliases {
					var enabled string

					if a.Enabled == "1" {
						enabled = "true"
					} else {
						enabled = "false"
					}

					data = append(data,
						[]string{enabled, a.Id, a.Name, a.Description},
					)
				}

				headers := []string{"Enabled", "Id", "Name", "Description"}
				oc.PrintTable(data, headers)
			}
			return nil
		},
	}
}
