package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/solderapp/solder-cli/solder"
)

// Profile provides the sub-command for the profile API.
func Profile() cli.Command {
	return cli.Command{
		Name:  "profile",
		Usage: "Profile related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:  "show",
				Usage: "Show profile details",
				Action: func(c *cli.Context) {
					Handle(c, ProfileShow)
				},
			},
			{
				Name:  "update",
				Usage: "Update your profile",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Provide a slug",
					},
					cli.StringFlag{
						Name:  "username",
						Value: "",
						Usage: "Provide a username",
					},
					cli.StringFlag{
						Name:  "email",
						Value: "",
						Usage: "Provide a email",
					},
					cli.StringFlag{
						Name:  "password",
						Value: "",
						Usage: "Provide a password",
					},
				},
				Action: func(c *cli.Context) {
					Handle(c, ProfileUpdate)
				},
			},
		},
	}
}

// ProfileShow provides the sub-command to show profile details.
func ProfileShow(c *cli.Context, client solder.API) error {
	record, err := client.ProfileGet()

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Key", "Value"})

	table.AppendBulk(
		[][]string{
			[]string{"ID", strconv.FormatInt(record.ID, 10)},
			[]string{"Slug", record.Slug},
			[]string{"Username", record.Username},
			[]string{"Email", record.Email},
			[]string{"Created", record.CreatedAt.Format(time.UnixDate)},
			[]string{"Updated", record.UpdatedAt.Format(time.UnixDate)},
		},
	)

	table.Render()
	return nil
}

// ProfileUpdate provides the sub-command to update the profile.
func ProfileUpdate(c *cli.Context, client solder.API) error {
	record, err := client.ProfileGet()

	if err != nil {
		return err
	}

	if val := c.String("slug"); val != record.Slug {
		record.Slug = val
	}

	if val := c.String("username"); val != record.Username {
		record.Username = val
	}

	if val := c.String("email"); val != record.Email {
		record.Email = val
	}

	if val := c.String("password"); val != "" {
		record.Password = val
	}

	_, patch := client.ProfilePatch(record)

	if patch != nil {
		return patch
	}

	fmt.Fprintf(os.Stderr, "Successfully updated\n")
	return nil
}
