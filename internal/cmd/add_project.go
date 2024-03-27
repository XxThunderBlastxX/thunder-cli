package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"

	"github.com/XxThunderBlastxX/thunder-cli/pkg/view"
)

func AddProjectAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		if _, err := tea.NewProgram(view.InitialModel()).Run(); err != nil {
			fmt.Printf("could not start program: %s\n", err)
			os.Exit(1)
		}
		return nil
	}
}
