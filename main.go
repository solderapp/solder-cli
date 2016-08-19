package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/kleister/kleister-cli/cmd"
	"github.com/kleister/kleister-cli/config"
	"github.com/sanbornm/go-selfupdate/selfupdate"
	"github.com/urfave/cli"
)

var (
	updates = "http://dl.webhippie.de/"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "kleister-cli"
	app.Version = config.Version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "Manage mod packs for Minecraft"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "server, s",
			Value:  "",
			Usage:  "Kleister API server",
			EnvVar: "KLEISTER_SERVER",
		},
		cli.StringFlag{
			Name:   "token, t",
			Value:  "",
			Usage:  "Kleister API token",
			EnvVar: "KLEISTER_TOKEN",
		},
		cli.BoolTFlag{
			Name:   "update, u",
			Usage:  "Enable auto update",
			EnvVar: "KLEISTER_UPDATE",
		},
	}

	app.Before = func(c *cli.Context) error {
		if c.BoolT("update") {
			if config.VersionDev == "dev" {
				fmt.Fprintf(os.Stderr, "Updates are disabled for development versions.\n")
			} else {
				updater := &selfupdate.Updater{
					CurrentVersion: fmt.Sprintf(
						"%d.%d.%d",
						config.VersionMajor,
						config.VersionMinor,
						config.VersionPatch,
					),
					ApiURL:  updates,
					BinURL:  updates,
					DiffURL: updates,
					Dir:     "updates/",
					CmdName: app.Name,
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
		cmd.Team(),
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
