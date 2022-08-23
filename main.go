package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "Anything Version Manager",
		Usage: "",
		Commands: []*cli.Command{
			ConfigCommand,
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
