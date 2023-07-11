package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var UseCommand = &cli.Command{
	Name:  "use",
	Usage: "Set the script for the given app name as the active app alias",
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
		if c.NArg() < 2 {
			return fmt.Errorf("missing args: see usage")
		}

		appName := args.Get(0)
		if appName == "" {
			return fmt.Errorf("missing arg: app name")
		}

		versionName := args.Get(1)
		if versionName == "" {
			return fmt.Errorf("missing arg: version name")
		}

		anyverYaml, err := ReadAnyverYaml(paths.ConfigFile)
		if err != nil {
			return err
		}

		if versionName == SystemVersion {
			if err := DeleteAppAlias(appName); err != nil {
				return err
			}

			if err := anyverYaml.UseVersion(appName, versionName); err != nil {
				return err
			}

			if err := SaveAnyverYaml(anyverYaml, paths.ConfigFile); err != nil {
				return err
			}

			write("Now using version %q for app %q", versionName, appName)
			return nil
		}

		versionScript, err := anyverYaml.FindVersionScript(appName, versionName)
		if err != nil {
			return err
		}

		if err := SetAppAlias(appName, versionScript); err != nil {
			return err
		}

		if err := anyverYaml.UseVersion(appName, versionName); err != nil {
			return err
		}

		if err := SaveAnyverYaml(anyverYaml, paths.ConfigFile); err != nil {
			return err
		}

		write("Now using version %q for app %q", versionName, appName)
		return nil
	},
}
