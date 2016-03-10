package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/codegangsta/cli"
	"github.com/solderapp/solder-cli/solder"
)

// Version provides the sub-command for the version API.
func Version() cli.Command {
	return cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Version related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all versions",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "mod, m",
						Value: "",
						Usage: "ID or slug of the related mod",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionList)
				},
			},
			{
				Name:  "show",
				Usage: "Display a version",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionShow)
				},
			},
			{
				Name:  "update",
				Usage: "Update a version",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionUpdate)
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"rm"},
				Usage:   "Delete a version",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionDelete)
				},
			},
			{
				Name:  "create",
				Usage: "Create a version",
				Action: func(c *cli.Context) {
					Handle(c, VersionCreate)
				},
			},
		},
	}
}

// VersionList provides the sub-command to list all versions.
func VersionList(c *cli.Context, client solder.API) error {
	mod := c.GlobalString("mod")

	if mod == "" {
		fmt.Println("Error: you must provide a mod ID or slug.")
		os.Exit(1)
	}

	records, err := client.VersionList(mod)

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

// VersionShow provides the sub-command to show version details.
func VersionShow(c *cli.Context, client solder.API) error {
	return nil
}

// VersionUpdate provides the sub-command to update a version.
func VersionUpdate(c *cli.Context, client solder.API) error {
	return nil
}

// VersionDelete provides the sub-command to delete a version.
func VersionDelete(c *cli.Context, client solder.API) error {
	return nil
}

// VersionCreate provides the sub-command to create a version.
func VersionCreate(c *cli.Context, client solder.API) error {
	return nil
}
