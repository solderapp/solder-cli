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

// Build provides the sub-command for the build API.
func Build() cli.Command {
	return cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all builds",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "pack, p",
						Value: "",
						Usage: "ID or slug of the related pack",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, BuildList)
				},
			},
			{
				Name:  "show",
				Usage: "Display a build",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Build ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, BuildShow)
				},
			},
			{
				Name:  "update",
				Usage: "Update a build",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
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
						Usage: "Provide a Minecraft ID",
					},
					cli.StringFlag{
						Name:  "forge",
						Value: "",
						Usage: "Provide a Forge ID",
					},
					cli.BoolFlag{
						Name:  "published",
						Usage: "Mark build published",
					},
					cli.BoolFlag{
						Name:  "private",
						Usage: "Mark build private",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, BuildUpdate)
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"rm"},
				Usage:   "Delete a build",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Build ID or slug to delete",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, BuildDelete)
				},
			},
			{
				Name:  "create",
				Usage: "Create a build",
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
						Usage: "Provide a Minecraft ID",
					},
					cli.StringFlag{
						Name:  "forge",
						Value: "",
						Usage: "Provide a Forge ID",
					},
					cli.BoolFlag{
						Name:  "published",
						Usage: "Mark build published",
					},
					cli.BoolFlag{
						Name:  "private",
						Usage: "Mark build private",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, BuildCreate)
				},
			},
			{
				Name:  "version-list",
				Usage: "List assigned versions",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Build ID or slug to list versions",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, BuildVersionList)
				},
			},
			{
				Name:  "version-append",
				Usage: "Append a version to build",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Build ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "version",
						Value: "",
						Usage: "Version ID or slug to append to",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, BuildVersionAppend)
				},
			},
			{
				Name:  "version-remove",
				Usage: "Remove a version from build",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Build ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "version",
						Value: "",
						Usage: "Version ID or slug to append to",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, BuildVersionRemove)
				},
			},
		},
	}
}

// BuildList provides the sub-command to list all builds.
func BuildList(c *cli.Context, client solder.API) error {
	records, err := client.BuildList(
		GetPackParam(c),
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

// BuildShow provides the sub-command to show build details.
func BuildShow(c *cli.Context, client solder.API) error {
	record, err := client.BuildGet(
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
			{"Slug", record.Slug},
			{"Name", record.Name},
			{"Pack", record.Pack},
			{"Minecraft", record.Minecraft},
			{"Forge", record.Forge},
			{"Java", record.MinJava},
			{"Memory", record.MinMemory},
			{"Published", strconv.FormatBool(record.Published)},
			{"Private", strconv.FormatBool(record.Private)},
			{"Created", record.CreatedAt.Format(time.UnixDate)},
			{"Updated", record.UpdatedAt.Format(time.UnixDate)},
		},
	)

	table.Render()
	return nil
}

// BuildDelete provides the sub-command to delete a build.
func BuildDelete(c *cli.Context, client solder.API) error {
	err := client.BuildDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// BuildUpdate provides the sub-command to update a build.
func BuildUpdate(c *cli.Context, client solder.API) error {
	record, err := client.BuildGet(
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

	if val, err := strconv.ParseInt(c.String("minecraft"), 10, 64); err == nil && val != record.MinecraftID {
		record.MinecraftID = val
	}

	if val, err := strconv.ParseInt(c.String("forge"), 10, 64); err == nil && val != record.ForgeID {
		record.ForgeID = val
	}

	if val := c.String("min-java"); val != record.MinJava {
		record.MinJava = val
	}

	if val := c.String("min-memory"); val != record.MinMemory {
		record.MinMemory = val
	}

	if val := c.Bool("published"); val != record.Published {
		record.Published = val
	}

	if val := c.Bool("private"); val != record.Private {
		record.Private = val
	}

	_, patch := client.BuildPatch(record)

	if patch != nil {
		return patch
	}

	fmt.Fprintf(os.Stderr, "Successfully updated\n")
	return nil
}

// BuildCreate provides the sub-command to create a build.
func BuildCreate(c *cli.Context, client solder.API) error {
	record := &solder.Build{}

	if val, err := strconv.ParseInt(c.String("pack"), 10, 64); err == nil && val != 0 {
		record.PackID = val
	} else {
		return fmt.Errorf("You must provide a pack.")
	}

	if val := c.String("name"); val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("You must provide a name.")
	}

	if val := c.String("slug"); val != "" {
		record.Slug = val
	}

	if val, err := strconv.ParseInt(c.String("minecraft"), 10, 64); err == nil && val != 0 {
		record.MinecraftID = val
	}

	if val, err := strconv.ParseInt(c.String("forge"), 10, 64); err == nil && val != 0 {
		record.ForgeID = val
	}

	if val := c.String("min-java"); val != "" {
		record.MinJava = val
	}

	if val := c.String("min-memory"); val != "" {
		record.MinMemory = val
	}

	if val := c.Bool("published"); val != false {
		record.Published = val
	}

	if val := c.Bool("private"); val != false {
		record.Private = val
	}

	_, err := client.BuildPost(record)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// BuildVersionList provides the sub-command to list versions of the build.
func BuildVersionList(c *cli.Context, client solder.API) error {
	return nil
}

// BuildVersionAppend provides the sub-command to append a version to the build.
func BuildVersionAppend(c *cli.Context, client solder.API) error {
	return nil
}

// BuildVersionRemove provides the sub-command to remove a version from the build.
func BuildVersionRemove(c *cli.Context, client solder.API) error {
	return nil
}
