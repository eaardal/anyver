package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const SystemVersion = "_system"

var (
	ErrNoActive            = errors.New("no active version is configured for the app")
	ErrNoScript            = errors.New("version script not found")
	ErrNoApp               = errors.New("app not found")
	ErrActiveVersionBroken = errors.New("active version refers to a nonexisting version")
	ErrNoVersion           = errors.New("version not found")
)

var DefaultAnyverYamlPath = fmt.Sprintf("%s/config.yaml", AnyverDirPath)

type AnyverYaml struct {
	Active   map[string]string            `yaml:"active"`
	Apps     map[string]map[string]string `yaml:"apps"`
	Contexts map[string]map[string]string `yaml:"contexts"`
}

func (a AnyverYaml) ActiveVersionName(app string) (versionName string, err error) {
	for appName, activeVersion := range a.Active {
		if appName == app {
			return activeVersion, nil
		}
	}
	return "", ErrNoActive
}

func (a AnyverYaml) FindVersionScript(appName string, versionName string) (versionScript string, err error) {
	foundApp := false

	for yamlAppName, yamlAppVersions := range a.Apps {
		if yamlAppName == appName {
			for appVersionName, appVersionScript := range yamlAppVersions {
				if versionName == appVersionName {
					return appVersionScript, nil
				}
			}

			foundApp = true
			break
		}
	}

	if foundApp {
		return "", ErrNoScript
	}
	return "", ErrNoApp
}

func (a AnyverYaml) FindActiveVersionScript(appName string) (versionName string, versionScript string, err error) {
	activeVersion, err := a.ActiveVersionName(appName)
	if err != nil {
		return "", "", err
	}

	script, err := a.FindVersionScript(appName, activeVersion)
	if err != nil {
		if err == ErrNoScript {
			return "", "", ErrActiveVersionBroken
		}
		return "", "", err
	}

	return activeVersion, script, nil
}

type Version struct {
	Name    string
	Script  string
	AppName string
}

func (a AnyverYaml) FindContextScripts(contextName string) (versions []Version, err error) {
	for ctxName, ctxVersions := range a.Contexts {
		if ctxName == contextName {
			for appName, versionName := range ctxVersions {
				versionScript, err := a.FindVersionScript(appName, versionName)
				if err != nil {
					return nil, err
				}
				versions = append(versions, Version{
					Name:    versionName,
					Script:  versionScript,
					AppName: appName,
				})
			}
		}
	}
	return versions, nil
}

func (a AnyverYaml) UseVersion(appName string, versionName string) error {
	foundVersion := false

	if versionName == SystemVersion {
		a.Active[appName] = SystemVersion
		return nil
	}

	for yamlAppName, yamlAppVersions := range a.Apps {
		if yamlAppName == appName {
			for appVersionName := range yamlAppVersions {
				if appVersionName == versionName {
					a.Active[appName] = versionName
					foundVersion = true
					break
				}
			}

			if !foundVersion {
				return ErrNoVersion
			}

			break
		}
	}

	if foundVersion {
		return nil
	}
	return ErrNoApp
}

func ReadAnyverYaml(file string) (*AnyverYaml, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var data AnyverYaml
	if err := yaml.Unmarshal(content, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func SaveAnyverYaml(anyverYaml *AnyverYaml, filePath string) error {
	content, err := yaml.Marshal(*anyverYaml)
	if err != nil {
		return err
	}

	stat, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, content, stat.Mode())
}
