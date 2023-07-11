package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var InitCommand = &cli.Command{
	Name:  "init",
	Usage: "Initialize Anyver. Creates a default Anyver YAML config file if it doesn't exist",
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

		if err := CreateDirIfNotExists(paths.RootDir); err != nil {
			return fmt.Errorf("failed to crate dir %s: %v", paths.RootDir, err)
		} else {
			write("Ensured %s exists", paths.RootDir)
		}

		if err := CreateDirIfNotExists(paths.AppsDir); err != nil {
			return fmt.Errorf("failed to create apps dir %s: %v", paths.AppsDir, err)
		} else {
			write("Ensured %s exists", paths.AppsDir)
		}

		if err := CreateFileIfNotExists(paths.ConfigFile); err != nil {
			return fmt.Errorf("failed to create config file %s: %v", paths.ConfigFile, err)
		} else {
			write("Ensured %s exists", paths.ConfigFile)
		}

		defaultConfig := &AnyverYaml{
			Active: map[string]string{
				"tryme": "demo",
			},
			Apps: map[string]map[string]string{
				"tryme": {
					"demo": "echo \"tryme demo\"",
				},
			},
			Contexts: map[string]map[string]string{
				"default": {
					"demo": SystemVersion,
				},
				"tryall": {
					"tryme": "demo",
				},
			},
		}

		if err := SaveAnyverYaml(defaultConfig, paths.ConfigFile); err != nil {
			return fmt.Errorf("failed to save %s: %v", paths.ConfigFile, err)
		}
		write("Created %s", paths.ConfigFile)

		return nil
	},
}
