package view

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"

	"github.com/XxThunderBlastxX/thunder-cli/internal/model"
	"github.com/XxThunderBlastxX/thunder-cli/pkg/utils"
)

type AddProjectViewState int

const (
	successAddProject AddProjectViewState = iota
	editing
)

var (
	proj  = &model.Project{}
	stack string
)

type AddProjectViewModel struct {
	state          AddProjectViewState
	form           *huh.Form
	quitting       bool
	err            error
	help           help.Model
	loadingSpinner spinner.Model
	keymap         utils.KeyMap
	isSubmitted    bool
}

func NewAddProjectView() AddProjectViewModel {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Name of the Project ?").
				Placeholder("Enter project name").
				CharLimit(50).
				Prompt("?").
				Value(&proj.Name),

			huh.NewInput().
				Title("Link to the Project ?").
				Placeholder("Enter project link").
				CharLimit(50).
				Prompt("?").
				Value(&proj.Link),

			huh.NewInput().
				Title("Description of the Project ?").
				Placeholder("Enter short description of project").
				CharLimit(100).
				Prompt("?").
				Value(&proj.Description),

			huh.NewInput().
				Title("Tech Stack used in the Project ?").
				Placeholder("Enter tech stack with comma seperated values").
				CharLimit(50).
				Prompt("?").
				Value(&stack),

			huh.NewConfirm().
				Title("Are you sure want to submit ?").
				Affirmative("Submit").
				Negative("Cancel"),
		).WithTheme(huh.ThemeCharm()),
	)

	loadingSpinner := spinner.New()
	loadingSpinner.Spinner = spinner.Dot

	m := AddProjectViewModel{
		state:          editing,
		help:           help.New(),
		loadingSpinner: loadingSpinner,
		form:           form,
		isSubmitted:    false,
		quitting:       false,
		keymap:         utils.DefaultKeyMap(),
	}

	return m
}

func (p AddProjectViewModel) Init() tea.Cmd {
	return tea.Batch(
		p.form.Init(),
	)
}

func (p AddProjectViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, p.keymap.Quit):
			p.quitting = true
			return p, tea.Quit
		case key.Matches(msg, p.keymap.NextInput):
			return p, p.form.NextField()
		case key.Matches(msg, p.keymap.PrevInput):
			return p, p.form.PrevField()
		case key.Matches(msg, p.keymap.Send):
			return p, nil
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := p.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		p.form = f
		cmds = append(cmds, cmd)
	}
	return p, tea.Batch(cmds...)
}

func (p AddProjectViewModel) View() string {
	return p.form.View()
}
