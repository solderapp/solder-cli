package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/kleister/kleister-go/kleister"
	"github.com/urfave/cli"
)

// PackFuncMap provides template helper functions.
var packFuncMap = template.FuncMap{
	"clientList": func(s []*kleister.Client) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"userList": func(s []*kleister.User) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"teamList": func(s []*kleister.Team) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
}

// tmplPackList represents a row within forge listing.
var tmplPackList = "Slug: \x1b[33m{{ .Slug }}\x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}
`

// tmplPackShow represents a pack within details view.
var tmplPackShow = "Slug: \x1b[33m{{ .Slug }}\x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}{{with .Website}}
Website: {{ . }}{{end}}{{with .Recommended}}
Recommended: {{ . }}{{end}}{{with .Latest}}
Latest: {{ . }}{{end}}{{with .Icon}}
Icon: {{ . }}{{end}}{{with .Logo}}
Logo: {{ . }}{{end}}{{with .Background}}
Background: {{ . }}{{end}}
Published: {{ .Published }}
Private: {{ .Private }}{{with .Clients}}
Clients: {{ clientList . }}{{end}}{{with .Users}}
Users: {{ userList . }}{{end}}{{with .Teams}}
Teams: {{ teamList . }}{{end}}
Created: {{ .CreatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
Updated: {{ .UpdatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
`

// tmplPackClientList represents a row within pack client listing.
var tmplPackClientList = "Slug: \x1b[33m{{ .Slug }}\x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}
`

// tmplPackUserList represents a row within pack user listing.
var tmplPackUserList = "Slug: \x1b[33m{{ .Slug }}\x1b[0m" + `
ID: {{ .ID }}
Username: {{ .Username }}
`

// tmplPackTeamList represents a row within pack team listing.
var tmplPackTeamList = "Slug: \x1b[33m{{ .Slug }}\x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}
`

// Pack provides the sub-command for the pack API.
func Pack() cli.Command {
	return cli.Command{
		Name:  "pack",
		Usage: "Pack related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all packs",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "format",
						Value: tmplPackList,
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
					return Handle(c, PackList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a pack",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Pack ID or slug to show",
					},
					cli.StringFlag{
						Name:  "format",
						Value: tmplPackShow,
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
					return Handle(c, PackShow)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a pack",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Pack ID or slug to update",
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
						Name:  "website",
						Value: "",
						Usage: "Provide a website",
					},
					cli.StringFlag{
						Name:  "recommended",
						Value: "",
						Usage: "Recommended build ID or slug",
					},
					cli.StringFlag{
						Name:  "latest",
						Value: "",
						Usage: "Latest build ID or slug",
					},
					cli.StringFlag{
						Name:  "icon-url",
						Value: "",
						Usage: "Provide an icon URL",
					},
					cli.StringFlag{
						Name:  "icon-path",
						Value: "",
						Usage: "Provide an icon path",
					},
					cli.StringFlag{
						Name:  "logo-url",
						Value: "",
						Usage: "Provide a logo URL",
					},
					cli.StringFlag{
						Name:  "logo-path",
						Value: "",
						Usage: "Provide a logo path",
					},
					cli.StringFlag{
						Name:  "bg-url",
						Value: "",
						Usage: "Provide a background URL",
					},
					cli.StringFlag{
						Name:  "bg-path",
						Value: "",
						Usage: "Provide a background path",
					},
					cli.BoolFlag{
						Name:  "published",
						Usage: "Mark pack published",
					},
					cli.BoolFlag{
						Name:  "hidden",
						Usage: "Mark pack hidden",
					},
					cli.BoolFlag{
						Name:  "private",
						Usage: "Mark pack private",
					},
					cli.BoolFlag{
						Name:  "public",
						Usage: "Mark pack public",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, PackUpdate)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a pack",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Pack ID or slug to delete",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, PackDelete)
				},
			},
			{
				Name:      "create",
				Usage:     "Create a pack",
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
						Name:  "website",
						Value: "",
						Usage: "Provide a website",
					},
					cli.StringFlag{
						Name:  "icon-url",
						Value: "",
						Usage: "Provide an icon URL",
					},
					cli.StringFlag{
						Name:  "icon-path",
						Value: "",
						Usage: "Provide an icon path",
					},
					cli.StringFlag{
						Name:  "logo-url",
						Value: "",
						Usage: "Provide a logo URL",
					},
					cli.StringFlag{
						Name:  "logo-path",
						Value: "",
						Usage: "Provide a logo path",
					},
					cli.StringFlag{
						Name:  "bg-url",
						Value: "",
						Usage: "Provide a background URL",
					},
					cli.StringFlag{
						Name:  "bg-path",
						Value: "",
						Usage: "Provide a background path",
					},
					cli.BoolFlag{
						Name:  "published",
						Usage: "Mark pack published",
					},
					cli.BoolFlag{
						Name:  "hidden",
						Usage: "Mark pack hidden",
					},
					cli.BoolFlag{
						Name:  "private",
						Usage: "Mark pack private",
					},
					cli.BoolFlag{
						Name:  "public",
						Usage: "Mark pack public",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, PackCreate)
				},
			},
			{
				Name:      "client-list",
				Usage:     "List assigned clients",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Pack ID or slug to list clients",
					},
					cli.StringFlag{
						Name:  "format",
						Value: tmplPackClientList,
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
					return Handle(c, PackClientList)
				},
			},
			{
				Name:      "client-append",
				Usage:     "Append a client to pack",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Pack ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "client, c",
						Value: "",
						Usage: "Client ID or slug to append",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, PackClientAppend)
				},
			},
			{
				Name:      "client-remove",
				Usage:     "Remove a client from pack",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Pack ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "client, c",
						Value: "",
						Usage: "Client ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, PackClientRemove)
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
						Usage: "Pack ID or slug to list users",
					},
					cli.StringFlag{
						Name:  "format",
						Value: tmplPackUserList,
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
					return Handle(c, PackUserList)
				},
			},
			{
				Name:      "user-append",
				Usage:     "Append a user to pack",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Pack ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "user, u",
						Value: "",
						Usage: "User ID or slug to append",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, PackUserAppend)
				},
			},
			{
				Name:      "user-remove",
				Usage:     "Remove a user from pack",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Pack ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "user, u",
						Value: "",
						Usage: "User ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, PackUserRemove)
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
						Usage: "Pack ID or slug to list teams",
					},
					cli.StringFlag{
						Name:  "format",
						Value: tmplPackTeamList,
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
					return Handle(c, PackTeamList)
				},
			},
			{
				Name:      "team-append",
				Usage:     "Append a team to pack",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Pack ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "team, t",
						Value: "",
						Usage: "Team ID or slug to append",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, PackTeamAppend)
				},
			},
			{
				Name:      "team-remove",
				Usage:     "Remove a team from pack",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Pack ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "team, t",
						Value: "",
						Usage: "Team ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, PackTeamRemove)
				},
			},
		},
	}
}

// PackList provides the sub-command to list all packs.
func PackList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.PackList()

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
		packFuncMap,
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

// PackShow provides the sub-command to show pack details.
func PackShow(c *cli.Context, client kleister.ClientAPI) error {
	record, err := client.PackGet(
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
		packFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, record)
}

// PackDelete provides the sub-command to delete a pack.
func PackDelete(c *cli.Context, client kleister.ClientAPI) error {
	err := client.PackDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// PackUpdate provides the sub-command to update a pack.
func PackUpdate(c *cli.Context, client kleister.ClientAPI) error {
	record, err := client.PackGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	changed := false

	if c.IsSet("recommended") {
		if match, _ := regexp.MatchString("^([0-9]+)$", c.String("recommended")); match {
			if val, err := strconv.ParseInt(c.String("recommended"), 10, 64); err == nil && val != record.RecommendedID {
				record.RecommendedID = val
				changed = true
			}
		} else {
			if c.String("recommended") != "" {
				related, err := client.BuildGet(
					GetIdentifierParam(c),
					c.String("recommended"),
				)

				if err != nil {
					return err
				}

				if related.ID != record.RecommendedID {
					record.RecommendedID = related.ID
					changed = true
				}
			}
		}
	}

	if c.IsSet("latest") {
		if match, _ := regexp.MatchString("^([0-9]+)$", c.String("latest")); match {
			if val, err := strconv.ParseInt(c.String("latest"), 10, 64); err == nil && val != record.LatestID {
				record.LatestID = val
				changed = true
			}
		} else {
			if c.String("latest") != "" {
				related, err := client.BuildGet(
					GetIdentifierParam(c),
					c.String("latest"),
				)

				if err != nil {
					return err
				}

				if related.ID != record.LatestID {
					record.LatestID = related.ID
					changed = true
				}
			}
		}
	}

	if val := c.String("name"); c.IsSet("name") && val != record.Name {
		record.Name = val
		changed = true
	}

	if val := c.String("slug"); c.IsSet("slug") && val != record.Slug {
		record.Slug = val
		changed = true
	}

	if val := c.String("website"); c.IsSet("website") && val != record.Website {
		record.Website = val
		changed = true
	}

	if val := c.String("icon-url"); c.IsSet("icon-url") && val != "" {
		err := record.DownloadIcon(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to download and encode icon.")
		}

		changed = true
	}

	if val := c.String("icon-path"); c.IsSet("icon-path") && val != "" {
		err := record.EncodeIcon(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to encode icon.")
		}

		changed = true
	}

	if val := c.String("logo-url"); c.IsSet("logo-url") && val != "" {
		err := record.DownloadLogo(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to download and encode logo.")
		}

		changed = true
	}

	if val := c.String("logo-path"); c.IsSet("logo-path") && val != "" {
		err := record.EncodeLogo(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to encode logo.")
		}

		changed = true
	}

	if val := c.String("bg-url"); c.IsSet("bg-url") && val != "" {
		err := record.DownloadBackground(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to download and encode background.")
		}

		changed = true
	}

	if val := c.String("bg-path"); c.IsSet("bg-path") && val != "" {
		err := record.EncodeBackground(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to encode background.")
		}

		changed = true
	}

	if c.IsSet("published") && c.IsSet("hidden") {
		return fmt.Errorf("Conflict, you can mark it only published OR hidden!")
	}

	if c.IsSet("published") {
		record.Published = true
		changed = true
	}

	if c.IsSet("hidden") {
		record.Published = false
		changed = true
	}

	if c.IsSet("private") && c.IsSet("public") {
		return fmt.Errorf("Conflict, you can mark it only private OR public!")
	}

	if c.IsSet("private") {
		record.Private = true
		changed = true
	}

	if c.IsSet("public") {
		record.Private = false
		changed = true
	}

	if changed {
		_, patch := client.PackPatch(
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

// PackCreate provides the sub-command to create a pack.
func PackCreate(c *cli.Context, client kleister.ClientAPI) error {
	record := &kleister.Pack{}

	if val := c.String("name"); c.IsSet("name") && val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("You must provide a name.")
	}

	if val := c.String("slug"); c.IsSet("slug") && val != "" {
		record.Slug = val
	}

	if val := c.String("website"); c.IsSet("website") && val != "" {
		record.Website = val
	}

	if val := c.String("icon-url"); c.IsSet("icon-url") && val != "" {
		err := record.DownloadIcon(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to download and encode icon.")
		}
	}

	if val := c.String("icon-path"); c.IsSet("icon-path") && val != "" {
		err := record.EncodeIcon(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to encode icon.")
		}
	}

	if val := c.String("logo-url"); c.IsSet("logo-url") && val != "" {
		err := record.DownloadLogo(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to download and encode logo.")
		}
	}

	if val := c.String("logo-path"); c.IsSet("logo-path") && val != "" {
		err := record.EncodeLogo(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to encode logo.")
		}
	}

	if val := c.String("bg-url"); c.IsSet("bg-url") && val != "" {
		err := record.DownloadBackground(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to download and encode background.")
		}
	}

	if val := c.String("bg-path"); c.IsSet("bg-path") && val != "" {
		err := record.EncodeBackground(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to encode background.")
		}
	}

	if c.IsSet("published") && c.IsSet("hidden") {
		return fmt.Errorf("Conflict, you can mark it only published OR hidden!")
	}

	if c.IsSet("published") {
		record.Published = true
	}

	if c.IsSet("hidden") {
		record.Published = false
	}

	if c.IsSet("private") && c.IsSet("public") {
		return fmt.Errorf("Conflict, you can mark it only private OR public!")
	}

	if c.IsSet("private") {
		record.Private = true
	}

	if c.IsSet("public") {
		record.Private = false
	}

	_, err := client.PackPost(
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// PackClientList provides the sub-command to list packs of the pack.
func PackClientList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.PackClientList(
		kleister.PackClientParams{
			Pack: GetIdentifierParam(c),
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
		packFuncMap,
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

// PackClientAppend provides the sub-command to append a client to the pack.
func PackClientAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.PackClientAppend(
		kleister.PackClientParams{
			Pack:   GetIdentifierParam(c),
			Client: GetClientParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to pack\n")
	return nil
}

// PackClientRemove provides the sub-command to remove a client from the pack.
func PackClientRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.PackClientDelete(
		kleister.PackClientParams{
			Pack:   GetIdentifierParam(c),
			Client: GetClientParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from pack\n")
	return nil
}

// PackUserList provides the sub-command to list users of the pack.
func PackUserList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.PackUserList(
		kleister.PackUserParams{
			Pack: GetIdentifierParam(c),
			User: GetUserParam(c),
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
		packFuncMap,
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

// PackUserAppend provides the sub-command to append a user to the pack.
func PackUserAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.PackUserAppend(
		kleister.PackUserParams{
			Pack: GetIdentifierParam(c),
			User: GetUserParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to pack\n")
	return nil
}

// PackUserRemove provides the sub-command to remove a user from the pack.
func PackUserRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.PackUserDelete(
		kleister.PackUserParams{
			Pack: GetIdentifierParam(c),
			User: GetUserParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from pack\n")
	return nil
}

// PackTeamList provides the sub-command to list teams of the pack.
func PackTeamList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.PackTeamList(
		kleister.PackTeamParams{
			Pack: GetIdentifierParam(c),
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
		packFuncMap,
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

// PackTeamAppend provides the sub-command to append a team to the pack.
func PackTeamAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.PackTeamAppend(
		kleister.PackTeamParams{
			Pack: GetIdentifierParam(c),
			Team: GetTeamParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to team\n")
	return nil
}

// PackTeamRemove provides the sub-command to remove a team from the pack.
func PackTeamRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.PackTeamDelete(
		kleister.PackTeamParams{
			Pack: GetIdentifierParam(c),
			Team: GetTeamParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from team\n")
	return nil
}
