package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"gtfo-tier/tiering"
	"strings"
)

type keyMap struct {
	Up   key.Binding
	Down key.Binding
	Quit key.Binding
}

var (
	KeyMap = keyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "Move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "Move down"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q/CTRL+C", "Quit"),
		),
	}
)

type model struct {
	choices []tiering.Situation
	cursor  int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, KeyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, KeyMap.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, KeyMap.Down):
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var s strings.Builder
	s.Grow(512)

	fmt.Fprintln(&s, "Situations:")

	for i, choice := range m.choices {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		fmt.Fprintf(&s, "  %s %s\n", cursor, choice)
	}

	fmt.Fprintf(&s, "\nPress q to quit\n")

	return s.String()
}

func NewModel() model {
	return model{
		choices: tiering.AllSituations,
		cursor:  0,
	}
}
