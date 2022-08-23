package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var UseCommand = &cli.Command{
	Name:  "use",
	Usage: "Set the given version as the active command when running the app",
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

		args := c.Args()
		if c.NArg() != 2 {
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

		anyverYaml, err := ReadAnyverYaml(yamlFilePath)
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

			if err := SaveAnyverYaml(anyverYaml, yamlFilePath); err != nil {
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

		if err := SaveAnyverYaml(anyverYaml, yamlFilePath); err != nil {
			return err
		}

		write("Now using version %q for app %q", versionName, appName)
		return nil
	},
}
