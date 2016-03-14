package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
	"github.com/solderapp/solder-cli/solder"
)

// Version provides the sub-command for the version API.
func Version() cli.Command {
	return cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Version related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all versions",
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
				Name:  "show",
				Usage: "Display a version",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionShow)
				},
			},
			{
				Name:  "update",
				Usage: "Update a version",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
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
				Name:    "delete",
				Aliases: []string{"rm"},
				Usage:   "Delete a version",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionDelete)
				},
			},
			{
				Name:  "create",
				Usage: "Create a version",
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
				Name:  "build-list",
				Usage: "List assigned builds",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to list builds",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionBuildList)
				},
			},
			{
				Name:  "build-append",
				Usage: "Append a mod version to build",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "build",
						Value: "",
						Usage: "Build ID or slug to append to",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, VersionBuildAppend)
				},
			},
			{
				Name:  "build-remove",
				Usage: "Remove a mod version from build",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Version ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "build",
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
func VersionList(c *cli.Context, client solder.API) error {
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
func VersionShow(c *cli.Context, client solder.API) error {
	record, err := client.VersionGet(
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
			{"ID", strconv.FormatInt(record.ID, 10)},
			{"Mod", record.Mod},
			{"Slug", record.Slug},
			{"Name", record.Name},
			{"File", record.File},
			{"Created", record.CreatedAt.Format(time.UnixDate)},
			{"Updated", record.UpdatedAt.Format(time.UnixDate)},
		},
	)

	table.Render()
	return nil
}

// VersionDelete provides the sub-command to delete a version.
func VersionDelete(c *cli.Context, client solder.API) error {
	err := client.VersionDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// VersionUpdate provides the sub-command to update a version.
func VersionUpdate(c *cli.Context, client solder.API) error {
	record, err := client.VersionGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	if val := c.String("name"); val != record.Name {
		record.Name = val
	}

	if val := c.String("slug"); val != record.Slug {
		record.Slug = val
	}

	// TODO(must): Implement URL import
	// if val := c.String("file-url"); val != record.FileURL {
	// 	record.FileURL = val
	// }

	// TODO(must): Implement path import
	// if val := c.String("file-path"); val != record.FilePath {
	// 	record.FilePath = val
	// }

	_, patch := client.VersionPatch(record)

	if patch != nil {
		return patch
	}

	fmt.Fprintf(os.Stderr, "Successfully updated\n")
	return nil
}

// VersionCreate provides the sub-command to create a version.
func VersionCreate(c *cli.Context, client solder.API) error {
	record := &solder.Version{}

	if c.String("mod") == "" {
		return fmt.Errorf("You must provide a mod.")
	}

	if match, _ := regexp.MatchString("([0-9]+)", c.String("mod")); match {
		if val, err := strconv.ParseInt(c.String("mod"), 10, 64); err == nil && val != 0 {
			record.ModID = val
		}
	} else {
		if c.String("mod") != "" {
			related, err := client.BuildGet(
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

	if val := c.String("name"); val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("You must provide a name.")
		os.Exit(1)
	}

	if val := c.String("slug"); val != "" {
		record.Slug = val
	}

	// TODO(must): Implement URL import
	// if val := c.String("file-url"); val != "" {
	// 	record.FileURL = val
	// }

	// TODO(must): Implement path import
	// if val := c.String("file-path"); val != "" {
	// 	record.FilePath = val
	// }

	_, err := client.VersionPost(record)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// VersionBuildList provides the sub-command to list builds of the version.
func VersionBuildList(c *cli.Context, client solder.API) error {
	return nil
}

// VersionBuildAppend provides the sub-command to append a build to the version.
func VersionBuildAppend(c *cli.Context, client solder.API) error {
	return nil
}

// VersionBuildRemove provides the sub-command to remove a build from the version.
func VersionBuildRemove(c *cli.Context, client solder.API) error {
	return nil
}
