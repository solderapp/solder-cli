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

// Pack provides the sub-command for the pack API.
func Pack() cli.Command {
	return cli.Command{
		Name:    "pack",
		Aliases: []string{"p"},
		Usage:   "Pack related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "List all packs",
				Action: func(c *cli.Context) {
					Handle(c, PackList)
				},
			},
			{
				Name:  "show",
				Usage: "Display a pack",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Pack ID or slug to show",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, PackShow)
				},
			},
			{
				Name:  "update",
				Usage: "Update a pack",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
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
						Usage: "Recommended build ID",
					},
					cli.StringFlag{
						Name:  "latest",
						Value: "",
						Usage: "Latest build ID",
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
						Name:  "hidden",
						Usage: "Mark pack hidden",
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
				Name:    "delete",
				Aliases: []string{"rm"},
				Usage:   "Delete a pack",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Pack ID or slug to delete",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, PackDelete)
				},
			},
			{
				Name:  "create",
				Usage: "Create a pack",
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
						Name:  "recommended",
						Value: "",
						Usage: "Recommended build ID",
					},
					cli.StringFlag{
						Name:  "latest",
						Value: "",
						Usage: "Latest build ID",
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
						Name:  "hidden",
						Usage: "Mark pack hidden",
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
				Name:  "client-list",
				Usage: "List assigned clients",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Pack ID or slug to list clients",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, PackClientList)
				},
			},
			{
				Name:  "client-append",
				Usage: "Append a client to pack",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Pack ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "client",
						Value: "",
						Usage: "Client ID or slug to append",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, PackClientAppend)
				},
			},
			{
				Name:  "client-remove",
				Usage: "Remove a client from pack",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id",
						Value: "",
						Usage: "Pack ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "client",
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
func PackList(c *cli.Context, client solder.API) error {
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
func PackShow(c *cli.Context, client solder.API) error {
	record, err := client.PackGet(
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
			{"Website", record.Website},
			{"Recommended", record.Recommended},
			{"Latest", record.Latest},
			{"Icon", record.Icon},
			{"Logo", record.Logo},
			{"Background", record.Background},
			{"Hidden", strconv.FormatBool(record.Hidden)},
			{"Private", strconv.FormatBool(record.Private)},
			{"Created", record.CreatedAt.Format(time.UnixDate)},
			{"Updated", record.UpdatedAt.Format(time.UnixDate)},
		},
	)

	table.Render()
	return nil
}

// PackDelete provides the sub-command to delete a pack.
func PackDelete(c *cli.Context, client solder.API) error {
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
func PackUpdate(c *cli.Context, client solder.API) error {
	record, err := client.PackGet(
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

	if val := c.String("website"); val != record.Website {
		record.Website = val
	}

	if val, err := strconv.ParseInt(c.String("recommended"), 10, 64); err == nil && val != record.RecommendedID {
		record.RecommendedID = val
	}

	if val, err := strconv.ParseInt(c.String("latest"), 10, 64); err == nil && val != record.LatestID {
		record.LatestID = val
	}

	// TODO(must): Implement URL import
	// if val := c.String("icon-url"); val != record.IconURL {
	// 	record.IconURL = val
	// }

	// TODO(must): Implement path import
	// if val := c.String("icon-path"); val != record.IconPath {
	// 	record.IconPath = val
	// }

	// TODO(must): Implement URL import
	// if val := c.String("logo-url"); val != record.LogoURL {
	// 	record.LogoURL = val
	// }

	// TODO(must): Implement path import
	// if val := c.String("logo-path"); val != record.LogoPath {
	// 	record.LogoPath = val
	// }

	// TODO(must): Implement URL import
	// if val := c.String("bg-url"); val != record.BackgrounURL {
	// 	record.BackgroundURL = val
	// }

	// TODO(must): Implement path import
	// if val := c.String("bg-path"); val != record.BackgroundPath {
	// 	record.BackgroundPath = val
	// }

	if val := c.Bool("hidden"); val != record.Hidden {
		record.Hidden = val
	}

	if val := c.Bool("private"); val != record.Private {
		record.Private = val
	}

	_, patch := client.PackPatch(record)

	if patch != nil {
		return patch
	}

	fmt.Fprintf(os.Stderr, "Successfully updated\n")
	return nil
}

// PackCreate provides the sub-command to create a pack.
func PackCreate(c *cli.Context, client solder.API) error {
	record := &solder.Pack{}

	if val := c.String("name"); val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("You must provide a name.")
	}

	if val := c.String("slug"); val != "" {
		record.Slug = val
	}

	if val := c.String("website"); val != "" {
		record.Website = val
	}

	if val, err := strconv.ParseInt(c.String("recommended"), 10, 64); err == nil && val != 0 {
		record.RecommendedID = val
	}

	if val, err := strconv.ParseInt(c.String("latest"), 10, 64); err == nil && val != 0 {
		record.LatestID = val
	}

	// TODO(must): Implement URL import
	// if val := c.String("icon-url"); val != "" {
	// 	record.IconURL = val
	// }

	// TODO(must): Implement path import
	// if val := c.String("icon-path"); val != "" {
	// 	record.IconPath = val
	// }

	// TODO(must): Implement URL import
	// if val := c.String("logo-url"); val != "" {
	// 	record.LogoURL = val
	// }

	// TODO(must): Implement path import
	// if val := c.String("logo-path"); val != "" {
	// 	record.LogoPath = val
	// }

	// TODO(must): Implement URL import
	// if val := c.String("bg-url"); val != "" {
	// 	record.BackgroundURL = val
	// }

	// TODO(must): Implement path import
	// if val := c.String("bg-path"); val != "" {
	// 	record.BackgroundPath = val
	// }

	if val := c.Bool("hidden"); val != false {
		record.Hidden = val
	}

	if val := c.Bool("private"); val != false {
		record.Private = val
	}

	_, err := client.PackPost(record)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// PackClientList provides the sub-command to list packs of the pack.
func PackClientList(c *cli.Context, client solder.API) error {
	return nil
}

// PackClientAppend provides the sub-command to append a client to the pack.
func PackClientAppend(c *cli.Context, client solder.API) error {
	return nil
}

// PackClientRemove provides the sub-command to remove a client from the pack.
func PackClientRemove(c *cli.Context, client solder.API) error {
	return nil
}
