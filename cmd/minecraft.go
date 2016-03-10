package cmd

import (
	"github.com/codegangsta/cli"
)

// Minecraft provides the sub-command for the minecraft API.
func Minecraft() cli.Command {
	return cli.Command{
		Name:  "minecraft",
		Usage: "Minecraft related sub-commands",
		Subcommands: []cli.Command{
			cli.Command{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all Minecraft versions",
				Action: func(c *cli.Context) {
					Handle(c, MinecraftList)
				},
			},
			cli.Command{
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
func MinecraftList() cli.Command {
	return nil
}

// MinecraftRefresh provides the sub-command to refresh the Minecraft versions.
func MinecraftRefresh() cli.Command {
	return nil
}
