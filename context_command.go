package main

import (
	"github.com/urfave/cli/v2"
)

var ContextCommand = &cli.Command{
	Name:  "context",
	Usage: "Manage multiple apps in one go with contexts",
	Subcommands: []*cli.Command{
		ContextUseCommand,
		ContextRestoreCommand,
	},
}
