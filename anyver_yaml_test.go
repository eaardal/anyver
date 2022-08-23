package main

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestActiveVersionName_WhenAppHasActiveVersionSet_ReturnsVersionName(t *testing.T) {
	const rawYaml = `
active:
  yq: system`
	anyverYaml := readTestYaml(t, rawYaml)
	versionName, err := anyverYaml.ActiveVersionName("yq")
	assert.NoError(t, err)
	assert.Equal(t, "system", versionName)
}

func TestActiveVersionName_WhenAppMissingActiveVersion_ReturnsNoActiveError(t *testing.T) {
	const rawYaml = `
active:
  yq: system`
	anyverYaml := readTestYaml(t, rawYaml)
	versionName, err := anyverYaml.ActiveVersionName("foo")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoActive)
	assert.Empty(t, versionName)
}

func TestFindVersionScript_WhenAppExistsAndVersionExists_ReturnsVersionScript(t *testing.T) {
	const rawYaml = `
apps:
  yq:
    system: /usr/local/bin/yq
    tv2_rp_docker: "docker run foo"
  mockery:
    system: echo "run mockery from usr/local/bin"
    tv2_rp_docker: echo "run mockery in docker"`
	anyverYaml := readTestYaml(t, rawYaml)
	versionName, err := anyverYaml.FindVersionScript("yq", "tv2_rp_docker")
	assert.NoError(t, err)
	assert.Equal(t, "docker run foo", versionName)
}

func TestFindVersionScript_WhenAppExistsButVersionDoesNotExist_ReturnsNoScriptError(t *testing.T) {
	const rawYaml = `
apps:
  yq:
    system: /usr/local/bin/yq
    tv2_rp_docker: "docker run foo"
  mockery:
    system: echo "run mockery from usr/local/bin"
    tv2_rp_docker: echo "run mockery in docker"`
	anyverYaml := readTestYaml(t, rawYaml)
	versionName, err := anyverYaml.FindVersionScript("yq", "other_version")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoScript)
	assert.Empty(t, versionName)
}

func TestFindVersionScript_WhenAppDoesNotExist_ReturnsNoAppError(t *testing.T) {
	const rawYaml = `
apps:
  yq:
    system: /usr/local/bin/yq
    tv2_rp_docker: "docker run foo"
  mockery:
    system: echo "run mockery from usr/local/bin"
    tv2_rp_docker: echo "run mockery in docker"`
	anyverYaml := readTestYaml(t, rawYaml)
	versionName, err := anyverYaml.FindVersionScript("other_app", "system")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoApp)
	assert.Empty(t, versionName)
}

func TestFindActiveVersionScript_WhenAppExistsAndActiveVersionExists_ReturnsVersionScript(t *testing.T) {
	const rawYaml = `
active:
  yq: tv2_rp_docker
apps:
  yq:
    system: /usr/local/bin/yq
    tv2_rp_docker: "docker run foo"
  mockery:
    system: echo "run mockery from usr/local/bin"
    tv2_rp_docker: echo "run mockery in docker"`
	anyverYaml := readTestYaml(t, rawYaml)
	versionName, versionScript, err := anyverYaml.FindActiveVersionScript("yq")
	assert.NoError(t, err)
	assert.Equal(t, "tv2_rp_docker", versionName)
	assert.Equal(t, "docker run foo", versionScript)
}

func TestFindActiveVersionScript_WhenAppExistsButItHasNoActiveVersion_ReturnsNoScriptError(t *testing.T) {
	const rawYaml = `
active:
  yq: tv2_rp_docker
apps:
  yq:
    system: /usr/local/bin/yq
    tv2_rp_docker: "docker run foo"
  mockery:
    system: echo "run mockery from usr/local/bin"
    tv2_rp_docker: echo "run mockery in docker"`
	anyverYaml := readTestYaml(t, rawYaml)
	versionName, versionScript, err := anyverYaml.FindActiveVersionScript("mockery")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoActive)
	assert.Empty(t, versionName)
	assert.Empty(t, versionScript)
}

func TestFindActiveVersionScript_WhenAppDoesNotExist_ReturnsNoAppError(t *testing.T) {
	const rawYaml = `
active:
  yq: system
  other_app: system
apps:
  yq:
    system: /usr/local/bin/yq
    tv2_rp_docker: "docker run foo"
  mockery:
    system: echo "run mockery from usr/local/bin"
    tv2_rp_docker: echo "run mockery in docker"`
	anyverYaml := readTestYaml(t, rawYaml)
	versionName, versionScript, err := anyverYaml.FindActiveVersionScript("other_app")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNoApp)
	assert.Empty(t, versionName)
	assert.Empty(t, versionScript)
}

func TestFindActiveVersionScript_WhenActiveVersionRefersToNonExistingVersion_ReturnsActiveVersionBrokenError(t *testing.T) {
	const rawYaml = `
active:
  yq: nonexisting
apps:
  yq:
    system: /usr/local/bin/yq
    tv2_rp_docker: "docker run foo"
  mockery:
    system: echo "run mockery from usr/local/bin"
    tv2_rp_docker: echo "run mockery in docker"`
	anyverYaml := readTestYaml(t, rawYaml)
	versionName, versionScript, err := anyverYaml.FindActiveVersionScript("yq")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrActiveVersionBroken)
	assert.Empty(t, versionName)
	assert.Empty(t, versionScript)
}

func readTestYaml(t *testing.T, yamlContent string) *AnyverYaml {
	var data AnyverYaml
	if err := yaml.Unmarshal([]byte(yamlContent), &data); err != nil {
		t.Fatalf("failed to read fake/test yaml: %+v", err)
	}
	return &data
}
