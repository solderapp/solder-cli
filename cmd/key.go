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

// Key provides the sub-command for the key API.
func Key() cli.Command {
	return cli.Command{
		Name:    "key",
		Aliases: []string{"k"},
		Usage:   "Key related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all keys",
				Action: func(c *cli.Context) {
					Handle(c, KeyList)
				},
			},
			{
				Name:  "show",
				Usage: "Display a key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Key ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, KeyShow)
				},
			},
			{
				Name:  "update",
				Usage: "Update a key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Key ID or slug to show",
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
						Name:  "key",
						Value: "",
						Usage: "Provide a key",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, KeyUpdate)
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"rm"},
				Usage:   "Delete a key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Key ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, KeyDelete)
				},
			},
			{
				Name:  "create",
				Usage: "Create a key",
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
						Name:  "key",
						Value: "",
						Usage: "Provide a key",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, KeyCreate)
				},
			},
		},
	}
}

// KeyList provides the sub-command to list all keys.
func KeyList(c *cli.Context, client solder.API) error {
	records, err := client.KeyList()

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

// KeyShow provides the sub-command to show key details.
func KeyShow(c *cli.Context, client solder.API) error {
	record, err := client.KeyGet(
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
			{"Key", record.Value},
			{"Created", record.CreatedAt.Format(time.UnixDate)},
			{"Updated", record.UpdatedAt.Format(time.UnixDate)},
		},
	)

	table.Render()
	return nil
}

// KeyDelete provides the sub-command to delete a key.
func KeyDelete(c *cli.Context, client solder.API) error {
	err := client.KeyDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// KeyUpdate provides the sub-command to update a key.
func KeyUpdate(c *cli.Context, client solder.API) error {
	record, err := client.KeyGet(
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

	if val := c.String("key"); val != record.Value {
		record.Value = val
	}

	_, patch := client.KeyPatch(record)

	if patch != nil {
		return patch
	}

	fmt.Fprintf(os.Stderr, "Successfully updated\n")
	return nil
}

// KeyCreate provides the sub-command to create a key.
func KeyCreate(c *cli.Context, client solder.API) error {
	record := &solder.Key{}

	if val := c.String("name"); val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("You must provide a name.")
	}

	if val := c.String("slug"); val != "" {
		record.Slug = val
	}

	if val := c.String("key"); val != "" {
		record.Value = val
	}

	_, err := client.KeyPost(record)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}
