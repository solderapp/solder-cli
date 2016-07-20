package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kleister/kleister-go/kleister"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// Forge provides the sub-command for the Forge API.
func Forge() cli.Command {
	return cli.Command{
		Name:  "forge",
		Usage: "Forge related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all Forge versions",
				ArgsUsage: " ",
				Action: func(c *cli.Context) {
					Handle(c, ForgeList)
				},
			},
			{
				Name:      "refresh",
				Aliases:   []string{"ref"},
				Usage:     "Refresh Forge versions",
				ArgsUsage: " ",
				Action: func(c *cli.Context) {
					Handle(c, ForgeRefresh)
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
						Usage: "Forge ID or slug to list builds",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ForgeBuildList)
				},
			},
			{
				Name:      "build-append",
				Usage:     "Append a build to Forge",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Forge ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "Pack ID or slug to append",
					},
					cli.StringFlag{
						Name:  "build, b",
						Value: "",
						Usage: "Build ID or slug to append",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ForgeBuildAppend)
				},
			},
			{
				Name:      "build-remove",
				Usage:     "Remove a build from Forge",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Forge ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "Pack ID or slug to remove",
					},
					cli.StringFlag{
						Name:  "build, b",
						Value: "",
						Usage: "Build ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ForgeBuildRemove)
				},
			},
		},
	}
}

// ForgeList provides the sub-command to list all Forge versions.
func ForgeList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.ForgeList()

	if err != nil {
		return err
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"ID", "Slug", "Version", "Minecraft"})

	for _, record := range records {
		table.Append(
			[]string{
				strconv.FormatInt(record.ID, 10),
				record.Slug,
				record.Version,
				record.Minecraft,
			},
		)
	}

	table.Render()
	return nil
}

// ForgeRefresh provides the sub-command to refresh the Forge versions.
func ForgeRefresh(c *cli.Context, client kleister.ClientAPI) error {
	err := client.ForgeRefresh()

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully refreshed\n")
	return nil
}

// ForgeBuildList provides the sub-command to list builds of the Forge.
func ForgeBuildList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.ForgeBuildList(
		kleister.ForgeBuildParams{
			Forge: GetIdentifierParam(c),
		},
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
	table.SetHeader([]string{"Pack", "Build"})

	for _, record := range records {
		pack := "n/a"

		if record.Pack != nil {
			pack = record.Pack.Slug
		}

		table.Append(
			[]string{
				pack,
				record.Slug,
			},
		)
	}

	table.Render()
	return nil
}

// ForgeBuildAppend provides the sub-command to append a build to the Forge.
func ForgeBuildAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.ForgeBuildAppend(
		kleister.ForgeBuildParams{
			Forge: GetIdentifierParam(c),
			Pack:  GetPackParam(c),
			Build: GetBuildParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to Forge\n")
	return nil
}

// ForgeBuildRemove provides the sub-command to remove a build from the Forge.
func ForgeBuildRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.ForgeBuildDelete(
		kleister.ForgeBuildParams{
			Forge: GetIdentifierParam(c),
			Pack:  GetPackParam(c),
			Build: GetBuildParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from Forge\n")
	return nil
}
