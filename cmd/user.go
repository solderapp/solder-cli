package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"text/template"

	"github.com/kleister/kleister-go/kleister"
	"github.com/urfave/cli"
)

// userFuncMap provides user template helper functions.
var userFuncMap = template.FuncMap{}

// tmplUserList represents a row within forge listing.
var tmplUserList = "Slug: \x1b[33m{{ .Slug }}\x1b[0m" + `
ID: {{ .ID }}
Username: {{ .Username }}
`

// tmplUserShow represents a user within details view.
var tmplUserShow = "Slug: \x1b[33m{{ .Slug }}\x1b[0m" + `
ID: {{ .ID }}
Username: {{ .Username }}
Email: {{ .Email }}
Active: {{ .Active }}
Admin: {{ .Admin }}{{with .Teams}}
Teams: {{ teamList . }}{{end}}{{with .Packs}}
Packs: {{ packList . }}{{end}}{{with .Mods}}
Mods: {{ modList . }}{{end}}
Created: {{ .CreatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
Updated: {{ .UpdatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
`

// tmplUserModList represents a row within user mod listing.
var tmplUserModList = "Slug: \x1b[33m{{ .Mod.Slug }}\x1b[0m" + `
ID: {{ .Mod.ID }}
Name: {{ .Mod.Name }}
`

// tmplUserPackList represents a row within user pack listing.
var tmplUserPackList = "Slug: \x1b[33m{{ .Pack.Slug }}\x1b[0m" + `
ID: {{ .Pack.ID }}
Name: {{ .Pack.Name }}
`

// tmplUserTeamList represents a row within user team listing.
var tmplUserTeamList = "Slug: \x1b[33m{{ .Team.Slug }}\x1b[0m" + `
ID: {{ .Team.ID }}
Name: {{ .Team.Name }}
`

// User provides the sub-command for the user API.
func User() cli.Command {
	return cli.Command{
		Name:  "user",
		Usage: "User related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all users",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "format",
						Value: tmplUserList,
						Usage: "Custom output format",
					},
					cli.BoolFlag{
						Name:  "json",
						Usage: "Print in JSON format",
					},
					cli.BoolFlag{
						Name:  "xml",
						Usage: "Print in XML format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserList)
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
					cli.StringFlag{
						Name:  "format",
						Value: tmplUserShow,
						Usage: "Custom output format",
					},
					cli.BoolFlag{
						Name:  "json",
						Usage: "Print in JSON format",
					},
					cli.BoolFlag{
						Name:  "xml",
						Usage: "Print in XML format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserShow)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
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
					cli.BoolFlag{
						Name:  "active",
						Usage: "Mark user as active",
					},
					cli.BoolFlag{
						Name:  "blocked",
						Usage: "Mark user as blocked",
					},
					cli.BoolFlag{
						Name:  "admin",
						Usage: "Mark user as admin",
					},
					cli.BoolFlag{
						Name:  "user",
						Usage: "Mark user as user",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserUpdate)
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
				Action: func(c *cli.Context) error {
					return Handle(c, UserDelete)
				},
			},
			{
				Name:      "create",
				Usage:     "Create a user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
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
					cli.BoolFlag{
						Name:  "active",
						Usage: "Mark user as active",
					},
					cli.BoolFlag{
						Name:  "blocked",
						Usage: "Mark user as blocked",
					},
					cli.BoolFlag{
						Name:  "admin",
						Usage: "Mark user as admin",
					},
					cli.BoolFlag{
						Name:  "user",
						Usage: "Mark user as user",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserCreate)
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
					cli.StringFlag{
						Name:  "format",
						Value: tmplUserModList,
						Usage: "Custom output format",
					},
					cli.BoolFlag{
						Name:  "json",
						Usage: "Print in JSON format",
					},
					cli.BoolFlag{
						Name:  "xml",
						Usage: "Print in XML format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserModList)
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
				Action: func(c *cli.Context) error {
					return Handle(c, UserModAppend)
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
				Action: func(c *cli.Context) error {
					return Handle(c, UserModRemove)
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
					cli.StringFlag{
						Name:  "format",
						Value: tmplUserPackList,
						Usage: "Custom output format",
					},
					cli.BoolFlag{
						Name:  "json",
						Usage: "Print in JSON format",
					},
					cli.BoolFlag{
						Name:  "xml",
						Usage: "Print in XML format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserPackList)
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
				Action: func(c *cli.Context) error {
					return Handle(c, UserPackAppend)
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
				Action: func(c *cli.Context) error {
					return Handle(c, UserPackRemove)
				},
			},
			{
				Name:      "team-list",
				Usage:     "List assigned teams",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to list teams",
					},
					cli.StringFlag{
						Name:  "format",
						Value: tmplUserTeamList,
						Usage: "Custom output format",
					},
					cli.BoolFlag{
						Name:  "json",
						Usage: "Print in JSON format",
					},
					cli.BoolFlag{
						Name:  "xml",
						Usage: "Print in XML format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserTeamList)
				},
			},
			{
				Name:      "team-append",
				Usage:     "Append a team to user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "team, p",
						Value: "",
						Usage: "Team ID or slug to append",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserTeamAppend)
				},
			},
			{
				Name:      "team-remove",
				Usage:     "Remove a team from user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "team, p",
						Value: "",
						Usage: "Team ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserTeamRemove)
				},
			},
		},
	}
}

// UserList provides the sub-command to list all users.
func UserList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.UserList()

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("Conflict, you can only use JSON or XML at once!")
	}

	if c.Bool("xml") {
		res, err := xml.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if c.Bool("json") {
		res, err := json.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		userFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	for _, record := range records {
		err := tmpl.Execute(os.Stdout, record)

		if err != nil {
			return err
		}
	}

	return nil
}

// UserShow provides the sub-command to show user details.
func UserShow(c *cli.Context, client kleister.ClientAPI) error {
	record, err := client.UserGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("Conflict, you can only use JSON or XML at once!")
	}

	if c.Bool("xml") {
		res, err := xml.MarshalIndent(record, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if c.Bool("json") {
		res, err := json.MarshalIndent(record, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		userFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, record)
}

// UserDelete provides the sub-command to delete a user.
func UserDelete(c *cli.Context, client kleister.ClientAPI) error {
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
func UserUpdate(c *cli.Context, client kleister.ClientAPI) error {
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

	if c.IsSet("active") && c.IsSet("blocked") {
		return fmt.Errorf("Conflict, you can mark it only active OR blocked!")
	}

	if c.IsSet("active") {
		record.Active = true
		changed = true
	}

	if c.IsSet("blocked") {
		record.Active = false
		changed = true
	}

	if c.IsSet("admin") && c.IsSet("user") {
		return fmt.Errorf("Conflict, you can mark it only admin OR user!")
	}

	if c.IsSet("admin") {
		record.Admin = true
		changed = true
	}

	if c.IsSet("user") {
		record.Admin = false
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
func UserCreate(c *cli.Context, client kleister.ClientAPI) error {
	record := &kleister.User{}

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

	if c.IsSet("active") && c.IsSet("blocked") {
		return fmt.Errorf("Conflict, you can mark it only active OR blocked!")
	}

	if c.IsSet("active") {
		record.Active = true
	}

	if c.IsSet("blocked") {
		record.Active = false
	}

	if c.IsSet("admin") && c.IsSet("user") {
		return fmt.Errorf("Conflict, you can mark it only admin OR user!")
	}

	if c.IsSet("admin") {
		record.Admin = true
	}

	if c.IsSet("user") {
		record.Admin = false
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
func UserModList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.UserModList(
		kleister.UserModParams{
			User: GetIdentifierParam(c),
		},
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("Conflict, you can only use JSON or XML at once!")
	}

	if c.Bool("xml") {
		res, err := xml.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if c.Bool("json") {
		res, err := json.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		userFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	for _, record := range records {
		err := tmpl.Execute(os.Stdout, record)

		if err != nil {
			return err
		}
	}

	return nil
}

// UserModAppend provides the sub-command to append a mod to the user.
func UserModAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.UserModAppend(
		kleister.UserModParams{
			User: GetIdentifierParam(c),
			Mod:  GetModParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to user\n")
	return nil
}

// UserModRemove provides the sub-command to remove a mod from the user.
func UserModRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.UserModDelete(
		kleister.UserModParams{
			User: GetIdentifierParam(c),
			Mod:  GetModParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from user\n")
	return nil
}

// UserPackList provides the sub-command to list packs of the user.
func UserPackList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.UserPackList(
		kleister.UserPackParams{
			User: GetIdentifierParam(c),
		},
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("Conflict, you can only use JSON or XML at once!")
	}

	if c.Bool("xml") {
		res, err := xml.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if c.Bool("json") {
		res, err := json.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		userFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	for _, record := range records {
		err := tmpl.Execute(os.Stdout, record)

		if err != nil {
			return err
		}
	}

	return nil
}

// UserPackAppend provides the sub-command to append a pack to the user.
func UserPackAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.UserPackAppend(
		kleister.UserPackParams{
			User: GetIdentifierParam(c),
			Pack: GetPackParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to user\n")
	return nil
}

// UserPackRemove provides the sub-command to remove a pack from the user.
func UserPackRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.UserPackDelete(
		kleister.UserPackParams{
			User: GetIdentifierParam(c),
			Pack: GetPackParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from user\n")
	return nil
}

// UserTeamList provides the sub-command to list teams of the user.
func UserTeamList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.UserTeamList(
		kleister.UserTeamParams{
			User: GetIdentifierParam(c),
		},
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("Conflict, you can only use JSON or XML at once!")
	}

	if c.Bool("xml") {
		res, err := xml.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if c.Bool("json") {
		res, err := json.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		globalFuncMap,
	).Funcs(
		userFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	for _, record := range records {
		err := tmpl.Execute(os.Stdout, record)

		if err != nil {
			return err
		}
	}

	return nil
}

// UserTeamAppend provides the sub-command to append a team to the user.
func UserTeamAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.UserTeamAppend(
		kleister.UserTeamParams{
			User: GetIdentifierParam(c),
			Team: GetTeamParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to user\n")
	return nil
}

// UserTeamRemove provides the sub-command to remove a team from the user.
func UserTeamRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.UserTeamDelete(
		kleister.UserTeamParams{
			User: GetIdentifierParam(c),
			Team: GetTeamParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from user\n")
	return nil
}
