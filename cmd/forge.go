package cmd

import (
	"github.com/codegangsta/cli"
)

// Forge provides the sub-command for the forge API.
func Forge() cli.Command {
	return cli.Command{
		Name:  "forge",
		Usage: "Forge related sub-commands",
		Subcommands: []cli.Command{
			cli.Command{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all Forge versions",
				Action: func(c *cli.Context) {
					Handle(c, ForgeList)
				},
			},
			cli.Command{
				Name:    "refresh",
				Aliases: []string{"ref"},
				Usage:   "Refresh the Forge versions",
				Action: func(c *cli.Context) {
					Handle(c, ForgeRefresh)
				},
			},
		},
	}
}

// ForgeList provides the sub-command to list all Forge versions.
func ForgeList() cli.Command {
	return nil
}

// ForgeRefresh provides the sub-command to refresh the Forge versions.
func ForgeRefresh() cli.Command {
	return nil
}
