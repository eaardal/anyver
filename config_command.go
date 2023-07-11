package main

import (
	"github.com/urfave/cli/v2"
	"os"
)

var ConfigCommand = &cli.Command{
	Name:  "config",
	Usage: "Print the Anyver YAML config file as-is",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			EnvVars: []string{"ANYVER_CONFIG"},
		},
		&cli.StringFlag{
			Name:    "apps-dir",
			Aliases: []string{"a"},
			EnvVars: []string{"ANYVER_APPS_DIR"},
		},
	},
	Aliases: []string{"yaml"},
	Action: func(c *cli.Context) error {
		paths := GetAnyverPaths(c)

		file, err := os.ReadFile(paths.ConfigFile)
		if err != nil {
			return err
		}

		write(paths.ConfigFile)
		write(string(file))

		return nil
	},
}
