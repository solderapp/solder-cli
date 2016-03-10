package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/solderapp/solder-cli/solder"
)

// Minecraft provides the sub-command for the minecraft API.
func Minecraft() cli.Command {
	return cli.Command{
		Name:  "minecraft",
		Usage: "Minecraft related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all Minecraft versions",
				Action: func(c *cli.Context) {
					Handle(c, MinecraftList)
				},
			},
			{
				Name:    "refresh",
				Aliases: []string{"ref"},
				Usage:   "Refresh the Minecraft versions",
				Action: func(c *cli.Context) {
					Handle(c, MinecraftRefresh)
				},
			},
		},
	}
}

// MinecraftList provides the sub-command to list all Minecraft versions.
func MinecraftList(c *cli.Context, client solder.API) error {
	return nil
}

// MinecraftRefresh provides the sub-command to refresh the Minecraft versions.
func MinecraftRefresh(c *cli.Context, client solder.API) error {
	return nil
}
