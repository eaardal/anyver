package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var AppsCommand = &cli.Command{
	Name:  "apps",
	Usage: "Print the currently active app aliases",
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

		appFiles, err := os.ReadDir(paths.AppsDir)
		if err != nil {
			return err
		}

		write("Active app aliases in %s:", paths.AppsDir)

		if len(appFiles) == 0 {
			write("No active app aliases")
			return nil
		}

		for _, appFile := range appFiles {
			appAliasFilePath := filepath.Join(paths.AppsDir, appFile.Name())

			file, err := os.ReadFile(appAliasFilePath)
			if err != nil {
				return fmt.Errorf("failed to read %s: %v", appFile.Name(), err)
			}

			out, err := exec.Command("which", appFile.Name()).CombinedOutput()
			if err != nil {
				return fmt.Errorf("failed to run which %s: %v", appFile.Name(), err)
			}
			trimmedOut := strings.TrimSuffix(string(out), "\n")

			whichOut := green(trimmedOut)
			if trimmedOut != appAliasFilePath {
				whichOut = red(trimmedOut)
			}
			which := fmt.Sprintf("which %s: %s", yellow(appFile.Name()), whichOut)

			appName := yellow(appFile.Name())
			appScript := cyan(fmt.Sprintf("%q", string(file)))
			write(fmt.Sprintf("%s: %s - %s", appName, appScript, which))
		}

		return nil
	},
}
