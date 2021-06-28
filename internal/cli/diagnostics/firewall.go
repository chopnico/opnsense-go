package diagnostics

import (
	"github.com/chopnico/opnsense"
	oc "github.com/chopnico/opnsense/internal/cli"

	"github.com/urfave/cli/v2"
)

func firewallLogs(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:    "logs",
		Aliases: []string{"l"},
		Usage:   "show firewall logs",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "limit",
				Usage: "limit the amount of logs returned",
				Value: "25",
			},
			&cli.StringFlag{
				Name:    "properties",
				Aliases: []string{"p"},
				Usage:   "show only these`PROPERTIES` (comma separated)",
			},
		},
		Action: func(c *cli.Context) error {
			log, err := api.DiagnosticsFirewallLog(c.String("limit"))
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				oc.PrintJson(log)
			case "list":
				var l []interface{}
				for _, i := range *log {
					l = append(l, i)
				}
				oc.PrintList(&l, c.String("properties"))
			default:
				data := [][]string{}
				for _, a := range *log {
					data = append(data,
						[]string{a.Src, a.Srcport, a.Dst, a.Dstport, a.Interface, a.Proto, a.Action},
					)
				}

				headers := []string{"Source", "SourcePort", "Destination", "DestinationPort", "Interface", "Protocol", "Action"}
				oc.PrintTable(data, headers)
			}
			return nil
		},
	}
}
