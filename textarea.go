package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Textarea struct {
	textarea textarea.Model
	err      error
}

func newTextarea() Textarea {
	area := textarea.New()
	area.Placeholder = "E.g. Added a fn to sort an unsorted 'map' using 'sort.Sort()'..."
	area.SetWidth(lineLength)
	area.Focus()

	return Textarea{
		textarea: area,
		err:      nil,
	}
}

func (t Textarea) Init() tea.Cmd {
	return textarea.Blink
}

func (t Textarea) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return t, tea.Quit
		case "ctrl+s":
			model, cmd := stepper.Next(t.textarea.Value())
			if model != nil {
				return model, cmd
			}

			return t, cmd
		}

	case error:
		t.err = msg
		return t, nil
	}

	var cmd tea.Cmd
	t.textarea, cmd = t.textarea.Update(msg)
	return t, cmd
}

func (t Textarea) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n",
		stepper.Current().question,
		t.textarea.View(),
		"Press CTRL+C to quit | CTRL+S to submit")
}
