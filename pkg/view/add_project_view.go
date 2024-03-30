package view

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"

	"github.com/XxThunderBlastxX/thunder-cli/internal/model"
	"github.com/XxThunderBlastxX/thunder-cli/internal/service"
	"github.com/XxThunderBlastxX/thunder-cli/pkg/style"
	"github.com/XxThunderBlastxX/thunder-cli/pkg/utils"
)

type AddProjectViewState int

// States of the AddProjectView
const (
	editing AddProjectViewState = iota
	successAddProject
	errorAddProject
	sendingAddProjectReq
)

// Values to be taken from the user
var (
	name        string
	link        string
	description string
	stack       string
)

type (
	successAddProjectMsg struct{}
	errorAddProjectMsg   struct {
		err error
	}
)

type AddProjectViewModel struct {
	// holds the current state of the view
	state AddProjectViewState

	// form to take input from user
	form *huh.Form

	// quitting is true when user wants to quit the view
	quitting bool

	err error

	// help is a help.Model that displays help text
	help help.Model

	// loadingSpinner is a spinner.Model that displays a spinner while
	loadingSpinner spinner.Model

	// keymap is the keybindings for the view
	keymap utils.KeyMap

	// isSubmitted is true when user has submitted the form
	isSubmitted bool

	// projService is the service instance for project
	projService *service.IProject
}

func NewAddProjectView(projService *service.IProject) AddProjectViewModel {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Name of the Project ?").
				Placeholder("Enter project name").
				CharLimit(50).
				Prompt("?").
				Value(&name),

			huh.NewInput().
				Title("Link to the Project ?").
				Placeholder("Enter project link").
				CharLimit(50).
				Prompt("?").
				Value(&link),

			huh.NewInput().
				Title("Description of the Project ?").
				Placeholder("Enter short description of project").
				CharLimit(100).
				Prompt("?").
				Value(&description),

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
		projService:    projService,
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
	case successAddProjectMsg:
		p.state = successAddProject
		return p, tea.Quit
	case errorAddProjectMsg:
		// TODO: Handle error properly
		log.Println("Error adding project: ", msg.err)
		p.err = errors.New("could not add project. Please try again")
		return p, tea.Quit
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

	// If the form is submitted, send the add project request
	if p.form.State == huh.StateCompleted && p.state == editing {
		p.state = sendingAddProjectReq
		cmds = append(cmds, p.addProjectCmd())
		cmds = append(cmds, p.loadingSpinner.Tick)
	}

	switch p.state {
	case sendingAddProjectReq:
		p.loadingSpinner, cmd = p.loadingSpinner.Update(msg)
		cmds = append(cmds, cmd)
	default:
	}

	return p, tea.Batch(cmds...)
}

func (p AddProjectViewModel) View() string {
	var s strings.Builder

	switch p.form.State {
	case huh.StateCompleted:
		s.WriteString(fmt.Sprintf("Wow Cool! Adding your %s project.", name))
		s.WriteString("\n\n")
		return style.BorderStyle.Render(s.String() + style.SecondaryStyle.Render(p.loadingSpinner.View()) + "Please wait a moment...")
	default:
		return p.form.View()
	}

}

// addProjectCmd returns a tea.Cmd that sends the add project request
func (p AddProjectViewModel) addProjectCmd() tea.Cmd {
	return func() tea.Msg {
		proj := model.Project{
			Name:        name,
			Link:        link,
			Description: description,
			Stacks:      nil,
		}

		// Split the stack string by comma and add to the project model
		stacks := strings.Split(stack, ",")
		for _, s := range stacks {
			proj.Stacks = append(proj.Stacks, model.TechStack{Name: s})
		}

		if err := (*p.projService).AddProject(proj); err != nil {
			return errorAddProjectMsg{
				err: err,
			}
		}
		return successAddProjectMsg{}
	}
}
