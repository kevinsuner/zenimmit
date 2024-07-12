package list

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	cursor  int
	options []string
}

func New(options []string) *Model {
	return &Model{
		options: options,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil // noop
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", " ":
			// TODO: Call stepper handler
		}
	}

	return m, nil
}

func (m *Model) View() string {
	var builder strings.Builder
	builder.WriteString("What type of change are you doing?\n\n")

	for i, option := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		builder.WriteString(fmt.Sprintf("%s %s\n", cursor, option))
	}

	builder.WriteString("\nPress q to quit.\n")
	return builder.String()
}
