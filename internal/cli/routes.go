package cli

import (
	"fmt"
	"strings"

	"github.com/chopnico/opnsense-go"

	"github.com/chopnico/output"
	"github.com/urfave/cli/v2"
)

func deleteRoute(app *cli.App) *cli.Command {
	flags := []cli.Flag{}
	flags = append(flags,
		&cli.StringFlag{
			Name:     "uuid",
			Aliases:  []string{"u"},
			Usage:    "the `UUID` of the route",
			Required: true,
		},
	)

	return &cli.Command{
		Name:    "delete",
		Usage:   "delete a single route",
		Aliases: []string{"d"},
		Flags:   flags,
		Action: func(c *cli.Context) error {
			api := c.Context.Value("api").(opnsense.Api)
			err := api.DeleteRoute(c.String("uuid"))
			if err != nil {
				return err
			}

			fmt.Println("route with uuid " + c.String("uuid") + "has been deleted")

			return nil
		},
	}
}

func setRoute(app *cli.App) *cli.Command {
	flags := []cli.Flag{}
	flags = append(flags,
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
	)

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

func getRoute(app *cli.App) *cli.Command {
	flags := globalFlags()
	flags = append(flags,
		&cli.StringFlag{
			Name:     "uuid",
			Aliases:  []string{"u"},
			Usage:    "the `UUID` of the route",
			Required: true,
		},
	)

	return &cli.Command{
		Name:    "get",
		Usage:   "get a single route",
		Aliases: []string{"g"},
		Flags:   flags,
		Action: func(c *cli.Context) error {
			api := c.Context.Value("api").(opnsense.Api)
			route, err := api.GetRouteByUuid(c.String("uuid"))
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
			return nil
		},
	}
}

func listRoutes(app *cli.App) *cli.Command {
	flags := globalFlags()

	return &cli.Command{
		Name:    "list",
		Usage:   "list all routes",
		Aliases: []string{"l"},
		Flags:   flags,
		Action: func(c *cli.Context) error {
			api := c.Context.Value("api").(opnsense.Api)
			routes, err := api.GetRoutes()
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
						[]string{i.UUID, i.Network, i.Gateway, i.Description, i.Disabled},
					)
				}

				headers := []string{"UUID", "Network", "Gateway", "Descriptioniption", "Disabled"}
				fmt.Print(output.FormatTable(data, headers))
			}
			return nil
		},
	}
}
