package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/codegangsta/cli"
	"github.com/solderapp/solder-cli/solder"
)

// Pack provides the sub-command for the pack API.
func Pack() cli.Command {
	return cli.Command{
		Name:    "pack",
		Aliases: []string{"p"},
		Usage:   "Pack related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all packs",
				Action: func(c *cli.Context) {
					Handle(c, PackList)
				},
			},
			{
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
			{
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
			{
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
			{
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
func PackList(c *cli.Context, client solder.API) error {
	records, err := client.PackList()

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

// PackShow provides the sub-command to show pack details.
func PackShow(c *cli.Context, client solder.API) error {
	return nil
}

// PackUpdate provides the sub-command to update a pack.
func PackUpdate(c *cli.Context, client solder.API) error {
	return nil
}

// PackDelete provides the sub-command to delete a pack.
func PackDelete(c *cli.Context, client solder.API) error {
	return nil
}

// PackCreate provides the sub-command to create a pack.
func PackCreate(c *cli.Context, client solder.API) error {
	return nil
}
