package main

import (
	"github.com/urfave/cli/v2"
	"path/filepath"
)

func SetAnyverPaths(c *cli.Context) (anyverYamlPath string, anyverRootDir string) {
	yamlFilePath := c.String("config")
	if yamlFilePath == "" {
		yamlFilePath = DefaultAnyverYamlPath
	}

	yamlFileDir := filepath.Dir(yamlFilePath)
	AnyverAppsDirPath = filepath.Join(yamlFileDir, "apps")

	return yamlFilePath, yamlFileDir
}
