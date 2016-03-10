package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/solderapp/solder-cli/solder"
)

// HandleFunc is the real handle implementation.
type HandleFunc func(*cli.Context, solder.Client) error

// Handle wraps the command function handler.
func Handle(c *cli.Context, fn handleFunc) {
	token := c.GlobalString("token")
	server := c.GlobalString("server")

	if server == "" {
		fmt.Println("Error: You must provide the server address.")
		os.Exit(1)
	}

	if token == "" {
		fmt.Println("Error: You must provide your access token.")
		os.Exit(1)
	}

	client := solder.NewClientToken(
		server,
		token,
	)

	if err := fn(c, client); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
