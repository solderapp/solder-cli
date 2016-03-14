package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

// GetIdentifierParam checks and returns the record id/slug parameter.
func GetIdentifierParam(c *cli.Context) string {
	val := c.String("id")

	if val == "" {
		fmt.Println("Error: You must provide an ID or a slug.")
		os.Exit(1)
	}

	return val
}

// GetModParam checks and returns the mod id/slug parameter.
func GetModParam(c *cli.Context) string {
	val := c.String("mod")

	if val == "" {
		fmt.Println("Error: You must provide a mod ID or slug.")
		os.Exit(1)
	}

	return val
}

// GetVersionParam checks and returns the version id/slug parameter.
func GetVersionParam(c *cli.Context) string {
	val := c.String("version")

	if val == "" {
		fmt.Println("Error: You must provide a version ID or slug.")
		os.Exit(1)
	}

	return val
}

// GetPackParam checks and returns the pack id/slug parameter.
func GetPackParam(c *cli.Context) string {
	val := c.String("pack")

	if val == "" {
		fmt.Println("Error: You must provide a pack ID or slug.")
		os.Exit(1)
	}

	return val
}

// GetBuildParam checks and returns the build id/slug parameter.
func GetBuildParam(c *cli.Context) string {
	val := c.String("build")

	if val == "" {
		fmt.Println("Error: You must provide a build ID or slug.")
		os.Exit(1)
	}

	return val
}

// GetClientParam checks and returns the client id/slug parameter.
func GetClientParam(c *cli.Context) string {
	val := c.String("client")

	if val == "" {
		fmt.Println("Error: You must provide a client ID or slug.")
		os.Exit(1)
	}

	return val
}
