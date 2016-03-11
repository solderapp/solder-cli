package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
	"github.com/solderapp/solder-cli/solder"
)

// Client provides the sub-command for the client API.
func Client() cli.Command {
	return cli.Command{
		Name:    "client",
		Aliases: []string{"c"},
		Usage:   "Client related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all clients",
				Action: func(c *cli.Context) {
					Handle(c, ClientList)
				},
			},
			{
				Name:  "show",
				Usage: "Display a build",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Client ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ClientShow)
				},
			},
			{
				Name:  "update",
				Usage: "Update a client",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Client ID or slug to update",
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
						Name:  "uuid",
						Value: "",
						Usage: "Provide a UUID",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ClientUpdate)
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"rm"},
				Usage:   "Delete a client",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Client ID or slug to delete",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ClientDelete)
				},
			},
			{
				Name:  "create",
				Usage: "Create a client",
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
						Name:  "uuid",
						Value: "",
						Usage: "Provide a UUID",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ClientCreate)
				},
			},
		},
	}
}

// ClientList provides the sub-command to list all clients.
func ClientList(c *cli.Context, client solder.API) error {
	records, err := client.ClientList()

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

// ClientShow provides the sub-command to show client details.
func ClientShow(c *cli.Context, client solder.API) error {
	record, err := client.ClientGet(
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
			[]string{"UUID", record.Value},
			[]string{"Created", record.CreatedAt.Format(time.UnixDate)},
			[]string{"Updated", record.UpdatedAt.Format(time.UnixDate)},
		},
	)

	table.Render()
	return nil
}

// ClientDelete provides the sub-command to delete a client.
func ClientDelete(c *cli.Context, client solder.API) error {
	err := client.ClientDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// ClientUpdate provides the sub-command to update a client.
func ClientUpdate(c *cli.Context, client solder.API) error {
	record, err := client.ClientGet(
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

	if val := c.String("uuid"); val != record.Value {
		record.Value = val
	}

	_, patch := client.ClientPatch(record)

	if patch != nil {
		return patch
	}

	fmt.Fprintf(os.Stderr, "Successfully updated\n")
	return nil
}

// ClientCreate provides the sub-command to create a client.
func ClientCreate(c *cli.Context, client solder.API) error {
	record := &solder.Client{}

	if val := c.String("name"); val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("You must provide a name.")
	}

	if val := c.String("slug"); val != "" {
		record.Slug = val
	}

	if val := c.String("uuid"); val != "" {
		record.Value = val
	}

	_, err := client.ClientPost(record)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}
