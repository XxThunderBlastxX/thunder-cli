package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v2"
)

const logo = `
████████╗██╗  ██╗██╗   ██╗███╗   ██╗██████╗ ███████╗██████╗        ██████╗██╗     ██╗
╚══██╔══╝██║  ██║██║   ██║████╗  ██║██╔══██╗██╔════╝██╔══██╗      ██╔════╝██║     ██║
   ██║   ███████║██║   ██║██╔██╗ ██║██║  ██║█████╗  ██████╔╝█████╗██║     ██║     ██║
   ██║   ██╔══██║██║   ██║██║╚██╗██║██║  ██║██╔══╝  ██╔══██╗╚════╝██║     ██║     ██║
   ██║   ██║  ██║╚██████╔╝██║ ╚████║██████╔╝███████╗██║  ██║      ╚██████╗███████╗██║
   ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚═════╝ ╚══════╝╚═╝  ╚═╝       ╚═════╝╚══════╝╚═╝
`

var (
	textStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#894BEF"))
)

func RootAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		fmt.Println(textStyle.Render(logo))
		return nil
	}
}
