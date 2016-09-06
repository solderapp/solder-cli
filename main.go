package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/kleister/kleister-cli/cmd"
	"github.com/kleister/kleister-cli/config"
	"github.com/urfave/cli"

	_ "github.com/joho/godotenv/autoload"
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
			Value:  "http://localhost:8080",
			Usage:  "Kleister API server",
			EnvVar: "KLEISTER_SERVER",
		},
		cli.StringFlag{
			Name:   "token, t",
			Value:  "",
			Usage:  "Kleister API token",
			EnvVar: "KLEISTER_TOKEN",
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

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
