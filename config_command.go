package main

import (
	"github.com/urfave/cli/v2"
	"os"
)

var ConfigCommand = &cli.Command{
	Name:  "config",
	Usage: "Print the entire Anyver config file as-is",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			EnvVars: []string{"ANYVER_CONFIG"},
		},
	},
	Action: func(c *cli.Context) error {
		yamlFilePath := c.String("config")
		if yamlFilePath == "" {
			yamlFilePath = DefaultAnyverYamlPath
		}

		file, err := os.ReadFile(yamlFilePath)
		if err != nil {
			return err
		}

		write(yamlFilePath)
		write(string(file))

		appFiles, err := os.ReadDir(AnyverAppsDirPath)
		if err != nil {
			return err
		}

		writeEmptyLine()
		write("Active app versions in %s:", AnyverAppsDirPath)

		if len(appFiles) == 0 {
			write("No active app versions")
		} else {
			for _, appFile := range appFiles {
				write(appFile.Name())
			}
		}

		return nil
	},
}
