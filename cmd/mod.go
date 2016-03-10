package cmd

import (
	"github.com/codegangsta/cli"
)

// Mod provides the sub-command for the mod API.
func Mod() cli.Command {
	return cli.Command{
		Name:    "mod",
		Aliases: []string{"m"},
		Usage:   "Mod related sub-commands",
		Subcommands: []cli.Command{
			cli.Command{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all mods",
				Action: func(c *cli.Context) {
					Handle(c, ModList)
				},
			},
			cli.Command{
				Name:  "show",
				Usage: "Display a mod",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ModShow)
				},
			},
			cli.Command{
				Name:  "update",
				Usage: "Update a mod",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ModUpdate)
				},
			},
			cli.Command{
				Name:    "delete",
				Aliases: []string{"rm"},
				Usage:   "Delete a mod",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ModDelete)
				},
			},
			cli.Command{
				Name:  "create",
				Usage: "Create a mod",
				Action: func(c *cli.Context) {
					Handle(c, ModCreate)
				},
			},
		},
	}
}

// ModList provides the sub-command to list all mods.
func ModList(c *cli.Context, client solder.Client) error {
	return nil
}

// ModShow provides the sub-command to show mod details.
func ModShow(c *cli.Context, client solder.Client) error {
	return nil
}

// ModUpdate provides the sub-command to update a mod.
func ModUpdate(c *cli.Context, client solder.Client) error {
	return nil
}

// ModDelete provides the sub-command to delete a mod.
func ModDelete(c *cli.Context, client solder.Client) error {
	return nil
}

// ModCreate provides the sub-command to create a mod.
func ModCreate(c *cli.Context, client solder.Client) error {
	return nil
}
