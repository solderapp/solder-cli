package cmd

import (
	"github.com/codegangsta/cli"
)

// Version provides the sub-command for the version API.
func Version() cli.Command {
	return cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Version related sub-commands",
		Subcommands: []cli.Command{
			cli.Command{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all versions",
				Action: func(c *cli.Context) {
					Handle(c, VersionList)
				},
			},
			cli.Command{
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
			cli.Command{
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
			cli.Command{
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
			cli.Command{
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
func VersionList(c *cli.Context, client solder.Client) error {
	return nil
}

// VersionShow provides the sub-command to show version details.
func VersionShow(c *cli.Context, client solder.Client) error {
	return nil
}

// VersionUpdate provides the sub-command to update a version.
func VersionUpdate(c *cli.Context, client solder.Client) error {
	return nil
}

// VersionDelete provides the sub-command to delete a version.
func VersionDelete(c *cli.Context, client solder.Client) error {
	return nil
}

// VersionCreate provides the sub-command to create a version.
func VersionCreate(c *cli.Context, client solder.Client) error {
	return nil
}
