package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
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
						Usage: "Version ID or slug to update",
					},
					cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Define an optional slug",
					},
					cli.StringFlag{
						Name:  "name",
						Value: "",
						Usage: "Define a required name",
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
						Usage: "Version ID or slug to delete",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, PackDelete)
				},
			},
			{
				Name:  "create",
				Usage: "Create a pack",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Define an optional slug",
					},
					cli.StringFlag{
						Name:  "name",
						Value: "",
						Usage: "Define a required name",
					},
				},
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

	if err != nil {
		return err
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Error: Empty result\n")
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

// PackShow provides the sub-command to show pack details.
func PackShow(c *cli.Context, client solder.API) error {
	record, err := client.PackGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Key", "Value"})

	table.AppendBulk(
		[][]string{
			[]string{"ID", strconv.FormatInt(record.ID, 10)},
			[]string{"Slug", record.Slug},
			[]string{"Name", record.Name},
			[]string{"Created", record.CreatedAt.Format(time.UnixDate)},
			[]string{"Updated", record.UpdatedAt.Format(time.UnixDate)},
		},
	)

	table.Render()
	return nil
}

// PackDelete provides the sub-command to delete a pack.
func PackDelete(c *cli.Context, client solder.API) error {
	err := client.PackDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// PackUpdate provides the sub-command to update a pack.
func PackUpdate(c *cli.Context, client solder.API) error {
	record, err := client.PackGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	if val := c.String("slug"); val != "" {
		record.Slug = val
	}

	if val := c.String("name"); val != "" {
		record.Name = val
	}

	_, patch := client.PackPatch(record)

	if patch != nil {
		return patch
	}

	fmt.Fprintf(os.Stderr, "Successfully updated\n")
	return nil
}

// PackCreate provides the sub-command to create a pack.
func PackCreate(c *cli.Context, client solder.API) error {
	record := &solder.Pack{}

	if val := c.String("slug"); val != "" {
		record.Slug = val
	}

	if val := c.String("name"); val != "" {
		record.Name = val
	} else {
		fmt.Println("Error: You must provide a name.")
		os.Exit(1)
	}

	_, err := client.PackPost(record)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}
