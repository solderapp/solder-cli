package main

import (
	"os"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/solderapp/solder-cli/cmd"
)

var (
	version string
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "solder-cli"
	app.Version = version
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
