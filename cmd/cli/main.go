package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"

	"github.com/XxThunderBlastxX/thunder-cli/internal/cmd"
	"github.com/XxThunderBlastxX/thunder-cli/internal/config"
)

func main() {
	f, err := tea.LogToFile("thunder-cli.log", "thunder-cli")
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	// Load app config
	appConfig, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:   "thunder",
		Usage:  "used to interact with thunder-api",
		Action: cmd.RootAction(),
		Commands: []*cli.Command{
			{
				Name:  "project",
				Usage: "used to interact with projects",
				Subcommands: []*cli.Command{
					{
						Name:   "add",
						Usage:  "add a project",
						Action: cmd.AddProjectAction(),
					},
				},
			},
			{
				Name:   "login",
				Usage:  "login to thunder-api",
				Action: cmd.LoginAction(appConfig),
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
