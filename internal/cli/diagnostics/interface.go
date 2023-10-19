package diagnostics

import (
	"strconv"

	"github.com/chopnico/opnsense-go"
	oc "github.com/chopnico/opnsense-go/internal/cli"

	"github.com/urfave/cli/v2"
)

func interfaceArp(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:    "arp",
		Usage:   "list arp entries",
		Aliases: []string{"a"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "properties",
				Aliases: []string{"p"},
				Usage:   "show only these `PROPERTIES` (comma separated)",
			},
		},
		Action: func(c *cli.Context) error {
			arp, err := api.DiagnosticsInterfaceArp()
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				oc.PrintJson(arp)
			case "list":
				var l []interface{}
				for _, i := range *arp {
					l = append(l, i)
				}
				oc.PrintList(&l, c.String("properties"))
			default:
				data := [][]string{}
				for _, a := range *arp {
					data = append(data,
						[]string{a.Mac, a.IP, a.Intf, a.Manufacturer, a.Hostname},
					)
				}

				headers := []string{"MacAddress", "IPAddress", "Interface", "Manufacture", "Hostname"}
				oc.PrintTable(data, headers)
			}
			return nil
		},
	}
}

func interfaces(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:  "all",
		Usage: "list all interfaces",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "properties",
				Aliases: []string{"p"},
				Usage:   "show only these `PROPERTIES` (comma separated)",
			},
		},
		Action: func(c *cli.Context) error {
			interfaces, err := api.DiagnosticsInterfaceStatistics()
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				oc.PrintJson(interfaces)
			case "list":
				var l []interface{}
				for _, i := range *interfaces {
					l = append(l, i)
				}
				oc.PrintList(&l, c.String("properties"))
			default:
				data := [][]string{}
				for _, i := range *interfaces {
					data = append(data,
						[]string{i.Name, i.Interface, i.Address, i.Network, strconv.FormatInt(i.ReceivedBytes, 10), strconv.FormatInt(i.SentBytes, 10)},
					)
				}

				headers := []string{"InterfaceName", "Interface", "Address", "Network", "RecievedBytes", "SentBytes"}
				oc.PrintTable(data, headers)
			}
			return nil
		},
	}
}

func interfaceNdp(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:    "ndp",
		Usage:   "list all ndp interfaces",
		Aliases: []string{"n"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "properties",
				Aliases: []string{"p"},
				Usage:   "show only these`PROPERTIES` (comma separated)",
			},
		},
		Action: func(c *cli.Context) error {
			interfaceNdp, err := api.DiagnosticsInterfaceNdp()
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				oc.PrintJson(interfaceNdp)
			case "list":
				var l []interface{}
				for _, i := range *interfaceNdp {
					l = append(l, i)
				}
				oc.PrintList(&l, c.String("properties"))
			default:
				data := [][]string{}
				for _, i := range *interfaceNdp {
					data = append(data,
						[]string{i.Intf, i.IntfDescription, i.Mac, i.IP, i.Manufacturer},
					)
				}

				headers := []string{"Interface", "Description", "MacAddress", "IP", "Manufacture"}
				oc.PrintTable(data, headers)
			}
			return nil
		},
	}
}
