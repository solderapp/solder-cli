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

// Mod provides the sub-command for the mod API.
func Mod() cli.Command {
	return cli.Command{
		Name:    "mod",
		Aliases: []string{"m"},
		Usage:   "Mod related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all mods",
				Action: func(c *cli.Context) {
					Handle(c, ModList)
				},
			},
			{
				Name:  "show",
				Usage: "Display a mod",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Mod ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ModShow)
				},
			},
			{
				Name:  "update",
				Usage: "Update a mod",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Mod ID or slug to update",
					},
					cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Provide a slug",
					},
					cli.StringFlag{
						Name:  "name",
						Value: "",
						Usage: "Provide a name",
					},
					cli.StringFlag{
						Name:  "description",
						Value: "",
						Usage: "Provide a description",
					},
					cli.StringFlag{
						Name:  "author",
						Value: "",
						Usage: "Provide an author",
					},
					cli.StringFlag{
						Name:  "website-link",
						Value: "",
						Usage: "Provide a website link",
					},
					cli.StringFlag{
						Name:  "donate-link",
						Value: "",
						Usage: "Provide a donate link",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ModUpdate)
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"rm"},
				Usage:   "Delete a mod",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Mod ID or slug to delete",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ModDelete)
				},
			},
			{
				Name:  "create",
				Usage: "Create a mod",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Provide a slug",
					},
					cli.StringFlag{
						Name:  "name",
						Value: "",
						Usage: "Provide a name",
					},
					cli.StringFlag{
						Name:  "description",
						Value: "",
						Usage: "Provide a description",
					},
					cli.StringFlag{
						Name:  "author",
						Value: "",
						Usage: "Provide an author",
					},
					cli.StringFlag{
						Name:  "website-link",
						Value: "",
						Usage: "Provide a website link",
					},
					cli.StringFlag{
						Name:  "donate-link",
						Value: "",
						Usage: "Provide a donate link",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ModCreate)
				},
			},
		},
	}
}

// ModList provides the sub-command to list all mods.
func ModList(c *cli.Context, client solder.API) error {
	records, err := client.ModList()

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

// ModShow provides the sub-command to show mod details.
func ModShow(c *cli.Context, client solder.API) error {
	record, err := client.ModGet(
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
			{"ID", strconv.FormatInt(record.ID, 10)},
			{"Slug", record.Slug},
			{"Name", record.Name},
			{"Description", record.Description},
			{"Author", record.Author},
			{"Website", record.Website},
			{"Donate", record.Donate},
			{"Created", record.CreatedAt.Format(time.UnixDate)},
			{"Updated", record.UpdatedAt.Format(time.UnixDate)},
		},
	)

	table.Render()
	return nil
}

// ModDelete provides the sub-command to delete a mod.
func ModDelete(c *cli.Context, client solder.API) error {
	err := client.ModDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// ModUpdate provides the sub-command to update a mod.
func ModUpdate(c *cli.Context, client solder.API) error {
	record, err := client.ModGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	if val := c.String("name"); val != record.Name {
		record.Name = val
	}

	if val := c.String("slug"); val != record.Slug {
		record.Slug = val
	}

	if val := c.String("description"); val != record.Description {
		record.Description = val
	}

	if val := c.String("author"); val != record.Author {
		record.Author = val
	}

	if val := c.String("website"); val != record.Website {
		record.Website = val
	}

	if val := c.String("donate"); val != record.Donate {
		record.Donate = val
	}

	_, patch := client.ModPatch(record)

	if patch != nil {
		return patch
	}

	fmt.Fprintf(os.Stderr, "Successfully updated\n")
	return nil
}

// ModCreate provides the sub-command to create a mod.
func ModCreate(c *cli.Context, client solder.API) error {
	record := &solder.Mod{}

	if val := c.String("name"); val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("You must provide a name.")
	}

	if val := c.String("slug"); val != "" {
		record.Slug = val
	}

	if val := c.String("description"); val != "" {
		record.Description = val
	}

	if val := c.String("author"); val != "" {
		record.Author = val
	}

	if val := c.String("website"); val != "" {
		record.Website = val
	}

	if val := c.String("donate"); val != "" {
		record.Donate = val
	}

	_, err := client.ModPost(record)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}
