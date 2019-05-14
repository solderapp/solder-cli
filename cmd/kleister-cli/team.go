package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"text/template"

	"github.com/kleister/kleister-go/kleister"
	"gopkg.in/urfave/cli.v2"
)

// tmplTeamList represents a row within user listing.
var tmplTeamList = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}
`

// tmplTeamShow represents a user within details view.
var tmplTeamShow = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}{{with .Users}}
Users: {{ userlist . }}{{end}}{{with .Packs}}
Packs: {{ packlist . }}{{end}}{{with .Mods}}
Mods: {{ modlist . }}{{end}}
Created: {{ .CreatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
Updated: {{ .UpdatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
`

// tmplTeamUserList represents a row within team user listing.
var tmplTeamUserList = "Slug: \x1b[33m{{ .User.Slug }} \x1b[0m" + `
ID: {{ .User.ID }}
Username: {{ .User.Username }}
Permission: {{ .Perm }}
`

// tmplTeamPackList represents a row within team pack listing.
var tmplTeamPackList = "Slug: \x1b[33m{{ .Pack.Slug }} \x1b[0m" + `
ID: {{ .Pack.ID }}
Name: {{ .Pack.Name }}
Permission: {{ .Perm }}
`

// tmplTeamModList represents a row within team mod listing.
var tmplTeamModList = "Slug: \x1b[33m{{ .Mod.Slug }} \x1b[0m" + `
ID: {{ .Mod.ID }}
Name: {{ .Mod.Name }}
Permission: {{ .Perm }}
`

// Team provides the sub-command for the team API.
func Team() *cli.Command {
	return &cli.Command{
		Name:  "team",
		Usage: "Team related sub-commands",
		Subcommands: []*cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all teams",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "format",
						Value: tmplTeamList,
						Usage: "Custom output format",
					},
					&cli.BoolFlag{
						Name:  "json",
						Value: false,
						Usage: "Print in JSON format",
					},
					&cli.BoolFlag{
						Name:  "xml",
						Value: false,
						Usage: "Print in XML format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a team",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Team ID or slug to show",
					},
					&cli.StringFlag{
						Name:  "format",
						Value: tmplTeamShow,
						Usage: "Custom output format",
					},
					&cli.BoolFlag{
						Name:  "json",
						Value: false,
						Usage: "Print in JSON format",
					},
					&cli.BoolFlag{
						Name:  "xml",
						Value: false,
						Usage: "Print in XML format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamShow)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a team",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Team ID or slug to show",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamDelete)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a team",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Team ID or slug to update",
					},
					&cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Provide a slug",
					},
					&cli.StringFlag{
						Name:  "name",
						Value: "",
						Usage: "Provide a name",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamUpdate)
				},
			},
			{
				Name:      "create",
				Usage:     "Create a team",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Provide a slug",
					},
					&cli.StringFlag{
						Name:  "name",
						Value: "",
						Usage: "Provide a name",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamCreate)
				},
			},
			{
				Name:  "user",
				Usage: "User assignments",
				Subcommands: []*cli.Command{
					{
						Name:      "list",
						Aliases:   []string{"ls"},
						Usage:     "List assigned users",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to list users",
							},
							&cli.StringFlag{
								Name:  "format",
								Value: tmplTeamUserList,
								Usage: "Custom output format",
							},
							&cli.BoolFlag{
								Name:  "json",
								Value: false,
								Usage: "Print in JSON format",
							},
							&cli.BoolFlag{
								Name:  "xml",
								Value: false,
								Usage: "Print in XML format",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamUserList)
						},
					},
					{
						Name:      "append",
						Usage:     "Append a user to team",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to append to",
							},
							&cli.StringFlag{
								Name:  "user, u",
								Value: "",
								Usage: "User ID or slug to append",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the user, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamUserAppend)
						},
					},
					{
						Name:      "perm",
						Usage:     "Update team user permissions",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to update",
							},
							&cli.StringFlag{
								Name:  "user, u",
								Value: "",
								Usage: "User ID or slug to update",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the user, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamUserPerm)
						},
					},
					{
						Name:      "remove",
						Aliases:   []string{"rm"},
						Usage:     "Remove a user from team",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to remove from",
							},
							&cli.StringFlag{
								Name:  "user, u",
								Value: "",
								Usage: "User ID or slug to remove",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamUserRemove)
						},
					},
				},
			},
			{
				Name:  "pack",
				Usage: "Pack assignments",
				Subcommands: []*cli.Command{
					{
						Name:      "list",
						Aliases:   []string{"ls"},
						Usage:     "List assigned packs",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to list packs",
							},
							&cli.StringFlag{
								Name:  "format",
								Value: tmplTeamPackList,
								Usage: "Custom output format",
							},
							&cli.BoolFlag{
								Name:  "json",
								Value: false,
								Usage: "Print in JSON format",
							},
							&cli.BoolFlag{
								Name:  "xml",
								Value: false,
								Usage: "Print in XML format",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamPackList)
						},
					},
					{
						Name:      "append",
						Usage:     "Append a pack to team",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to append to",
							},
							&cli.StringFlag{
								Name:  "pack, u",
								Value: "",
								Usage: "Pack ID or slug to append",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the pack, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamPackAppend)
						},
					},
					{
						Name:      "perm",
						Usage:     "Update team pack permissions",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to update",
							},
							&cli.StringFlag{
								Name:  "pack, u",
								Value: "",
								Usage: "Pack ID or slug to update",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the pack, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamPackPerm)
						},
					},
					{
						Name:      "remove",
						Aliases:   []string{"rm"},
						Usage:     "Remove a pack from team",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to remove from",
							},
							&cli.StringFlag{
								Name:  "pack, u",
								Value: "",
								Usage: "Pack ID or slug to remove",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamPackRemove)
						},
					},
				},
			},
			{
				Name:  "mod",
				Usage: "Mod assignments",
				Subcommands: []*cli.Command{
					{
						Name:      "list",
						Aliases:   []string{"ls"},
						Usage:     "List assigned mods",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to list mods",
							},
							&cli.StringFlag{
								Name:  "format",
								Value: tmplTeamModList,
								Usage: "Custom output format",
							},
							&cli.BoolFlag{
								Name:  "json",
								Value: false,
								Usage: "Print in JSON format",
							},
							&cli.BoolFlag{
								Name:  "xml",
								Value: false,
								Usage: "Print in XML format",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamModList)
						},
					},
					{
						Name:      "append",
						Usage:     "Append a mod to team",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to append to",
							},
							&cli.StringFlag{
								Name:  "mod, u",
								Value: "",
								Usage: "Mod ID or slug to append",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the mod, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamModAppend)
						},
					},
					{
						Name:      "perm",
						Usage:     "Update team mod permissions",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to update",
							},
							&cli.StringFlag{
								Name:  "mod, u",
								Value: "",
								Usage: "Mod ID or slug to update",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the mod, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamModPerm)
						},
					},
					{
						Name:      "remove",
						Aliases:   []string{"rm"},
						Usage:     "Remove a mod from team",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to remove from",
							},
							&cli.StringFlag{
								Name:  "mod, u",
								Value: "",
								Usage: "Mod ID or slug to remove",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamModRemove)
						},
					},
				},
			},
		},
	}
}

// TeamList provides the sub-command to list all teams.
func TeamList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.TeamList()

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("conflict, you can only use json or xml at once")
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
		sprigFuncMap,
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

// TeamShow provides the sub-command to show team details.
func TeamShow(c *cli.Context, client kleister.ClientAPI) error {
	record, err := client.TeamGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("conflict, you can only use json or xml at once")
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
		sprigFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, record)
}

// TeamDelete provides the sub-command to delete a team.
func TeamDelete(c *cli.Context, client kleister.ClientAPI) error {
	err := client.TeamDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// TeamUpdate provides the sub-command to update a team.
func TeamUpdate(c *cli.Context, client kleister.ClientAPI) error {
	record, err := client.TeamGet(
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

	if val := c.String("name"); c.IsSet("name") && val != record.Name {
		record.Name = val
		changed = true
	}

	if changed {
		_, patch := client.TeamPatch(
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

// TeamCreate provides the sub-command to create a team.
func TeamCreate(c *cli.Context, client kleister.ClientAPI) error {
	record := &kleister.Team{}

	if val := c.String("slug"); c.IsSet("slug") && val != "" {
		record.Slug = val
	}

	if val := c.String("name"); c.IsSet("name") && val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("you must provide a name")
	}

	_, err := client.TeamPost(
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// TeamUserList provides the sub-command to list users of the team.
func TeamUserList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.TeamUserList(
		kleister.TeamUserParams{
			Team: GetIdentifierParam(c),
		},
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("conflict, you can only use json or xml at once")
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
		sprigFuncMap,
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

// TeamUserAppend provides the sub-command to append a user to the team.
func TeamUserAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.TeamUserAppend(
		kleister.TeamUserParams{
			Team: GetIdentifierParam(c),
			User: GetUserParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to user\n")
	return nil
}

// TeamUserPerm provides the sub-command to update team user permissions.
func TeamUserPerm(c *cli.Context, client kleister.ClientAPI) error {
	err := client.TeamUserPerm(
		kleister.TeamUserParams{
			Team: GetIdentifierParam(c),
			User: GetUserParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully updated permissions\n")
	return nil
}

// TeamUserRemove provides the sub-command to remove a user from the team.
func TeamUserRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.TeamUserDelete(
		kleister.TeamUserParams{
			Team: GetIdentifierParam(c),
			User: GetUserParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from user\n")
	return nil
}

// TeamPackList provides the sub-command to list packs of the team.
func TeamPackList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.TeamPackList(
		kleister.TeamPackParams{
			Team: GetIdentifierParam(c),
		},
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("conflict, you can only use json or xml at once")
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
		sprigFuncMap,
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

// TeamPackAppend provides the sub-command to append a pack to the team.
func TeamPackAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.TeamPackAppend(
		kleister.TeamPackParams{
			Team: GetIdentifierParam(c),
			Pack: GetPackParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to pack\n")
	return nil
}

// TeamPackPerm provides the sub-command to update team pack permissions.
func TeamPackPerm(c *cli.Context, client kleister.ClientAPI) error {
	err := client.TeamPackPerm(
		kleister.TeamPackParams{
			Team: GetIdentifierParam(c),
			Pack: GetPackParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully updated permissions\n")
	return nil
}

// TeamPackRemove provides the sub-command to remove a pack from the team.
func TeamPackRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.TeamPackDelete(
		kleister.TeamPackParams{
			Team: GetIdentifierParam(c),
			Pack: GetPackParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from pack\n")
	return nil
}

// TeamModList provides the sub-command to list mods of the team.
func TeamModList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.TeamModList(
		kleister.TeamModParams{
			Team: GetIdentifierParam(c),
		},
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("conflict, you can only use json or xml at once")
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
		sprigFuncMap,
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

// TeamModAppend provides the sub-command to append a mod to the team.
func TeamModAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.TeamModAppend(
		kleister.TeamModParams{
			Team: GetIdentifierParam(c),
			Mod:  GetModParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to mod\n")
	return nil
}

// TeamModPerm provides the sub-command to update team mod permissions.
func TeamModPerm(c *cli.Context, client kleister.ClientAPI) error {
	err := client.TeamModPerm(
		kleister.TeamModParams{
			Team: GetIdentifierParam(c),
			Mod:  GetModParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully updated permissions\n")
	return nil
}

// TeamModRemove provides the sub-command to remove a mod from the team.
func TeamModRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.TeamModDelete(
		kleister.TeamModParams{
			Team: GetIdentifierParam(c),
			Mod:  GetModParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from mod\n")
	return nil
}
