package main

import (
	"github.com/urfave/cli/v2"
)

var RestoreAllCommand = &cli.Command{
	Name:  "restore-all",
	Usage: "Removes all the current Anyver app aliases. This should make the system's default executables the active ones for all apps",
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

		anyverYaml, err := ReadAnyverYaml(paths.ConfigFile)
		if err != nil {
			return err
		}

		for appName := range anyverYaml.Apps {
			if err := DeleteAppAlias(appName); err != nil {
				return err
			}

			if err := anyverYaml.UseVersion(appName, SystemVersion); err != nil {
				return err
			}

			write("Restored app %q", appName)
		}

		if err := SaveAnyverYaml(anyverYaml, paths.ConfigFile); err != nil {
			return err
		}

		return nil
	},
}
