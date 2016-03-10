package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
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
	records, err := client.MinecraftList()

	if err != nil || len(records) == 0 {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"ID", "Version", "Type"})

	for _, record := range records {
		table.Append(
			[]string{
				strconv.FormatInt(record.ID, 10),
				record.Version,
				record.Type,
			},
		)
	}

	table.Render()
	return nil
}

// MinecraftRefresh provides the sub-command to refresh the Minecraft versions.
func MinecraftRefresh(c *cli.Context, client solder.API) error {
	err := client.MinecraftRefresh()

	if err != nil {
		return err
	}

	fmt.Println("Successfully refreshed Minecraft versions")
	return nil
}
