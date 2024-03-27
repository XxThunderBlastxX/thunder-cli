package view

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/XxThunderBlastxX/thunder-cli/internal/model"
	"github.com/XxThunderBlastxX/thunder-cli/pkg/utils"
)

const (
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"
	dotChar           = " • "
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	viewStyle           = lipgloss.NewStyle().MarginLeft(2)

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))

	// Gradient colors we'll use for the progress bar
	ramp = utils.MakeRampStyles("#B14FFF", "#00FFA3", progressBarWidth)
)

type ProjectView struct {
	focusIndex  int
	inputs      []textinput.Model
	cursorMode  cursor.Mode
	project     model.Project
	isSubmitted bool
}

func InitialModel() ProjectView {
	m := ProjectView{
		inputs:      make([]textinput.Model, 4),
		project:     model.Project{},
		isSubmitted: false,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 64

		switch i {
		case 0:
			t.Placeholder = "Name"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Link"
		case 2:
			t.Placeholder = "Description"
			t.CharLimit = 256

		case 3:
			t.Placeholder = "Stacks"
			t.CharLimit = 128
		}

		m.inputs[i] = t
	}

	return m
}

func (p ProjectView) Init() tea.Cmd {
	return textinput.Blink
}

func (p ProjectView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return p, tea.Quit

		// Change cursor mode
		case "ctrl+r":
			p.cursorMode++
			if p.cursorMode > cursor.CursorHide {
				p.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(p.inputs))
			for i := range p.inputs {
				cmds[i] = p.inputs[i].Cursor.SetMode(p.cursorMode)
			}
			return p, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && p.focusIndex == len(p.inputs) {
				p.isSubmitted = true
				return p, nil
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				p.focusIndex--
			} else {
				p.focusIndex++
			}

			if p.focusIndex > len(p.inputs) {
				p.focusIndex = 0
			} else if p.focusIndex < 0 {
				p.focusIndex = len(p.inputs)
			}

			cmds := make([]tea.Cmd, len(p.inputs))
			for i := 0; i <= len(p.inputs)-1; i++ {
				if i == p.focusIndex {
					// Set focused state
					cmds[i] = p.inputs[i].Focus()
					p.inputs[i].PromptStyle = focusedStyle
					p.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				p.inputs[i].Blur()
				p.inputs[i].PromptStyle = noStyle
				p.inputs[i].TextStyle = noStyle

				switch i {
				case 0:
					p.project.Name = p.inputs[i].Value()
				case 1:
					p.project.Link = p.inputs[i].Value()
				case 2:
					p.project.Description = p.inputs[i].Value()
				case 3:
					a := strings.Split(p.inputs[i].Value(), ",")
					p.project.Stacks = nil
					for i := range a {
						p.project.Stacks = append(p.project.Stacks, model.TechStack{Name: a[i]})
					}
				}
			}

			return p, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := p.updateInputs(msg)

	return p, cmd
}

func (p *ProjectView) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(p.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range p.inputs {
		p.inputs[i], cmds[i] = p.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (p ProjectView) View() string {
	if p.isSubmitted {
		return p.submitFormView()
	}
	return p.formView()
}

func (p *ProjectView) formView() string {
	var b strings.Builder

	for i := range p.inputs {
		b.WriteString(p.inputs[i].View())
		if i < len(p.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if p.focusIndex == len(p.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(p.cursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	return viewStyle.Render(b.String())
}

func (p *ProjectView) submitFormView() string {
	return viewStyle.Render("Loading...")
}
