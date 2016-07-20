package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/kleister/kleister-go/kleister"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// Version provides the sub-command for the version API.
func Version() cli.Command {
	return cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Version related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all versions",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "mod, m",
						Value: "",
						Usage: "ID or slug of the related mod",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a version",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "mod, m",
						Value: "",
						Usage: "ID or slug of the related mod",
					},
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionShow)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a version",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "mod, m",
						Value: "",
						Usage: "ID or slug of the related mod",
					},
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Version ID or slug to show",
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
						Name:  "file-url",
						Value: "",
						Usage: "Provide a file URL",
					},
					cli.StringFlag{
						Name:  "file-path",
						Value: "",
						Usage: "Provide a file path",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionUpdate)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a version",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "mod, m",
						Value: "",
						Usage: "ID or slug of the related mod",
					},
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionDelete)
				},
			},
			{
				Name:      "create",
				Usage:     "Create a version",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "mod, m",
						Value: "",
						Usage: "ID or slug of the related mod",
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
						Name:  "file-url",
						Value: "",
						Usage: "Provide a file URL",
					},
					cli.StringFlag{
						Name:  "file-path",
						Value: "",
						Usage: "Provide a file path",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionCreate)
				},
			},
			{
				Name:      "build-list",
				Usage:     "List assigned builds",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "mod, m",
						Value: "",
						Usage: "ID or slug of the related mod",
					},
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Version ID or slug to list builds",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionBuildList)
				},
			},
			{
				Name:      "build-append",
				Usage:     "Append a build to version",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "mod, m",
						Value: "",
						Usage: "ID or slug of the related mod",
					},
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Version ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "Pack ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "build, b",
						Value: "",
						Usage: "Build ID or slug to append to",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionBuildAppend)
				},
			},
			{
				Name:      "build-remove",
				Usage:     "Remove a build from version",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "mod, m",
						Value: "",
						Usage: "ID or slug of the related mod",
					},
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Version ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "Pack ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "build, b",
						Value: "",
						Usage: "Build ID or slug to append to",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionBuildRemove)
				},
			},
		},
	}
}

// VersionList provides the sub-command to list all versions.
func VersionList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.VersionList(
		GetModParam(c),
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

// VersionShow provides the sub-command to show version details.
func VersionShow(c *cli.Context, client kleister.ClientAPI) error {
	record, err := client.VersionGet(
		GetModParam(c),
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

	if record.Mod != nil {
		table.Append(
			[]string{
				"Mod",
				record.Mod.String(),
			},
		)
	}

	if record.File != nil {
		table.Append(
			[]string{
				"File",
				record.File.String(),
			},
		)
	}

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

// VersionDelete provides the sub-command to delete a version.
func VersionDelete(c *cli.Context, client kleister.ClientAPI) error {
	err := client.VersionDelete(
		GetModParam(c),
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// VersionUpdate provides the sub-command to update a version.
func VersionUpdate(c *cli.Context, client kleister.ClientAPI) error {
	record, err := client.VersionGet(
		GetModParam(c),
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

	if val := c.String("file-url"); c.IsSet("file-url") && val != "" {
		err := record.DownloadFile(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to download and encode file.")
		}

		changed = true
	}

	if val := c.String("file-path"); c.IsSet("file-path") && val != "" {
		err := record.EncodeFile(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to encode file.")
		}

		changed = true
	}

	if changed {
		_, patch := client.VersionPatch(
			GetModParam(c),
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

// VersionCreate provides the sub-command to create a version.
func VersionCreate(c *cli.Context, client kleister.ClientAPI) error {
	record := &kleister.Version{}

	if c.String("mod") == "" {
		return fmt.Errorf("You must provide a mod ID or slug.")
	}

	if c.IsSet("mod") {
		if match, _ := regexp.MatchString("^([0-9]+)$", c.String("mod")); match {
			if val, err := strconv.ParseInt(c.String("mod"), 10, 64); err == nil && val != 0 {
				record.ModID = val
			}
		} else {
			if c.String("mod") != "" {
				related, err := client.ModGet(
					c.String("mod"),
				)

				if err != nil {
					return err
				}

				if related.ID != record.ModID {
					record.ModID = related.ID
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

	if val := c.String("file-url"); c.IsSet("file-url") && val != "" {
		err := record.DownloadFile(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to download and encode file.")
		}
	}

	if val := c.String("file-path"); c.IsSet("file-path") && val != "" {
		err := record.EncodeFile(
			val,
		)

		if err != nil {
			return fmt.Errorf("Failed to encode file.")
		}
	}

	_, err := client.VersionPost(
		GetModParam(c),
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// VersionBuildList provides the sub-command to list builds of the version.
func VersionBuildList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.VersionBuildList(
		kleister.VersionBuildParams{
			Mod:     GetModParam(c),
			Version: GetIdentifierParam(c),
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
	table.SetHeader([]string{"Pack", "Build"})

	for _, record := range records {
		pack := "n/a"

		if record.Pack != nil {
			pack = record.Pack.Slug
		}

		table.Append(
			[]string{
				pack,
				record.Slug,
			},
		)
	}

	table.Render()
	return nil
}

// VersionBuildAppend provides the sub-command to append a build to the version.
func VersionBuildAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.VersionBuildAppend(
		kleister.VersionBuildParams{
			Mod:     GetModParam(c),
			Version: GetIdentifierParam(c),
			Pack:    GetPackParam(c),
			Build:   GetBuildParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to version\n")
	return nil
}

// VersionBuildRemove provides the sub-command to remove a build from the version.
func VersionBuildRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.VersionBuildDelete(
		kleister.VersionBuildParams{
			Mod:     GetModParam(c),
			Version: GetIdentifierParam(c),
			Pack:    GetPackParam(c),
			Build:   GetBuildParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from version\n")
	return nil
}
