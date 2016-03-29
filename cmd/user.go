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

// User provides the sub-command for the user API.
func User() cli.Command {
	return cli.Command{
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "User related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all users",
				ArgsUsage: " ",
				Action: func(c *cli.Context) {
					Handle(c, UserList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, UserShow)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a user",
				ArgsUsage: " ",
				Flags: append(
					[]cli.Flag{
						cli.StringFlag{
							Name:  "id, i",
							Value: "",
							Usage: "User ID or slug to update",
						},
						cli.StringFlag{
							Name:  "slug",
							Value: "",
							Usage: "Provide a slug",
						},
						cli.StringFlag{
							Name:  "username",
							Value: "",
							Usage: "Provide an username",
						},
						cli.StringFlag{
							Name:  "email",
							Value: "",
							Usage: "Provide an email",
						},
						cli.StringFlag{
							Name:  "password",
							Value: "",
							Usage: "Provide a password",
						},
					},
					userPermissionFlags()...,
				),
				Action: func(c *cli.Context) {
					Handle(c, UserUpdate)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, UserDelete)
				},
			},
			{
				Name:      "create",
				Usage:     "Create a user",
				ArgsUsage: " ",
				Flags: append(
					[]cli.Flag{
						cli.StringFlag{
							Name:  "slug",
							Value: "",
							Usage: "Provide a slug",
						},
						cli.StringFlag{
							Name:  "username",
							Value: "",
							Usage: "Provide an username",
						},
						cli.StringFlag{
							Name:  "email",
							Value: "",
							Usage: "Provide an email",
						},
						cli.StringFlag{
							Name:  "password",
							Value: "",
							Usage: "Provide a password",
						},
					},
					userPermissionFlags()...,
				),
				Action: func(c *cli.Context) {
					Handle(c, UserCreate)
				},
			},
			{
				Name:      "mod-list",
				Usage:     "List assigned mods",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to list mods",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, UserModList)
				},
			},
			{
				Name:      "mod-append",
				Usage:     "Append a mod to user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "mod, m",
						Value: "",
						Usage: "Mod ID or slug to append",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, UserModAppend)
				},
			},
			{
				Name:      "mod-remove",
				Usage:     "Remove a mod from user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "mod, m",
						Value: "",
						Usage: "Mod ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, UserModRemove)
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
						Usage: "User ID or slug to list packs",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, UserPackList)
				},
			},
			{
				Name:      "pack-append",
				Usage:     "Append a pack to user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "Pack ID or slug to append",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, UserPackAppend)
				},
			},
			{
				Name:      "pack-remove",
				Usage:     "Remove a pack from user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "Pack ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, UserPackRemove)
				},
			},
		},
	}
}

// UserList provides the sub-command to list all users.
func UserList(c *cli.Context, client solder.ClientAPI) error {
	records, err := client.UserList()

	if err != nil {
		return err
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"ID", "Username", "Email"})

	for _, record := range records {
		table.Append(
			[]string{
				strconv.FormatInt(record.ID, 10),
				record.Username,
				record.Email,
			},
		)
	}

	table.Render()
	return nil
}

// UserShow provides the sub-command to show user details.
func UserShow(c *cli.Context, client solder.ClientAPI) error {
	record, err := client.UserGet(
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
			"Username",
			record.Username,
		},
	)

	table.Append(
		[]string{
			"Email",
			record.Email,
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

// UserDelete provides the sub-command to delete a user.
func UserDelete(c *cli.Context, client solder.ClientAPI) error {
	err := client.UserDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// UserUpdate provides the sub-command to update a user.
func UserUpdate(c *cli.Context, client solder.ClientAPI) error {
	record, err := client.UserGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	changed := false

	if val := c.String("slug"); c.IsSet("slug") && val != record.Slug {
		record.Slug = val
		changed = true
	}

	if val := c.String("username"); c.IsSet("username") && val != record.Username {
		record.Username = val
		changed = true
	}

	if val := c.String("email"); c.IsSet("email") && val != record.Email {
		record.Email = val
		changed = true
	}

	if val := c.String("password"); c.IsSet("password") {
		record.Password = val
		changed = true
	}

	if changed {
		_, patch := client.UserPatch(
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

// UserCreate provides the sub-command to create a user.
func UserCreate(c *cli.Context, client solder.ClientAPI) error {
	record := &solder.User{}

	if val := c.String("slug"); c.IsSet("slug") && val != "" {
		record.Slug = val
	}

	if val := c.String("username"); c.IsSet("username") && val != "" {
		record.Username = val
	} else {
		return fmt.Errorf("You must provide an username.")
	}

	if val := c.String("email"); c.IsSet("email") && val != "" {
		record.Email = val
	} else {
		return fmt.Errorf("You must provide an email.")
	}

	if val := c.String("password"); c.IsSet("password") && val != "" {
		record.Password = val
	} else {
		return fmt.Errorf("You must provide a password.")
	}

	_, err := client.UserPost(
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// UserModList provides the sub-command to list mods of the user.
func UserModList(c *cli.Context, client solder.ClientAPI) error {
	records, err := client.UserModList(
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

// UserModAppend provides the sub-command to append a mod to the user.
func UserModAppend(c *cli.Context, client solder.ClientAPI) error {
	err := client.UserModAppend(
		GetIdentifierParam(c),
		GetModParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to user\n")
	return nil
}

// UserModRemove provides the sub-command to remove a mod from the user.
func UserModRemove(c *cli.Context, client solder.ClientAPI) error {
	err := client.UserModDelete(
		GetIdentifierParam(c),
		GetModParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from user\n")
	return nil
}

// UserPackList provides the sub-command to list packs of the user.
func UserPackList(c *cli.Context, client solder.ClientAPI) error {
	records, err := client.UserPackList(
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

// UserPackAppend provides the sub-command to append a pack to the user.
func UserPackAppend(c *cli.Context, client solder.ClientAPI) error {
	err := client.UserPackAppend(
		GetIdentifierParam(c),
		GetPackParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to user\n")
	return nil
}

// UserPackRemove provides the sub-command to remove a pack from the user.
func UserPackRemove(c *cli.Context, client solder.ClientAPI) error {
	err := client.UserPackDelete(
		GetIdentifierParam(c),
		GetPackParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from user\n")
	return nil
}

func userPermissionFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "users-display",
			Usage: "Permit display users",
		},
		cli.BoolFlag{
			Name:  "no-users-display",
			Usage: "Deny display users",
		},
		cli.BoolFlag{
			Name:  "users-change",
			Usage: "Permit change users",
		},
		cli.BoolFlag{
			Name:  "no-users-change",
			Usage: "Deny change users",
		},
		cli.BoolFlag{
			Name:  "users-delete",
			Usage: "Permit delete users",
		},
		cli.BoolFlag{
			Name:  "no-users-delete",
			Usage: "Deny delete users",
		},
		cli.BoolFlag{
			Name:  "keys-display",
			Usage: "Permit display keys",
		},
		cli.BoolFlag{
			Name:  "no-keys-display",
			Usage: "Deny display keys",
		},
		cli.BoolFlag{
			Name:  "keys-change",
			Usage: "Permit change keys",
		},
		cli.BoolFlag{
			Name:  "no-keys-change",
			Usage: "Deny change keys",
		},
		cli.BoolFlag{
			Name:  "keys-delete",
			Usage: "Permit delete keys",
		},
		cli.BoolFlag{
			Name:  "no-keys-delete",
			Usage: "Deny delete keys",
		},
		cli.BoolFlag{
			Name:  "clients-display",
			Usage: "Permit display clients",
		},
		cli.BoolFlag{
			Name:  "no-clients-display",
			Usage: "Deny display clients",
		},
		cli.BoolFlag{
			Name:  "clients-change",
			Usage: "Permit change clients",
		},
		cli.BoolFlag{
			Name:  "no-clients-change",
			Usage: "Deny change clients",
		},
		cli.BoolFlag{
			Name:  "clients-delete",
			Usage: "Permit delete clients",
		},
		cli.BoolFlag{
			Name:  "no-clients-delete",
			Usage: "Deny delete clients",
		},
		cli.BoolFlag{
			Name:  "packs-display",
			Usage: "Permit display packs",
		},
		cli.BoolFlag{
			Name:  "no-packs-display",
			Usage: "Deny display packs",
		},
		cli.BoolFlag{
			Name:  "packs-change",
			Usage: "Permit change packs",
		},
		cli.BoolFlag{
			Name:  "no-packs-change",
			Usage: "Deny change packs",
		},
		cli.BoolFlag{
			Name:  "packs-delete",
			Usage: "Permit delete packs",
		},
		cli.BoolFlag{
			Name:  "no-packs-delete",
			Usage: "Deny delete packs",
		},
		cli.BoolFlag{
			Name:  "mods-display",
			Usage: "Permit display mods",
		},
		cli.BoolFlag{
			Name:  "no-mods-display",
			Usage: "Deny display mods",
		},
		cli.BoolFlag{
			Name:  "mods-change",
			Usage: "Permit change mods",
		},
		cli.BoolFlag{
			Name:  "no-mods-change",
			Usage: "Deny change mods",
		},
		cli.BoolFlag{
			Name:  "mods-delete",
			Usage: "Permit delete mods",
		},
		cli.BoolFlag{
			Name:  "no-mods-delete",
			Usage: "Deny delete mods",
		},
	}
}
