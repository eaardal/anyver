package main

import (
	"errors"
	"fmt"
	"os"
	"path"
)

var AnyverDirPath = path.Join(os.Getenv("HOME"), ".anyver")
var AnyverAppsDirPath = path.Join(AnyverDirPath, "apps")

func EnsureAnyverAppsDirExists() error {
	if err := os.MkdirAll(AnyverAppsDirPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to ensure %s exists: %+v", AnyverAppsDirPath, err)
	}
	return nil
}

func SetAppAlias(appName string, aliasContent string) error {
	if err := EnsureAnyverAppsDirExists(); err != nil {
		return err
	}

	aliasFilePath := path.Join(AnyverAppsDirPath, appName)
	if fileExists(aliasFilePath) {
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
	if err := EnsureAnyverAppsDirExists(); err != nil {
		return err
	}

	aliasFilePath := path.Join(AnyverAppsDirPath, appName)
	if fileExists(aliasFilePath) {
		if err := os.Remove(aliasFilePath); err != nil {
			return fmt.Errorf("failed to remove alias file %s: %+v", aliasFilePath, err)
		}
	}

	return nil
}

func fileExists(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
