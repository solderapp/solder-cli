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
				Action: func(c *cli.Context) {
					Handle(c, ProfileUpdate)
				},
			},
		},
	}
}

// ProfileShow provides the sub-command to show profile details.
func ProfileShow(c *cli.Context, client solder.API) error {
	return nil
}

// ProfileUpdate provides the sub-command to update the profile.
func ProfileUpdate(c *cli.Context, client solder.API) error {
	return nil
}
