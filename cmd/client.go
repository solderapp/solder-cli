package cmd

import (
	"github.com/codegangsta/cli"
)

// Client provides the sub-command for the client API.
func Client() cli.Command {
	return cli.Command{
		Name:    "client",
		Aliases: []string{"c"},
		Usage:   "Client related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all clients",
				Action: func(c *cli.Context) {
					Handle(c, ClientList)
				},
			},
			{
				Name:  "show",
				Usage: "Display a build",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ClientShow)
				},
			},
			{
				Name:  "update",
				Usage: "Update a client",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ClientUpdate)
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"rm"},
				Usage:   "Delete a client",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ClientDelete)
				},
			},
			{
				Name:  "create",
				Usage: "Create a client",
				Action: func(c *cli.Context) {
					Handle(c, ClientCreate)
				},
			},
		},
	}
}

// ClientList provides the sub-command to list all clients.
func ClientList(c *cli.Context, client solder.Client) error {
	return nil
}

// ClientShow provides the sub-command to show client details.
func ClientShow(c *cli.Context, client solder.Client) error {
	return nil
}

// ClientUpdate provides the sub-command to update a client.
func ClientUpdate(c *cli.Context, client solder.Client) error {
	return nil
}

// ClientDelete provides the sub-command to delete a client.
func ClientDelete(c *cli.Context, client solder.Client) error {
	return nil
}

// ClientCreate provides the sub-command to create a client.
func ClientCreate(c *cli.Context, client solder.Client) error {
	return nil
}
