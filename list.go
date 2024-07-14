package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type List struct {
	cursor  int
	options []string
}

func newList(options []string) List {
	return List{
		options: options,
	}
}

func (l List) Init() tea.Cmd {
	return nil // noop
}

func (l List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return l, tea.Quit
		case "up", "k":
			if l.cursor > 0 {
				l.cursor--
			}
		case "down", "j":
			if l.cursor < len(l.options)-1 {
				l.cursor++
			}
		case "enter", " ":
			return stepper.Next(l.options[l.cursor]), nil
		}
	}

	return l, nil
}

func (l List) View() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s\n\n", stepper.Current().question))

	for i, option := range l.options {
		cursor := ""
		if l.cursor == i {
			cursor = ">"
		}

		builder.WriteString(fmt.Sprintf("%s %s\n", cursor, option))
	}

	builder.WriteString("\nPress q to quit.\n")
	return builder.String()
}
