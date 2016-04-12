package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/sanbornm/go-selfupdate/selfupdate"
	"github.com/solderapp/solder-cli/cmd"
	"github.com/solderapp/solder-cli/config"
)

var (
	updates string = "http://dl.webhippie.de/"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "solder-cli"
	app.Version = config.Version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "Manage mod packs for the Technic launcher"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "server, s",
			Value:  "",
			Usage:  "Solder API server",
			EnvVar: "SOLDER_SERVER",
		},
		cli.StringFlag{
			Name:   "token, t",
			Value:  "",
			Usage:  "Solder API token",
			EnvVar: "SOLDER_TOKEN",
		},
		cli.BoolFlag{
			Name:   "update, u",
			Usage:  "Enable auto update",
			EnvVar: "SOLDER_UPDATE",
		},
	}

	app.Before = func(c *cli.Context) error {
		if c.Bool("update") {
			if config.VersionDev == "dev" {
				fmt.Fprintf(os.Stderr, "Updates are disabled for development versions.\n")
			} else {
				updater := &selfupdate.Updater{
					CurrentVersion: config.StrippedVersion,
					ApiURL:         updates,
					BinURL:         updates,
					DiffURL:        updates,
					Dir:            "updates/",
					CmdName:        app.Name,
				}

				go updater.BackgroundRun()
			}
		}

		return nil
	}

	app.Commands = []cli.Command{
		cmd.Pack(),
		cmd.Build(),
		cmd.Mod(),
		cmd.Version(),
		cmd.Minecraft(),
		cmd.Forge(),
		cmd.User(),
		cmd.Key(),
		cmd.Client(),
		cmd.Profile(),
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Show the help, so what you see now",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Print the current version of that tool",
	}

	app.Run(os.Args)
}
