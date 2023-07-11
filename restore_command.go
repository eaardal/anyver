package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var RestoreCommand = &cli.Command{
	Name:  "restore",
	Usage: "Removes the current Anyver app alias. This should make the system's default executable for the given app the active one.",
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
	Action: func(c *cli.Context) error {
		paths := GetAnyverPaths(c)

		args := c.Args()
		if c.NArg() < 1 {
			return fmt.Errorf("missing args: see usage")
		}

		appName := args.Get(0)
		if appName == "" {
			return fmt.Errorf("missing arg: app name")
		}

		anyverYaml, err := ReadAnyverYaml(paths.ConfigFile)
		if err != nil {
			return err
		}

		if err := DeleteAppAlias(appName); err != nil {
			return err
		}

		if err := anyverYaml.UseVersion(appName, SystemVersion); err != nil {
			return err
		}

		if err := SaveAnyverYaml(anyverYaml, paths.ConfigFile); err != nil {
			return err
		}

		write("Restored app %q", appName)
		return nil
	},
}
