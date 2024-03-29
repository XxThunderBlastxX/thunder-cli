package utils

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	NextInput key.Binding
	PrevInput key.Binding
	Send      key.Binding
	Quit      key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		NextInput: key.NewBinding(
			key.WithKeys("tab", "enter"),
			key.WithHelp("tab", "next"),
		),
		PrevInput: key.NewBinding(
			key.WithKeys("shift+tab"),
		),
		Send: key.NewBinding(
			key.WithKeys("ctrl+d"),
			key.WithHelp("ctrl+d", "send"),
		),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c", "q", "esc"),
			key.WithHelp("ctrl+c", "quit"),
		),
	}
}
