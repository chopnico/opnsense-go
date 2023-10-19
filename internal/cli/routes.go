package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/chopnico/opnsense-go"

	"github.com/chopnico/output"
	"github.com/urfave/cli/v2"
)

func routesDelete(app *cli.App) *cli.Command {
	flags := []cli.Flag{}

	return &cli.Command{
		Name:      "delete",
		Usage:     "delete route",
		Aliases:   []string{"d"},
		ArgsUsage: "UUID",
		Flags:     flags,
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				cli.ShowCommandHelp(c, "delete")
				return errors.New("you must supply a route")
			} else {
				for i := 0; i < c.Args().Len(); i++ {
					api := c.Context.Value("api").(opnsense.Api)
					err := api.DeleteRoute(c.Args().Get(i))
					if err != nil {
						return err
					}

					fmt.Println("route with uuid " + c.Args().Get(i) + " has been deleted")
				}
				return nil
			}
		},
	}
}

func routesSet(app *cli.App) *cli.Command {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "destination",
			Aliases:  []string{"n"},
			Usage:    "the `DESTINATION` address",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "gateway",
			Aliases:  []string{"g"},
			Usage:    "the `GATEWAY` address",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "description",
			Aliases:  []string{"d"},
			Usage:    "a `DESCRIPTION` for the route",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "disable",
			Usage:    "should we `DISABLE` the route",
			Value:    false,
			Required: false,
		},
	}

	return &cli.Command{
		Name:    "set",
		Usage:   "create a new route",
		Aliases: []string{"s"},
		Flags:   flags,
		Action: func(c *cli.Context) error {
			api := c.Context.Value("api").(opnsense.Api)
			route := opnsense.Route{
				Network:     c.String("destination"),
				Gateway:     c.String("gateway"),
				Description: c.String("description"),
			}

			if c.Bool("disable") {
				route.Disabled = "1"
			} else {
				route.Disabled = "0"
			}

			r, err := api.SetRoute(route)
			if err != nil {
				return err
			}

			switch c.String("format") {
			case "json":
				fmt.Print(output.FormatItemAsJson(r))
			default:
				if c.String("properties") == "" {
					fmt.Print(output.FormatItemAsList(r, nil))
				} else {
					p := strings.Split(c.String("properties"), ",")
					fmt.Print(output.FormatItemAsList(r, p))
				}
			}
			return nil
		},
	}
}

func routesGet(app *cli.App) *cli.Command {
	flags := globalFlags(nil)

	return &cli.Command{
		Name:      "get",
		Usage:     "get a route",
		ArgsUsage: "UUID",
		Aliases:   []string{"g"},
		Flags:     flags,
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				cli.ShowCommandHelp(c, "get")
				return errors.New("you must supply a uuid of a route")
			} else {
				for i := 0; i < c.Args().Len(); i++ {
					api := c.Context.Value("api").(opnsense.Api)
					route, err := api.GetRouteByUuid(c.Args().Get(i))
					if err != nil {
						return err
					}

					switch c.String("format") {
					case "json":
						fmt.Print(output.FormatItemAsJson(route))
					default:
						if c.String("properties") == "" {
							fmt.Print(output.FormatItemAsList(route, nil))
						} else {
							p := strings.Split(c.String("properties"), ",")
							fmt.Print(output.FormatItemAsList(route, p))
						}
					}
				}
				return nil
			}

		},
	}
}

func routesList(app *cli.App) *cli.Command {
	flags := addQuietFlag(globalFlags(
		[]cli.Flag{
			&cli.StringFlag{
				Name:     "network",
				Aliases:  []string{"n"},
				Usage:    "list only routes destined for `NETWORK`",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "description",
				Aliases:  []string{"d"},
				Usage:    "list only routes with `DESCRIPTION`",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "enabled",
				Aliases:  []string{"e"},
				Usage:    "list only enabled routes",
				Value:    false,
				Required: false,
			},
		},
	))

	return &cli.Command{
		Name:    "list",
		Usage:   "list all routes",
		Aliases: []string{"l"},
		Flags:   flags,
		Action: func(c *cli.Context) error {
			api := c.Context.Value("api").(opnsense.Api)
			var (
				routes *[]opnsense.Route
				err    error
			)
			if c.String("network") != "" && c.String("description") != "" {
				routes, err = api.GetRoutesByNetworkWithDescription(
					c.String("network"),
					c.String("description"),
					c.Bool("enabled"),
				)
			} else if c.String("network") != "" {
				routes, err = api.GetRoutesByNetwork(c.String("network"), c.Bool("enabled"))
			} else if c.String("description") != "" {
				routes, err = api.GetRoutesByDescription(c.String("description"), c.Bool("enabled"))
			} else {
				routes, err = api.GetRoutes(c.Bool("enabled"))
			}
			if err != nil {
				return err
			}

			if c.Bool("quiet") {
				for _, r := range *routes {
					fmt.Println(r.UUID)
				}
			} else {
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
							[]string{i.UUID, i.Network, i.Gateway, i.Description, i.Disabled},
						)
					}

					headers := []string{"UUID", "Network", "Gateway", "Description", "Disabled"}
					fmt.Print(output.FormatTable(data, headers))
				}
			}
			return nil
		},
	}
}
