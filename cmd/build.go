package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
	"github.com/solderapp/solder-go/solder"
)

// Build provides the sub-command for the build API.
func Build() cli.Command {
	return cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build related sub-commands",
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
				},
				Action: func(c *cli.Context) {
					Handle(c, BuildList)
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
				},
				Action: func(c *cli.Context) {
					Handle(c, BuildShow)
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
				Action: func(c *cli.Context) {
					Handle(c, BuildUpdate)
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
				Action: func(c *cli.Context) {
					Handle(c, BuildDelete)
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
				Action: func(c *cli.Context) {
					Handle(c, BuildCreate)
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
				},
				Action: func(c *cli.Context) {
					Handle(c, BuildVersionList)
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
						Name:  "version, V",
						Value: "",
						Usage: "Version ID or slug to append",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, BuildVersionAppend)
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
						Name:  "version, V",
						Value: "",
						Usage: "Version ID or slug to remove",
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
func BuildList(c *cli.Context, client solder.ClientAPI) error {
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
func BuildShow(c *cli.Context, client solder.ClientAPI) error {
	record, err := client.BuildGet(
		GetPackParam(c),
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

	if record.Pack != nil {
		table.Append(
			[]string{
				"Pack",
				record.Pack.String(),
			},
		)
	}

	if record.Minecraft != nil {
		table.Append(
			[]string{
				"Minecraft",
				record.Minecraft.String(),
			},
		)
	}

	if record.Forge != nil {
		table.Append(
			[]string{
				"Forge",
				record.Forge.String(),
			},
		)
	}

	if record.MinJava != "" {
		table.Append(
			[]string{
				"Java",
				record.MinJava,
			},
		)
	}

	if record.MinMemory != "" {
		table.Append(
			[]string{
				"Memory",
				record.MinMemory,
			},
		)
	}

	table.Append(
		[]string{
			"Published",
			strconv.FormatBool(record.Published),
		},
	)

	table.Append(
		[]string{
			"Private",
			strconv.FormatBool(record.Private),
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

// BuildDelete provides the sub-command to delete a build.
func BuildDelete(c *cli.Context, client solder.ClientAPI) error {
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
func BuildUpdate(c *cli.Context, client solder.ClientAPI) error {
	record, err := client.BuildGet(
		GetPackParam(c),
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	changed := false

	if c.IsSet("minecraft") {
		if match, _ := regexp.MatchString("([0-9]+)", c.String("minecraft")); match {
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
		if match, _ := regexp.MatchString("([0-9]+)", c.String("forge")); match {
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
func BuildCreate(c *cli.Context, client solder.ClientAPI) error {
	record := &solder.Build{}

	if c.String("pack") == "" {
		return fmt.Errorf("You must provide a pack ID or slug.")
	}

	if c.IsSet("pack") {
		if match, _ := regexp.MatchString("([0-9]+)", c.String("pack")); match {
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
		if match, _ := regexp.MatchString("([0-9]+)", c.String("minecraft")); match {
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
		if match, _ := regexp.MatchString("([0-9]+)", c.String("forge")); match {
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
func BuildVersionList(c *cli.Context, client solder.ClientAPI) error {
	records, err := client.BuildVersionList(
		GetPackParam(c),
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

// BuildVersionAppend provides the sub-command to append a version to the build.
func BuildVersionAppend(c *cli.Context, client solder.ClientAPI) error {
	err := client.BuildVersionAppend(
		GetPackParam(c),
		GetIdentifierParam(c),
		GetVersionParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to build\n")
	return nil
}

// BuildVersionRemove provides the sub-command to remove a version from the build.
func BuildVersionRemove(c *cli.Context, client solder.ClientAPI) error {
	err := client.BuildVersionDelete(
		GetPackParam(c),
		GetIdentifierParam(c),
		GetVersionParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from build\n")
	return nil
}
