package diagnostics

import (
	"strconv"

	"github.com/chopnico/opnsense-go"
	oc "github.com/chopnico/opnsense-go/internal/cli"

	"github.com/urfave/cli/v2"
)

func interfaceMemoryStatistics(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:    "memory",
		Aliases: []string{"m"},
		Usage:   "show interface memory statistics",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "properties",
				Aliases: []string{"p"},
				Usage:   "show only these`PROPERTIES` (comma separated)",
			},
		},
		Action: func(c *cli.Context) error {
			interfaceMemoryStatistics, err := api.DiagnosticsInterfaceMemoryStatistics()
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				oc.PrintJson(interfaceMemoryStatistics)
			default:
				var l []interface{}
				l = append(l, interfaceMemoryStatistics)
				oc.PrintList(&l, c.String("properties"))
			}
			return nil
		},
	}
}

func interfaceBpfStatistics(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:    "bpf",
		Usage:   "list bpf statistics",
		Aliases: []string{"b"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "properties",
				Aliases: []string{"p"},
				Usage:   "show only these`PROPERTIES` (comma separated)",
			},
		},
		Action: func(c *cli.Context) error {
			bpfStasticsEntries, err := api.DiagnosticsInterfaceBpfStatistics()
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				oc.PrintJson(bpfStasticsEntries)
			case "list":
				var l []interface{}
				for _, i := range *bpfStasticsEntries {
					l = append(l, i)
				}
				oc.PrintList(&l, c.String("properties"))
			default:
				data := [][]string{}
				for _, a := range *bpfStasticsEntries {
					data = append(data,
						[]string{a.InterfaceName, a.Direction, strconv.Itoa(a.ReceivedPackets), a.Process},
					)
				}

				headers := []string{"InterfaceName", "Direction", "ReceivedPackets", "Process"}
				oc.PrintTable(data, headers)
			}
			return nil
		},
	}
}

func interfaceStatistics(app *cli.App, api *opnsense.Api) *cli.Command {
	var commands []*cli.Command

	commands = append(commands,
		interfaceMemoryStatistics(app, api),
		interfaceBpfStatistics(app, api),
	)

	return &cli.Command{
		Name:        "statistics",
		Aliases:     []string{"s"},
		Usage:       "show diagnostics statistics",
		Subcommands: commands,
	}
}
