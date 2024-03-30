package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"

	"github.com/XxThunderBlastxX/thunder-cli/internal/service"
	"github.com/XxThunderBlastxX/thunder-cli/pkg/view"
)

func AddProjectAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		h := &http.Client{
			Timeout: time.Minute * 2,
		}
		s := service.NewProjectService(h, "http://localhost:4040")

		if _, err := tea.NewProgram(view.NewAddProjectView(&s)).Run(); err != nil {
			fmt.Printf("could not start program: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("Project added successfully! 🎉")
		return nil
	}
}
