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
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all users",
				Action: func(c *cli.Context) {
					Handle(c, UserList)
				},
			},
			{
				Name:  "show",
				Usage: "Display a user",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "User ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, UserShow)
				},
			},
			{
				Name:  "update",
				Usage: "Update a user",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "User ID or slug to show",
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
				Action: func(c *cli.Context) {
					Handle(c, UserUpdate)
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"rm"},
				Usage:   "Delete a user",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "User ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, UserDelete)
				},
			},
			{
				Name:  "create",
				Usage: "Create a user",
				Action: func(c *cli.Context) {
					Handle(c, UserCreate)
				},
			},
		},
	}
}

// UserList provides the sub-command to list all users.
func UserList(c *cli.Context, client solder.API) error {
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
func UserShow(c *cli.Context, client solder.API) error {
	record, err := client.UserGet(
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
			[]string{"Username", record.Username},
			[]string{"Email", record.Email},
			[]string{"Created", record.CreatedAt.Format(time.UnixDate)},
			[]string{"Updated", record.UpdatedAt.Format(time.UnixDate)},
		},
	)

	table.Render()
	return nil
}

// UserDelete provides the sub-command to delete a user.
func UserDelete(c *cli.Context, client solder.API) error {
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
func UserUpdate(c *cli.Context, client solder.API) error {
	record, err := client.UserGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	if val := c.String("slug"); val != record.Slug {
		record.Slug = val
	}

	if val := c.String("username"); val != record.Username {
		record.Username = val
	}

	if val := c.String("email"); val != record.Email {
		record.Email = val
	}

	if val := c.String("password"); val != "" {
		record.Password = val
	}

	_, patch := client.UserPatch(record)

	if patch != nil {
		return patch
	}

	fmt.Fprintf(os.Stderr, "Successfully updated\n")
	return nil
}

// UserCreate provides the sub-command to create a user.
func UserCreate(c *cli.Context, client solder.API) error {
	record := &solder.User{}

	if val := c.String("slug"); val != "" {
		record.Slug = val
	}

	if val := c.String("username"); val != "" {
		record.Username = val
	} else {
		return fmt.Errorf("You must provide an username.")
	}

	if val := c.String("email"); val != "" {
		record.Email = val
	} else {
		return fmt.Errorf("You must provide an email.")
	}

	if val := c.String("password"); val != "" {
		record.Password = val
	} else {
		return fmt.Errorf("You must provide a password.")
	}

	_, err := client.UserPost(record)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}
