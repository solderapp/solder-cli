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

// minecraftFuncMap provides template helper functions.
var minecraftFuncMap = template.FuncMap{}

// tmplMinecraftList represents a row within forge listing.
var tmplMinecraftList = "Slug: \x1b[33m{{ .Slug }}\x1b[0m" + `
ID: {{ .ID }}
Version: {{ .Version }}
Type: {{ .Type }}
`

// tmplMinecraftBuildList represents a row within minecraft build listing.
var tmplMinecraftBuildList = "Slug: \x1b[33m{{ .Slug }}\x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}
Pack: {{with .Pack}}{{ . }}{{else}}n/a{{end}}
`

// Minecraft provides the sub-command for the Minecraft API.
func Minecraft() cli.Command {
	return cli.Command{
		Name:  "minecraft",
		Usage: "Minecraft related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all Minecraft versions",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "format",
						Value: tmplMinecraftList,
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
					return Handle(c, MinecraftList)
				},
			},
			{
				Name:      "refresh",
				Aliases:   []string{"ref"},
				Usage:     "Refresh Minecraft versions",
				ArgsUsage: " ",
				Action: func(c *cli.Context) error {
					return Handle(c, MinecraftRefresh)
				},
			},
			{
				Name:  "build",
				Usage: "Build assignments",
				Subcommands: []cli.Command{
					{
						Name:      "list",
						Aliases:   []string{"ls"},
						Usage:     "List assigned builds",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Minecraft ID or slug to list builds",
							},
							cli.StringFlag{
								Name:  "format",
								Value: tmplMinecraftBuildList,
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
							return Handle(c, MinecraftBuildList)
						},
					},
					{
						Name:      "append",
						Usage:     "Append a build to Minecraft",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Minecraft ID or slug to append to",
							},
							cli.StringFlag{
								Name:  "pack, p",
								Value: "",
								Usage: "Pack ID or slug to append",
							},
							cli.StringFlag{
								Name:  "build, b",
								Value: "",
								Usage: "Build ID or slug to append",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, MinecraftBuildAppend)
						},
					},
					{
						Name:      "remove",
						Aliases:   []string{"rm"},
						Usage:     "Remove a build from Minecraft",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Minecraft ID or slug to remove from",
							},
							cli.StringFlag{
								Name:  "pack, p",
								Value: "",
								Usage: "Pack ID or slug to remove",
							},
							cli.StringFlag{
								Name:  "build, b",
								Value: "",
								Usage: "Build ID or slug to remove",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, MinecraftBuildRemove)
						},
					},
				},
			},
		},
	}
}

// MinecraftList provides the sub-command to list all Minecraft versions.
func MinecraftList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.MinecraftList()

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
		minecraftFuncMap,
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

// MinecraftRefresh provides the sub-command to refresh the Minecraft versions.
func MinecraftRefresh(c *cli.Context, client kleister.ClientAPI) error {
	err := client.MinecraftRefresh()

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully refreshed\n")
	return nil
}

// MinecraftBuildList provides the sub-command to list builds of the Minecraft.
func MinecraftBuildList(c *cli.Context, client kleister.ClientAPI) error {
	records, err := client.MinecraftBuildList(
		kleister.MinecraftBuildParams{
			Minecraft: GetIdentifierParam(c),
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
		minecraftFuncMap,
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

// MinecraftBuildAppend provides the sub-command to append a build to the Minecraft.
func MinecraftBuildAppend(c *cli.Context, client kleister.ClientAPI) error {
	err := client.MinecraftBuildAppend(
		kleister.MinecraftBuildParams{
			Minecraft: GetIdentifierParam(c),
			Pack:      GetPackParam(c),
			Build:     GetBuildParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to Minecraft\n")
	return nil
}

// MinecraftBuildRemove provides the sub-command to remove a build from the Minecraft.
func MinecraftBuildRemove(c *cli.Context, client kleister.ClientAPI) error {
	err := client.MinecraftBuildDelete(
		kleister.MinecraftBuildParams{
			Minecraft: GetIdentifierParam(c),
			Pack:      GetPackParam(c),
			Build:     GetBuildParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from Minecraft\n")
	return nil
}
