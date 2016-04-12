package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
	"github.com/solderapp/solder-go/solder"
)

// Minecraft provides the sub-command for the Minecraft API.
func Minecraft() cli.Command {
	return cli.Command{
		Name:  "minecraft",
		Usage: "Minecraft related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all Minecraft versions",
				ArgsUsage: " ",
				Action: func(c *cli.Context) {
					Handle(c, MinecraftList)
				},
			},
			{
				Name:      "refresh",
				Aliases:   []string{"ref"},
				Usage:     "Refresh Minecraft versions",
				ArgsUsage: " ",
				Action: func(c *cli.Context) {
					Handle(c, MinecraftRefresh)
				},
			},
			{
				Name:      "build-list",
				Usage:     "List assigned builds",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Minecraft ID or slug to list builds",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, MinecraftBuildList)
				},
			},
			{
				Name:      "build-append",
				Usage:     "Append a build to Minecraft",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Minecraft ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "build, b",
						Value: "",
						Usage: "Build ID or slug to append",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, MinecraftBuildAppend)
				},
			},
			{
				Name:      "build-remove",
				Usage:     "Remove a build from Minecraft",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Minecraft ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "build, b",
						Value: "",
						Usage: "Build ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, MinecraftBuildRemove)
				},
			},
		},
	}
}

// MinecraftList provides the sub-command to list all Minecraft versions.
func MinecraftList(c *cli.Context, client solder.ClientAPI) error {
	records, err := client.MinecraftList()

	if err != nil {
		return err
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"ID", "Slug", "Version", "Type"})

	for _, record := range records {
		table.Append(
			[]string{
				strconv.FormatInt(record.ID, 10),
				record.Slug,
				record.Version,
				record.Type,
			},
		)
	}

	table.Render()
	return nil
}

// MinecraftRefresh provides the sub-command to refresh the Minecraft versions.
func MinecraftRefresh(c *cli.Context, client solder.ClientAPI) error {
	err := client.MinecraftRefresh()

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully refreshed\n")
	return nil
}

// MinecraftBuildList provides the sub-command to list builds of the Minecraft.
func MinecraftBuildList(c *cli.Context, client solder.ClientAPI) error {
	records, err := client.MinecraftBuildList(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"ID", "Slug", "Name"})

	for _, record := range records {
		table.Append(
			[]string{
				strconv.FormatInt(record.ID, 10),
				record.Slug,
				record.Name,
			},
		)
	}

	table.Render()
	return nil
}

// MinecraftBuildAppend provides the sub-command to append a build to the Minecraft.
func MinecraftBuildAppend(c *cli.Context, client solder.ClientAPI) error {
	err := client.MinecraftBuildAppend(
		GetIdentifierParam(c),
		GetBuildParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to Minecraft\n")
	return nil
}

// MinecraftBuildRemove provides the sub-command to remove a build from the Minecraft.
func MinecraftBuildRemove(c *cli.Context, client solder.ClientAPI) error {
	err := client.MinecraftBuildDelete(
		GetIdentifierParam(c),
		GetBuildParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from Minecraft\n")
	return nil
}
