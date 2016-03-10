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

// GetPackParam checks and returns the pack id/slug parameter.
func GetPackParam(c *cli.Context) string {
	val := c.String("pack")

	if val == "" {
		fmt.Println("Error: You must provide a pack ID or slug.")
		os.Exit(1)
	}

	return val
}
