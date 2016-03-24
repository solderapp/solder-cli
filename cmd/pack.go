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

// Pack provides the sub-command for the pack API.
func Pack() cli.Command {
	return cli.Command{
		Name:    "pack",
		Aliases: []string{"p"},
		Usage:   "Pack related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all packs",
				ArgsUsage: " ",
				Action: func(c *cli.Context) {
					Handle(c, PackList)
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
				},
				Action: func(c *cli.Context) {
					Handle(c, PackShow)
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
						Name:  "private",
						Usage: "Mark pack private",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, PackUpdate)
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
				Action: func(c *cli.Context) {
					Handle(c, PackDelete)
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
						Name:  "private",
						Usage: "Mark pack private",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, PackCreate)
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
				},
				Action: func(c *cli.Context) {
					Handle(c, PackClientList)
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
				Action: func(c *cli.Context) {
					Handle(c, PackClientAppend)
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
				Action: func(c *cli.Context) {
					Handle(c, PackClientRemove)
				},
			},
		},
	}
}

// PackList provides the sub-command to list all packs.
func PackList(c *cli.Context, client solder.ClientAPI) error {
	records, err := client.PackList()

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

// PackShow provides the sub-command to show pack details.
func PackShow(c *cli.Context, client solder.ClientAPI) error {
	record, err := client.PackGet(
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

	table.Append(
		[]string{
			"Website",
			record.Website,
		},
	)

	if record.Recommended != nil {
		table.Append(
			[]string{
				"Recommended",
				record.Recommended.String(),
			},
		)
	}

	if record.Latest != nil {
		table.Append(
			[]string{
				"Latest",
				record.Latest.String(),
			},
		)
	}

	if record.Icon != nil {
		table.Append(
			[]string{
				"Icon",
				record.Icon.String(),
			},
		)
	}

	if record.Logo != nil {
		table.Append(
			[]string{
				"Logo",
				record.Logo.String(),
			},
		)
	}

	if record.Background != nil {
		table.Append(
			[]string{
				"Background",
				record.Background.String(),
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

// PackDelete provides the sub-command to delete a pack.
func PackDelete(c *cli.Context, client solder.ClientAPI) error {
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
func PackUpdate(c *cli.Context, client solder.ClientAPI) error {
	record, err := client.PackGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	changed := false

	if c.IsSet("recommended") {
		if match, _ := regexp.MatchString("([0-9]+)", c.String("recommended")); match {
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
		if match, _ := regexp.MatchString("([0-9]+)", c.String("latest")); match {
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

	if val := c.Bool("published"); c.IsSet("published") && val != record.Published {
		record.Published = val
		changed = true
	}

	if val := c.Bool("private"); c.IsSet("private") && val != record.Private {
		record.Private = val
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
func PackCreate(c *cli.Context, client solder.ClientAPI) error {
	record := &solder.Pack{}

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

	if c.IsSet("published") {
		record.Published = c.Bool("published")
	}

	if c.IsSet("private") {
		record.Private = c.Bool("private")
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
func PackClientList(c *cli.Context, client solder.ClientAPI) error {
	records, err := client.PackClientList(
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

// PackClientAppend provides the sub-command to append a client to the pack.
func PackClientAppend(c *cli.Context, client solder.ClientAPI) error {
	err := client.PackClientAppend(
		GetIdentifierParam(c),
		GetClientParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to pack\n")
	return nil
}

// PackClientRemove provides the sub-command to remove a client from the pack.
func PackClientRemove(c *cli.Context, client solder.ClientAPI) error {
	err := client.PackClientDelete(
		GetIdentifierParam(c),
		GetClientParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from pack\n")
	return nil
}
