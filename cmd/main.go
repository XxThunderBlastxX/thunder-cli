package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/XxThunderBlastxX/thunder-cli/pkg/cmd"
)

func main() {
	app := &cli.App{
		Name:   "thunder-cli",
		Usage:  "used to interact with thunder-api",
		Action: cmd.RootAction(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
