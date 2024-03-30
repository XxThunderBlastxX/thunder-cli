package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"

	"github.com/XxThunderBlastxX/thunder-cli/internal/service"
	"github.com/XxThunderBlastxX/thunder-cli/pkg/style"
	"github.com/XxThunderBlastxX/thunder-cli/pkg/view"
)

func AddProjectAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		h := &http.Client{
			Timeout: time.Minute * 2,
		}
		s := service.NewProjectService(h, "http://localhost:4040")

		p, err := tea.NewProgram(view.NewAddProjectView(&s)).Run()
		if err != nil {
			fmt.Printf("could not start program: %s\n", err)
			os.Exit(1)
		}

		m := p.(view.AddProjectViewModel)

		if m.Err != nil {
			fmt.Println(style.ErrorStyle.PaddingLeft(1).Render(m.Err.Error()))
		}

		if !m.Abort && m.IsSubmitted && m.Err == nil {
			fmt.Println(style.AccentStyle.PaddingLeft(1).Render("> ") + "Project added successfully! 🎉")
		}

		return nil
	}
}
