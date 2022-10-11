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
	},
	Aliases: []string{"yaml"},
	Action: func(c *cli.Context) error {
		yamlFilePath, _ := SetAnyverPaths(c)

		file, err := os.ReadFile(yamlFilePath)
		if err != nil {
			return err
		}

		write(yamlFilePath)
		write(string(file))

		return nil
	},
}
