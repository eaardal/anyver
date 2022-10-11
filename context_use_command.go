package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var ContextUseCommand = &cli.Command{
	Name:  "use",
	Usage: "Set the script for all apps in the given context as the active app aliases",
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

		contextName := args.Get(0)
		if contextName == "" {
			return fmt.Errorf("missing arg: context name")
		}

		anyverYaml, err := ReadAnyverYaml(yamlFilePath)
		if err != nil {
			return err
		}

		versions, err := anyverYaml.FindContextScripts(contextName)
		if err != nil {
			return err
		}

		if len(versions) == 0 {
			return fmt.Errorf("found no versions for context %q", contextName)
		}

		for _, version := range versions {
			if version.Name == SystemVersion {
				if err := DeleteAppAlias(version.AppName); err != nil {
					return err
				}

				if err := anyverYaml.UseVersion(version.AppName, version.Name); err != nil {
					return err
				}

				continue
			}

			if err := SetAppAlias(version.AppName, version.Script); err != nil {
				return err
			}

			if err := anyverYaml.UseVersion(version.AppName, version.Name); err != nil {
				return err
			}

			write("Now using version %q for app %q", version.Name, version.AppName)
		}

		if err := SaveAnyverYaml(anyverYaml, yamlFilePath); err != nil {
			return err
		}

		return nil
	},
}
