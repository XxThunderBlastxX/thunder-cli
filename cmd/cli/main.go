package main

import (
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"

	"github.com/XxThunderBlastxX/thunder-cli/internal/cmd"
	"github.com/XxThunderBlastxX/thunder-cli/internal/config"
)

func init() {
	//Get home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	//Get absolute path of config directory
	dirPath, err := filepath.Abs(homeDir + "/.config/thunder-cli")
	if err != nil {
		return
	}

	// Check if directory does not exist
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			return
		}
	}

	// Create a new config.json
	configPath := dirPath + "/config.json"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		file, err := os.Create(configPath)
		if err != nil {
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				return
			}
		}(file)
	}
}

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
