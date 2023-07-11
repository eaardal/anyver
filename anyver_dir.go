package main

import (
	"fmt"
	"os"
	"path"
)

func EnsureAnyverAppsDirExists() (string, error) {
	paths := GetAnyverPaths(nil)

	if err := os.MkdirAll(paths.AppsDir, os.ModePerm); err != nil {
		return paths.AppsDir, fmt.Errorf("failed to ensure %s exists: %+v", paths.AppsDir, err)
	}
	return paths.AppsDir, nil
}

func SetAppAlias(appName string, aliasContent string) error {
	appsDir, err := EnsureAnyverAppsDirExists()
	if err != nil {
		return err
	}

	aliasFilePath := path.Join(appsDir, appName)
	if FileExists(aliasFilePath) {
		if err := os.Remove(aliasFilePath); err != nil {
			return fmt.Errorf("failed to remove alias file %s: %+v", aliasFilePath, err)
		}
	}

	if err := os.WriteFile(aliasFilePath, []byte(aliasContent), 0777); err != nil {
		return fmt.Errorf("failed to create alias file %s: %+v", aliasFilePath, err)
	}

	return nil
}

func DeleteAppAlias(appName string) error {
	appsDir, err := EnsureAnyverAppsDirExists()
	if err != nil {
		return err
	}

	aliasFilePath := path.Join(appsDir, appName)
	if FileExists(aliasFilePath) {
		if err := os.Remove(aliasFilePath); err != nil {
			return fmt.Errorf("failed to remove alias file %s: %+v", aliasFilePath, err)
		}
	}

	return nil
}
