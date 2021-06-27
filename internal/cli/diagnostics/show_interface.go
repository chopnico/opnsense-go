package diagnostics

import (
	"strconv"

	"github.com/chopnico/opnsense"
	oc "github.com/chopnico/opnsense/internal/cli"

	"github.com/urfave/cli/v2"
)

func showInterfaceArp(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:  "arp",
		Usage: "show firmware info",
		Action: func(c *cli.Context) error {
			arp, err := api.DiagnosticsInterfaceArp()
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				err := oc.PrintJson(arp)
				if err != nil {
					return err
				}
				return nil
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

func showInterfaceBpfStatistics(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:  "bpf-statistics",
		Usage: "show bpf statistics",
		Action: func(c *cli.Context) error {
			bpfStastics, err := api.DiagnosticsInterfaceBpfStatistics()
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				err := oc.PrintJson(bpfStastics)
				if err != nil {
					return err
				}
				return nil
			default:
				data := [][]string{}
				for _, a := range *&bpfStastics.BpfStatistics.BpfEntry {
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

func showInterfaces(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:  "interfaces",
		Usage: "show interfaces",
		Action: func(c *cli.Context) error {
			interfaces, err := api.DiagnosticsInterfaceStatistics()
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				err := oc.PrintJson(interfaces)
				if err != nil {
					return err
				}
				return nil
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

func showMemoryStatistics(app *cli.App, api *opnsense.Api) *cli.Command {
	return &cli.Command{
		Name:  "mbuf-statistics",
		Usage: "show interface memory statistics",
		Action: func(c *cli.Context) error {
			interfaces, err := api.DiagnosticsInterfaceMemoryStatistics()
			if err != nil {
				return err
			}

			switch api.Options.Print {
			case "json":
				err := oc.PrintJson(interfaces)
				if err != nil {
					return err
				}
				return nil
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
