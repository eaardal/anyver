package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"path"
	"path/filepath"
)

var defaultAnyverRoot = path.Join(os.Getenv("HOME"), ".anyver")
var defaultAnyverAppsDirPath = path.Join(defaultAnyverRoot, "apps")
var defaultAnyverConfigFilePath = path.Join(defaultAnyverRoot, "config.yaml")

type AnyverPaths struct {
	ConfigFile string
	AppsDir    string
	RootDir    string
}

func getAnyverPathsFromCLIContext(c *cli.Context) *AnyverPaths {
	paths := &AnyverPaths{}

	yamlFilePath := c.String("config")
	if yamlFilePath == "" {
		paths.ConfigFile = defaultAnyverConfigFilePath
	} else {
		paths.ConfigFile = yamlFilePath
	}

	paths.RootDir = filepath.Dir(paths.ConfigFile)

	appsDir := c.String("apps-dir")
	if appsDir == "" {
		paths.AppsDir = defaultAnyverAppsDirPath
	} else {
		paths.AppsDir = appsDir
	}

	return paths
}

func getAnyverPathsFromEnvsOrDefault() *AnyverPaths {
	paths := &AnyverPaths{}

	if value, ok := os.LookupEnv("ANYVER_CONFIG"); ok {
		paths.ConfigFile = value
	} else {
		paths.ConfigFile = defaultAnyverConfigFilePath
	}

	paths.RootDir = filepath.Dir(paths.ConfigFile)

	if value, ok := os.LookupEnv("ANYVER_APPS_DIR"); ok {
		paths.AppsDir = value
	} else {
		paths.AppsDir = defaultAnyverAppsDirPath
	}

	return paths
}

func GetAnyverPaths(c *cli.Context) *AnyverPaths {
	if c != nil {
		return getAnyverPathsFromCLIContext(c)
	}
	return getAnyverPathsFromEnvsOrDefault()
}
