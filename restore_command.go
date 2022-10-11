package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var RestoreCommand = &cli.Command{
	Name:  "restore",
	Usage: "Removes the Anyver shortcut which should make the system's default executable the main one",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			EnvVars: []string{"ANYVER_CONFIG"},
		},
	},
	Action: func(c *cli.Context) error {
		yamlFilePath, _ := SetAnyverPaths(c)

		args := c.Args()
		if c.NArg() < 1 {
			return fmt.Errorf("missing args: see usage")
		}

		appName := args.Get(0)
		if appName == "" {
			return fmt.Errorf("missing arg: app name")
		}

		anyverYaml, err := ReadAnyverYaml(yamlFilePath)
		if err != nil {
			return err
		}

		if err := DeleteAppAlias(appName); err != nil {
			return err
		}

		if err := anyverYaml.UseVersion(appName, SystemVersion); err != nil {
			return err
		}

		if err := SaveAnyverYaml(anyverYaml, yamlFilePath); err != nil {
			return err
		}

		write("Restored app %q", appName)
		return nil
	},
}
