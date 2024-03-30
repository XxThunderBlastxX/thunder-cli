package utils

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

func DefaultKeyMap() *huh.KeyMap {
	k := huh.NewDefaultKeyMap()

	k.Input.Next = key.NewBinding(
		key.WithKeys("tab", "enter", "down"),
		key.WithHelp("tab", "next"),
	)

	k.Input.Prev = key.NewBinding(
		key.WithKeys("shift+tab", "up"),
		key.WithHelp("shift+tab", "back"),
	)

	k.Quit = key.NewBinding(
		key.WithKeys("ctrl+c", "esc"),
		key.WithHelp("ctrl+c", "quit"),
	)

	return k
}
