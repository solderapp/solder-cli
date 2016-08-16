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

// Mod provides the sub-command for the mod API.
func Mod() cli.Command {
	return cli.Command{
		Name:    "mod",
		Aliases: []string{"m"},
		Usage:   "Mod related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all mods",
				ArgsUsage: " ",
				Action: func(c *cli.Context) error {
					return Handle(c, ModList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a mod",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Mod ID or slug to show",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, ModShow)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a mod",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
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
				Action: func(c *cli.Context) error {
					return Handle(c, ModUpdate)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a mod",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Mod ID or slug to delete",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, ModDelete)
				},
			},
			{
				Name:      "create",
				Usage:     "Create a mod",
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
				Action: func(c *cli.Context) error {
					return Handle(c, ModCreate)
				},
			},
			{
				Name:      "user-list",
				Usage:     "List assigned users",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Mod ID or slug to list users",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, ModUserList)
				},
			},
			{
				Name:      "user-append",
				Usage:     "Append a user to mod",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Mod ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "user, u",
						Value: "",
						Usage: "User ID or slug to append",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, ModUserAppend)
				},
			},
			{
				Name:      "user-remove",
				Usage:     "Remove a user from mod",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Mod ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "user, u",
						Value: "",
						Usage: "User ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, ModUserRemove)
				},
			},
		},
	}
}

// ModList provides the sub-command to list all mods.
func ModList(c *cli.Context, client kleister.ClientAPI) error {
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
func ModShow(c *cli.Context, client kleister.ClientAPI) error {
	record, err := client.ModGet(
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
			"Description",
			record.Description,
		},
	)

	table.Append(
		[]string{
			"Author",
			record.Author,
		},
	)

	table.Append(
		[]string{
			"Website",
			record.Website,
		},
	)

	table.Append(
		[]string{
			"Donate",
			record.Donate,
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

// ModDelete provides the sub-command to delete a mod.
func ModDelete(c *cli.Context, client kleister.ClientAPI) error {
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
func ModUpdate(c *cli.Context, client kleister.ClientAPI) error {
	record, err := client.ModGet(
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

	if val := c.String("slug"); c.IsSet("slug") && val != record.Slug {
		record.Slug = val
		changed = true
	}

	if val := c.String("description"); c.IsSet("description") && val != record.Description {
		record.Description = val
		changed = true
	}

	if val := c.String("author"); c.IsSet("author") && val != record.Author {
		record.Author = val
		changed = true
	}

	if val := c.String("website"); c.IsSet("website") && val != record.Website {
		record.Website = val
		changed = true
	}

	if val := c.String("donate"); c.IsSet("donate") && val != record.Donate {
		record.Donate = val
		changed = true
	}

	if changed {
		_, patch := client.ModPatch(
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

// ModCreate provides the sub-command to create a mod.
func ModCreate(c *cli.Context, client kleister.ClientAPI) error {
	record := &kleister.Mod{}

	if val := c.String("name"); c.IsSet("name") && val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("You must provide a name.")
	}

	if val := c.String("slug"); c.IsSet("slug") && val != "" {
		record.Slug = val
	}

	if val := c.String("description"); c.IsSet("description") && val != "" {
		record.Description = val
	}

	if val := c.String("author"); c.IsSet("author") && val != "" {
		record.Author = val
	}

	if val := c.String("website"); c.IsSet("website") && val != "" {
		record.Website = val
	}

	if val := c.String("donate"); c.IsSet("donate") && val != "" {
		record.Donate = val
	}

	_, err := client.ModPost(
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// ModUserList provides the sub-command to list users of the mod.
func ModUserList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.ModUserList(
		kleister.ModUserParams{
			Mod: GetIdentifierParam(c),
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
	table.SetHeader([]string{"User"})

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

// ModUserAppend provides the sub-command to append a user to the mod.
func ModUserAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.ModUserAppend(
		kleister.ModUserParams{
			Mod:  GetIdentifierParam(c),
			User: GetUserParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to mod\n")
	return nil
}

// ModUserRemove provides the sub-command to remove a user from the mod.
func ModUserRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.ModUserDelete(
		kleister.ModUserParams{
			Mod:  GetIdentifierParam(c),
			User: GetUserParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from mod\n")
	return nil
}
