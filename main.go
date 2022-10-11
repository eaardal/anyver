package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:        "Anything Version Manager",
		Usage:       "A version manager to temporarily replace any program with any script",
		Version:     "0.0.1",
		Description: "See readme at https://github.com/eaardal/anyver for instructions",
		Commands: []*cli.Command{
			InitCommand,
			ConfigCommand,
			AppsCommand,
			UseCommand,
			RestoreCommand,
			RestoreAllCommand,
			ContextCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
