package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chopnico/opnsense-go"

	"github.com/chopnico/output"
	"github.com/urfave/cli/v2"
)

func interfacesShowStatistics(app *cli.App) *cli.Command {
	flags := globalFlags(nil)

	return &cli.Command{
		Name:    "statistics",
		Aliases: []string{"s"},
		Usage:   "show interface statistics",
		Flags:   flags,
		Action: func(c *cli.Context) error {
			api := c.Context.Value("api").(opnsense.Api)
			statistics, err := api.DiagnosticsInterfaceStatistics()
			if err != nil {
				return err
			}

			switch c.String("format") {
			case "json":
				fmt.Print(output.FormatItemsAsJson(statistics))
			case "list":
				if c.String("properties") == "" {
					fmt.Print(output.FormatItemsAsList(statistics, nil))
				} else {
					p := strings.Split(c.String("properties"), ",")
					fmt.Print(output.FormatItemsAsList(statistics, p))
				}
			default:
				data := [][]string{}
				for _, i := range *statistics {
					data = append(data,
						[]string{
							i.Interface,
							strconv.FormatInt(i.ReceivedBytes, 10),
							strconv.FormatInt(i.SentBytes, 10),
							strconv.Itoa(i.ReceivedErrors),
							strconv.Itoa(i.SendErrors),
							strconv.Itoa(i.DroppedPackets),
						},
					)
				}

				headers := []string{"Interface", "Recieved Bytes", "Sent Bytes", "Recieved Errors", "Sent Errors", "Dropped Packets"}
				fmt.Print(output.FormatTable(data, headers))
			}

			return nil
		},
	}
}

func interfacesShowBpfStatistics(app *cli.App) *cli.Command {
	flags := globalFlags(nil)

	return &cli.Command{
		Name:    "bfp-statistics",
		Aliases: []string{"b"},
		Usage:   "show bfp statistics",
		Flags:   flags,
		Action: func(c *cli.Context) error {
			api := c.Context.Value("api").(opnsense.Api)
			bpfStatistics, err := api.DiagnosticsInterfaceBpfStatistics()
			if err != nil {
				return err
			}

			switch c.String("format") {
			case "json":
				fmt.Print(output.FormatItemsAsJson(bpfStatistics))
			case "list":
				if c.String("properties") == "" {
					fmt.Print(output.FormatItemsAsList(bpfStatistics, nil))
				} else {
					p := strings.Split(c.String("properties"), ",")
					fmt.Print(output.FormatItemsAsList(bpfStatistics, p))
				}
			default:
				data := [][]string{}
				for _, i := range *bpfStatistics {
					data = append(data,
						[]string{
							i.InterfaceName,
							strconv.Itoa(i.ReceivedPackets),
							strconv.Itoa(i.DroppedPackets),
							i.Direction,
						},
					)
				}

				headers := []string{"Interface", "Recieved Packets", "Dropped Packets", "Direction"}
				fmt.Print(output.FormatTable(data, headers))
			}

			return nil
		},
	}
}

func interfacesShowArp(app *cli.App) *cli.Command {
	flags := globalFlags(nil)

	return &cli.Command{
		Name:    "arp",
		Aliases: []string{"a"},
		Usage:   "show arp",
		Flags:   flags,
		Action: func(c *cli.Context) error {
			api := c.Context.Value("api").(opnsense.Api)
			arp, err := api.DiagnosticsInterfaceArp()
			if err != nil {
				return err
			}

			switch c.String("format") {
			case "json":
				fmt.Print(output.FormatItemsAsJson(arp))
			case "list":
				if c.String("properties") == "" {
					fmt.Print(output.FormatItemsAsList(arp, nil))
				} else {
					p := strings.Split(c.String("properties"), ",")
					fmt.Print(output.FormatItemsAsList(arp, p))
				}
			default:
				data := [][]string{}
				for _, i := range *arp {
					data = append(data,
						[]string{i.Intf, i.Mac, i.IP, i.IntfDescription, i.Hostname, i.Manufacturer},
					)
				}

				headers := []string{"Interface", "MAC", "IP", "Interface Description", "Hostname", "Manufacturer"}
				fmt.Print(output.FormatTable(data, headers))
			}

			return nil
		},
	}
}

func interfacesShowRoutes(app *cli.App) *cli.Command {
	flags := globalFlags(nil)

	return &cli.Command{
		Name:    "routes",
		Aliases: []string{"r"},
		Usage:   "show interface routes",
		Flags:   flags,
		Action: func(c *cli.Context) error {
			api := c.Context.Value("api").(opnsense.Api)
			routes, err := api.DiagnosticsInterfaceRoutes()
			if err != nil {
				return err
			}

			switch c.String("format") {
			case "json":
				fmt.Print(output.FormatItemsAsJson(routes))
			case "list":
				if c.String("properties") == "" {
					fmt.Print(output.FormatItemsAsList(routes, nil))
				} else {
					p := strings.Split(c.String("properties"), ",")
					fmt.Print(output.FormatItemsAsList(routes, p))
				}
			default:
				data := [][]string{}
				for _, i := range *routes {
					data = append(data,
						[]string{i.Netif, i.Destination, i.Gateway, i.IntfDescription},
					)
				}

				headers := []string{"Interface", "Destination", "Gateway", "Description"}
				fmt.Print(output.FormatTable(data, headers))
			}

			return nil
		},
	}
}
