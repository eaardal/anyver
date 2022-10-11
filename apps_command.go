package main

import (
	"github.com/urfave/cli/v2"
	"os"
)

var AppsCommand = &cli.Command{
	Name:  "apps",
	Usage: "Print the currently active app aliases",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			EnvVars: []string{"ANYVER_CONFIG"},
		},
	},
	Action: func(c *cli.Context) error {
		appFiles, err := os.ReadDir(AnyverAppsDirPath)
		if err != nil {
			return err
		}

		writeEmptyLine()
		write("Active app aliases in %s:", AnyverAppsDirPath)

		if len(appFiles) == 0 {
			write("No active app aliases")
		} else {
			for _, appFile := range appFiles {
				write(appFile.Name())
			}
		}

		return nil
	},
}
