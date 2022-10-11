package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"path/filepath"
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
	},
	Action: func(c *cli.Context) error {
		yamlFilePath, yamlFileDir := SetAnyverPaths(c)

		if err := CreateDirIfNotExists(yamlFileDir); err != nil {
			return fmt.Errorf("failed to crate dir %s: %v", yamlFileDir, err)
		} else {
			write("Ensured %s exists", yamlFileDir)
		}

		AnyverAppsDirPath = filepath.Join(yamlFileDir, "apps")

		if err := CreateDirIfNotExists(AnyverAppsDirPath); err != nil {
			return fmt.Errorf("failed to create apps dir %s: %v", AnyverAppsDirPath, err)
		} else {
			write("Ensured %s exists", AnyverAppsDirPath)
		}

		if err := CreateFileIfNotExists(yamlFilePath); err != nil {
			return fmt.Errorf("failed to create config file %s: %v", yamlFilePath, err)
		} else {
			write("Ensured %s exists", yamlFilePath)
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

		if err := SaveAnyverYaml(defaultConfig, yamlFilePath); err != nil {
			return fmt.Errorf("failed to save %s: %v", yamlFilePath, err)
		}
		write("Created %s", yamlFilePath)

		return nil
	},
}
