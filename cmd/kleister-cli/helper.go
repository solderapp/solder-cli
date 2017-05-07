package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/kleister/kleister-go/kleister"
	"github.com/urfave/cli"
)

// globalFuncMap provides global template helper functions.
var globalFuncMap = template.FuncMap{
	"split":    strings.Split,
	"join":     strings.Join,
	"toUpper":  strings.ToUpper,
	"toLower":  strings.ToLower,
	"contains": strings.Contains,
	"replace":  strings.Replace,
	"buildList": func(s []*kleister.Build) string {
		res := []string{}

		for _, row := range s {
			if row.Pack != nil {
				res = append(res, fmt.Sprintf("%s@%s", row.Pack.Slug, row.String()))
			} else {
				res = append(res, row.String())
			}
		}

		return strings.Join(res, ", ")
	},
	"clientList": func(s []*kleister.Client) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"modList": func(s []*kleister.Mod) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"packList": func(s []*kleister.Pack) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"teamList": func(s []*kleister.Team) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"userList": func(s []*kleister.User) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"versionList": func(s []*kleister.Version) string {
		res := []string{}

		for _, row := range s {
			if row.Mod != nil {
				res = append(res, fmt.Sprintf("%s@%s", row.Mod.Slug, row.String()))
			} else {
				res = append(res, row.String())
			}
		}

		return strings.Join(res, ", ")
	},
}

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

// GetUserParam checks and returns the user id/slug parameter.
func GetUserParam(c *cli.Context) string {
	val := c.String("user")

	if val == "" {
		fmt.Println("Error: You must provide a user ID or slug.")
		os.Exit(1)
	}

	return val
}

// GetTeamParam checks and returns the team id/slug parameter.
func GetTeamParam(c *cli.Context) string {
	val := c.String("team")

	if val == "" {
		fmt.Println("Error: You must provide a team ID or slug.")
		os.Exit(1)
	}

	return val
}

// GetPermParam checks and returns the permission parameter.
func GetPermParam(c *cli.Context) string {
	val := c.String("perm")

	if val == "" {
		fmt.Println("Error: You must provide a permission.")
		os.Exit(1)
	}

	for _, perm := range []string{"user", "admin", "owner"} {
		if perm == val {
			return val
		}
	}

	fmt.Println("Error: Invalid permission, can be user, admin or owner.")
	os.Exit(1)

	return ""
}
