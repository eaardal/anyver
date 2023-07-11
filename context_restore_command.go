package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var ContextRestoreCommand = &cli.Command{
	Name:  "restore",
	Usage: "Removes all Anyver app aliases defined in the given context, which should make the system's default executable the main one",
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

		contextName := args.Get(0)
		if contextName == "" {
			return fmt.Errorf("missing arg: context name")
		}

		anyverYaml, err := ReadAnyverYaml(paths.ConfigFile)
		if err != nil {
			return err
		}

		contextApps := anyverYaml.Contexts[contextName]
		if contextApps == nil || len(contextApps) == 0 {
			return fmt.Errorf("no apps found for context %s", contextName)
		}

		for appName := range contextApps {
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
