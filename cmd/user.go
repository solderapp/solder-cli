package cmd

import (
	"github.com/codegangsta/cli"
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
						Usage: "Version ID or slug to show",
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
						Usage: "Version ID or slug to show",
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
						Usage: "Version ID or slug to show",
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
func UserList(c *cli.Context, client solder.Client) error {
	return nil
}

// UserShow provides the sub-command to show user details.
func UserShow(c *cli.Context, client solder.Client) error {
	return nil
}

// UserUpdate provides the sub-command to update a user.
func UserUpdate(c *cli.Context, client solder.Client) error {
	return nil
}

// UserDelete provides the sub-command to delete a user.
func UserDelete(c *cli.Context, client solder.Client) error {
	return nil
}

// UserCreate provides the sub-command to create a user.
func UserCreate(c *cli.Context, client solder.Client) error {
	return nil
}
