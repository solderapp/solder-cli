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

// buildFuncMap provides template helper functions.
var buildFuncMap = template.FuncMap{
	"versionList": func(s []*kleister.Version) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
}

// tmplBuildList represents a row within build listing.
var tmplBuildList = "Slug: \x1b[33m{{ .Slug }}\x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}
`

// tmplBuildShow represents a build within details view.
var tmplBuildShow = "Slug: \x1b[33m{{ .Slug }}\x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}{{with .Pack}}
Pack: {{ .Name }}{{end}}{{with .Minecraft}}
Minecraft: {{ . }}{{end}}{{with .Forge}}
Forge: {{ . }}{{end}}{{with .MinJava}}
Java: {{ . }}{{end}}{{with .MinMemory}}
Memory: {{ . }}{{end}}
Published: {{ .Published }}
Private: {{ .Private }}{{with .Versions}}
Versions: {{ versionList . }}{{end}}
Created: {{ .CreatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
Updated: {{ .UpdatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
`

// tmplBuildVersionList represents a row within build version listing.
var tmplBuildVersionList = "Slug: \x1b[33m{{ .Version.Slug }}\x1b[0m" + `
ID: {{ .Version.ID }}
Name: {{ .Version.Name }}
Mod: {{with .Version.Mod}}{{ . }}{{else}}n/a{{end}}
`

// Build provides the sub-command for the build API.
func Build() cli.Command {
	return cli.Command{
		Name:  "build",
		Usage: "Build related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all builds",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "ID or slug of the related pack",
					},
					cli.StringFlag{
						Name:  "format",
						Value: tmplBuildList,
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
					return Handle(c, BuildList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a build",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "ID or slug of the related pack",
					},
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Build ID or slug to show",
					},
					cli.StringFlag{
						Name:  "format",
						Value: tmplBuildShow,
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
					return Handle(c, BuildShow)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a build",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "ID or slug of the related pack",
					},
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Build ID or slug to update",
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
						Name:  "min-java",
						Value: "",
						Usage: "Minimal Java version",
					},
					cli.StringFlag{
						Name:  "min-memory",
						Value: "",
						Usage: "Minimal memory alloc",
					},
					cli.StringFlag{
						Name:  "minecraft",
						Value: "",
						Usage: "Provide a Minecraft ID or slug",
					},
					cli.StringFlag{
						Name:  "forge",
						Value: "",
						Usage: "Provide a Forge ID or slug",
					},
					cli.BoolFlag{
						Name:  "published",
						Usage: "Mark build published",
					},
					cli.BoolFlag{
						Name:  "hidden",
						Usage: "Mark pack hidden",
					},
					cli.BoolFlag{
						Name:  "private",
						Usage: "Mark build private",
					},
					cli.BoolFlag{
						Name:  "public",
						Usage: "Mark pack public",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, BuildUpdate)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a build",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "ID or slug of the related pack",
					},
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Build ID or slug to delete",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, BuildDelete)
				},
			},
			{
				Name:      "create",
				Usage:     "Create a build",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "ID or slug of the related pack",
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
						Name:  "min-java",
						Value: "",
						Usage: "Minimal Java version",
					},
					cli.StringFlag{
						Name:  "min-memory",
						Value: "",
						Usage: "Minimal memory alloc",
					},
					cli.StringFlag{
						Name:  "minecraft",
						Value: "",
						Usage: "Provide a Minecraft ID or slug",
					},
					cli.StringFlag{
						Name:  "forge",
						Value: "",
						Usage: "Provide a Forge ID or slug",
					},
					cli.BoolFlag{
						Name:  "published",
						Usage: "Mark build published",
					},
					cli.BoolFlag{
						Name:  "hidden",
						Usage: "Mark pack hidden",
					},
					cli.BoolFlag{
						Name:  "private",
						Usage: "Mark build private",
					},
					cli.BoolFlag{
						Name:  "public",
						Usage: "Mark pack public",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, BuildCreate)
				},
			},
			{
				Name:      "version-list",
				Usage:     "List assigned versions",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "ID or slug of the related pack",
					},
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Build ID or slug to list versions",
					},
					cli.StringFlag{
						Name:  "format",
						Value: tmplBuildVersionList,
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
					return Handle(c, BuildVersionList)
				},
			},
			{
				Name:      "version-append",
				Usage:     "Append a version to build",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "ID or slug of the related pack",
					},
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Build ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "mod, m",
						Value: "",
						Usage: "Mod ID or slug to append",
					},
					cli.StringFlag{
						Name:  "version, V",
						Value: "",
						Usage: "Version ID or slug to append",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, BuildVersionAppend)
				},
			},
			{
				Name:      "version-remove",
				Usage:     "Remove a version from build",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "ID or slug of the related pack",
					},
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Build ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "mod, m",
						Value: "",
						Usage: "Mod ID or slug to remove",
					},
					cli.StringFlag{
						Name:  "version, V",
						Value: "",
						Usage: "Version ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, BuildVersionRemove)
				},
			},
		},
	}
}

// BuildList provides the sub-command to list all builds.
func BuildList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.BuildList(
		GetPackParam(c),
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
		buildFuncMap,
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

// BuildShow provides the sub-command to show build details.
func BuildShow(c *cli.Context, client kleister.ClientAPI) error {
	record, err := client.BuildGet(
		GetPackParam(c),
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
		buildFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, record)
}

// BuildDelete provides the sub-command to delete a build.
func BuildDelete(c *cli.Context, client kleister.ClientAPI) error {
	err := client.BuildDelete(
		GetPackParam(c),
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// BuildUpdate provides the sub-command to update a build.
func BuildUpdate(c *cli.Context, client kleister.ClientAPI) error {
	record, err := client.BuildGet(
		GetPackParam(c),
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	changed := false

	if c.IsSet("minecraft") {
		if match, _ := regexp.MatchString("^([0-9]+)$", c.String("minecraft")); match {
			if val, err := strconv.ParseInt(c.String("minecraft"), 10, 64); err == nil && val != record.MinecraftID {
				record.MinecraftID = val
				changed = true
			}
		} else {
			if c.String("minecraft") != "" {
				related, err := client.MinecraftGet(
					c.String("minecraft"),
				)

				if err != nil {
					return err
				}

				if related.ID != record.MinecraftID {
					record.MinecraftID = related.ID
					changed = true
				}
			}
		}
	}

	if c.IsSet("forge") {
		if match, _ := regexp.MatchString("^([0-9]+)$", c.String("forge")); match {
			if val, err := strconv.ParseInt(c.String("forge"), 10, 64); err == nil && val != record.ForgeID {
				record.ForgeID = val
				changed = true
			}
		} else {
			if c.String("forge") != "" {
				related, err := client.ForgeGet(
					c.String("forge"),
				)

				if err != nil {
					return err
				}

				if related.ID != record.ForgeID {
					record.ForgeID = related.ID
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

	if val := c.String("min-java"); c.IsSet("min-java") && val != record.MinJava {
		record.MinJava = val
		changed = true
	}

	if val := c.String("min-memory"); c.IsSet("min-memory") && val != record.MinMemory {
		record.MinMemory = val
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
		_, patch := client.BuildPatch(
			GetPackParam(c),
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

// BuildCreate provides the sub-command to create a build.
func BuildCreate(c *cli.Context, client kleister.ClientAPI) error {
	record := &kleister.Build{}

	if c.String("pack") == "" {
		return fmt.Errorf("You must provide a pack ID or slug.")
	}

	if c.IsSet("pack") {
		if match, _ := regexp.MatchString("^([0-9]+)$", c.String("pack")); match {
			if val, err := strconv.ParseInt(c.String("pack"), 10, 64); err == nil && val != 0 {
				record.PackID = val
			}
		} else {
			if c.String("pack") != "" {
				related, err := client.PackGet(
					c.String("pack"),
				)

				if err != nil {
					return err
				}

				if related.ID != record.PackID {
					record.PackID = related.ID
				}
			}
		}
	}

	if c.IsSet("minecraft") {
		if match, _ := regexp.MatchString("^([0-9]+)$", c.String("minecraft")); match {
			if val, err := strconv.ParseInt(c.String("minecraft"), 10, 64); err == nil && val != 0 {
				record.MinecraftID = val
			}
		} else {
			if c.String("minecraft") != "" {
				related, err := client.MinecraftGet(
					c.String("minecraft"),
				)

				if err != nil {
					return err
				}

				if related.ID != record.MinecraftID {
					record.MinecraftID = related.ID
				}
			}
		}
	}

	if c.IsSet("forge") {
		if match, _ := regexp.MatchString("^([0-9]+)$", c.String("forge")); match {
			if val, err := strconv.ParseInt(c.String("forge"), 10, 64); err == nil && val != 0 {
				record.ForgeID = val
			}
		} else {
			if c.String("forge") != "" {
				related, err := client.ForgeGet(
					c.String("forge"),
				)

				if err != nil {
					return err
				}

				if related.ID != record.ForgeID {
					record.ForgeID = related.ID
				}
			}
		}
	}

	if val := c.String("name"); c.IsSet("name") && val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("You must provide a name.")
	}

	if val := c.String("slug"); c.IsSet("slug") && val != "" {
		record.Slug = val
	}

	if val := c.String("min-java"); c.IsSet("min-java") && val != "" {
		record.MinJava = val
	}

	if val := c.String("min-memory"); c.IsSet("min-memory") && val != "" {
		record.MinMemory = val
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

	_, err := client.BuildPost(
		GetPackParam(c),
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// BuildVersionList provides the sub-command to list versions of the build.
func BuildVersionList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.BuildVersionList(
		kleister.BuildVersionParams{
			Pack:  GetPackParam(c),
			Build: GetIdentifierParam(c),
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
		buildFuncMap,
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

// BuildVersionAppend provides the sub-command to append a version to the build.
func BuildVersionAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.BuildVersionAppend(
		kleister.BuildVersionParams{
			Pack:    GetPackParam(c),
			Build:   GetIdentifierParam(c),
			Mod:     GetModParam(c),
			Version: GetVersionParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to build\n")
	return nil
}

// BuildVersionRemove provides the sub-command to remove a version from the build.
func BuildVersionRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.BuildVersionDelete(
		kleister.BuildVersionParams{
			Pack:    GetPackParam(c),
			Build:   GetIdentifierParam(c),
			Mod:     GetModParam(c),
			Version: GetVersionParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from build\n")
	return nil
}
