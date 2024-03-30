package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"

	c "github.com/XxThunderBlastxX/thunder-cli/internal/cmd"
)

func main() {
	f, err := tea.LogToFile("bubbletea.log", "thunder-cli")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	app := &cli.App{
		Name:   "thunder",
		Usage:  "used to interact with thunder-api",
		Action: c.RootAction(),
		Commands: []*cli.Command{
			{
				Name:  "project",
				Usage: "used to interact with projects",
				Subcommands: []*cli.Command{
					{
						Name:   "add",
						Usage:  "add a project",
						Action: c.AddProjectAction(),
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
