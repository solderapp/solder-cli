package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kleister/kleister-go/kleister"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// Client provides the sub-command for the client API.
func Client() cli.Command {
	return cli.Command{
		Name:    "client",
		Aliases: []string{"c"},
		Usage:   "Client related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all clients",
				ArgsUsage: " ",
				Action: func(c *cli.Context) {
					Handle(c, ClientList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a client",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Client ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ClientShow)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a client",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
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
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a client",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Client ID or slug to delete",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ClientDelete)
				},
			},
			{
				Name:      "create",
				Usage:     "Create a client",
				ArgsUsage: " ",
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
			{
				Name:      "pack-list",
				Usage:     "List assigned packs",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Client ID or slug to list packs",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ClientPackList)
				},
			},
			{
				Name:      "pack-append",
				Usage:     "Append a pack to client",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Client ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "Pack ID or slug to append",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ClientPackAppend)
				},
			},
			{
				Name:      "pack-remove",
				Usage:     "Remove a pack from client",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Client ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "Pack ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ClientPackRemove)
				},
			},
		},
	}
}

// ClientList provides the sub-command to list all clients.
func ClientList(c *cli.Context, client kleister.ClientAPI) error {
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
func ClientShow(c *cli.Context, client kleister.ClientAPI) error {
	record, err := client.ClientGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Key", "Value"})

	table.Append(
		[]string{
			"ID",
			strconv.FormatInt(record.ID, 10),
		},
	)

	table.Append(
		[]string{
			"Slug",
			record.Slug,
		},
	)

	table.Append(
		[]string{
			"Name",
			record.Name,
		},
	)

	table.Append(
		[]string{
			"UUID",
			record.Value,
		},
	)

	table.Append(
		[]string{
			"Created",
			record.CreatedAt.Format(time.UnixDate),
		},
	)

	table.Append(
		[]string{
			"Updated",
			record.UpdatedAt.Format(time.UnixDate),
		},
	)

	table.Render()
	return nil
}

// ClientDelete provides the sub-command to delete a client.
func ClientDelete(c *cli.Context, client kleister.ClientAPI) error {
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
func ClientUpdate(c *cli.Context, client kleister.ClientAPI) error {
	record, err := client.ClientGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	changed := false

	if val := c.String("name"); c.IsSet("name") && val != record.Name {
		record.Name = val
		changed = true
	}

	if val := c.String("uuid"); c.IsSet("uuid") && val != record.Value {
		record.Value = val
		changed = true
	}

	if val := c.String("slug"); c.IsSet("slug") && val != record.Slug {
		record.Slug = val
		changed = true
	}

	if changed {
		_, patch := client.ClientPatch(
			record,
		)

		if patch != nil {
			return patch
		}

		fmt.Fprintf(os.Stderr, "Successfully updated\n")
	} else {
		fmt.Fprintf(os.Stderr, "Nothing to update...\n")
	}

	return nil
}

// ClientCreate provides the sub-command to create a client.
func ClientCreate(c *cli.Context, client kleister.ClientAPI) error {
	record := &kleister.Client{}

	if val := c.String("name"); c.IsSet("name") && val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("You must provide a name.")
	}

	if val := c.String("uuid"); c.IsSet("uuid") && val != "" {
		record.Value = val
	} else {
		return fmt.Errorf("You must provide a UUID.")
	}

	if val := c.String("slug"); c.IsSet("slug") && val != "" {
		record.Slug = val
	}

	_, err := client.ClientPost(
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// ClientPackList provides the sub-command to list packs of the client.
func ClientPackList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.ClientPackList(
		kleister.ClientPackParams{
			Client: GetIdentifierParam(c),
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
	table.SetHeader([]string{"Pack"})

	for _, record := range records {
		table.Append(
			[]string{
				record.Slug,
			},
		)
	}

	table.Render()
	return nil
}

// ClientPackAppend provides the sub-command to append a pack to the client.
func ClientPackAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.ClientPackAppend(
		kleister.ClientPackParams{
			Client: GetIdentifierParam(c),
			Pack:   GetPackParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to client\n")
	return nil
}

// ClientPackRemove provides the sub-command to remove a pack from the client.
func ClientPackRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.ClientPackDelete(
		kleister.ClientPackParams{
			Client: GetIdentifierParam(c),
			Pack:   GetPackParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from client\n")
	return nil
}
