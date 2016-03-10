package cmd

import (
	"github.com/codegangsta/cli"
)

// Key provides the sub-command for the key API.
func Key() cli.Command {
	return cli.Command{
		Name:    "key",
		Aliases: []string{"k"},
		Usage:   "Key related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all keys",
				Action: func(c *cli.Context) {
					Handle(c, KeyList)
				},
			},
			{
				Name:  "show",
				Usage: "Display a key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, KeyShow)
				},
			},
			{
				Name:  "update",
				Usage: "Update a key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, KeyUpdate)
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"rm"},
				Usage:   "Delete a key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, KeyDelete)
				},
			},
			{
				Name:  "create",
				Usage: "Create a key",
				Action: func(c *cli.Context) {
					Handle(c, KeyCreate)
				},
			},
		},
	}
}

// KeyList provides the sub-command to list all keys.
func KeyList(c *cli.Context, client solder.Client) error {
	return nil
}

// KeyShow provides the sub-command to show key details.
func KeyShow(c *cli.Context, client solder.Client) error {
	return nil
}

// KeyUpdate provides the sub-command to update a key.
func KeyUpdate(c *cli.Context, client solder.Client) error {
	return nil
}

// KeyDelete provides the sub-command to delete a key.
func KeyDelete(c *cli.Context, client solder.Client) error {
	return nil
}

// KeyCreate provides the sub-command to create a key.
func KeyCreate(c *cli.Context, client solder.Client) error {
	return nil
}
