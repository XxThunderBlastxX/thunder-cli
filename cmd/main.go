package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	cmd2 "github.com/XxThunderBlastxX/thunder-cli/internal/cmd"
)

func main() {
	app := &cli.App{
		Name:   "thunder",
		Usage:  "used to interact with thunder-api",
		Action: cmd2.RootAction(),
		Commands: []*cli.Command{
			{
				Name:  "project",
				Usage: "used to interact with projects",
				Subcommands: []*cli.Command{
					{
						Name:   "add",
						Usage:  "add a project",
						Action: cmd2.AddProjectAction(),
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
