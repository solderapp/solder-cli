package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/codegangsta/cli"
	"github.com/solderapp/solder-cli/solder"
)

// Mod provides the sub-command for the mod API.
func Mod() cli.Command {
	return cli.Command{
		Name:    "mod",
		Aliases: []string{"m"},
		Usage:   "Mod related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all mods",
				Action: func(c *cli.Context) {
					Handle(c, ModList)
				},
			},
			{
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
			{
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
			{
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
			{
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
func ModList(c *cli.Context, client solder.API) error {
	records, err := client.ModList()

	if err != nil || len(records) == 0 {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)

	for _, record := range records {
		fmt.Fprintf(w, "%s\t%s\t%s\n", record.ID, record.CreatedAt, record.UpdatedAt)
	}

	w.Flush()
	return nil
}

// ModShow provides the sub-command to show mod details.
func ModShow(c *cli.Context, client solder.API) error {
	return nil
}

// ModUpdate provides the sub-command to update a mod.
func ModUpdate(c *cli.Context, client solder.API) error {
	return nil
}

// ModDelete provides the sub-command to delete a mod.
func ModDelete(c *cli.Context, client solder.API) error {
	return nil
}

// ModCreate provides the sub-command to create a mod.
func ModCreate(c *cli.Context, client solder.API) error {
	return nil
}
