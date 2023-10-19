package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chopnico/opnsense-go"

	"github.com/chopnico/output"
	"github.com/urfave/cli/v2"
)

func interfacesList(app *cli.App) *cli.Command {
	flags := globalFlags(nil)

	return &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list interfaces",
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
							i.Address,
							i.Network,
							strconv.Itoa(i.Mtu),
						},
					)
				}

				headers := []string{"Interface", "Address", "Network", "MTU"}
				fmt.Print(output.FormatTable(data, headers))
			}

			return nil
		},
	}
}
