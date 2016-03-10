package cmd

import (
	"github.com/codegangsta/cli"
)

// Pack provides the sub-command for the pack API.
func Pack() cli.Command {
	return cli.Command{
		Name:    "pack",
		Aliases: []string{"p"},
		Usage:   "Pack related sub-commands",
		Subcommands: []cli.Command{
			cli.Command{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all packs",
				Action: func(c *cli.Context) {
					Handle(c, PackList)
				},
			},
			cli.Command{
				Name:  "show",
				Usage: "Display a pack",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, PackShow)
				},
			},
			cli.Command{
				Name:  "update",
				Usage: "Update a pack",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, PackUpdate)
				},
			},
			cli.Command{
				Name:    "delete",
				Aliases: []string{"rm"},
				Usage:   "Delete a pack",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, PackDelete)
				},
			},
			cli.Command{
				Name:  "create",
				Usage: "Create a pack",
				Action: func(c *cli.Context) {
					Handle(c, PackCreate)
				},
			},
		},
	}
}

// PackList provides the sub-command to list all packs.
func PackList(c *cli.Context, client solder.Client) error {
	return nil
}

// PackShow provides the sub-command to show pack details.
func PackShow(c *cli.Context, client solder.Client) error {
	return nil
}

// PackUpdate provides the sub-command to update a pack.
func PackUpdate(c *cli.Context, client solder.Client) error {
	return nil
}

// PackDelete provides the sub-command to delete a pack.
func PackDelete(c *cli.Context, client solder.Client) error {
	return nil
}

// PackCreate provides the sub-command to create a pack.
func PackCreate(c *cli.Context, client solder.Client) error {
	return nil
}
