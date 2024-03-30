package style

import "github.com/charmbracelet/lipgloss"

const (
	purpleColor = "#874BFC"
	pinkColor   = "#FF5F87"
	tintColor   = "#07F996"
)

var (
	Base = lipgloss.NewStyle()

	PaddedStyle = Base.Padding(1)

	PrimaryStyle   = Base.Foreground(lipgloss.Color(purpleColor))
	SecondaryStyle = Base.Foreground(lipgloss.Color(pinkColor))
	AccentStyle    = Base.Foreground(lipgloss.Color(tintColor))

	BorderStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true, true, true, true).BorderForeground(lipgloss.Color(purpleColor)).Padding(1).Margin(1)

	ErrorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5F87"))
)
