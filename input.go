package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type Input struct {
	input textinput.Model
	err   error
}

func newInput() Input {
	input := textinput.New()
	input.Placeholder = "refactored spaghetti code, now it's linguini"
	input.CharLimit = 50
	input.Width = 50
	input.Focus()

	return Input{input: input, err: nil}
}

func (i Input) Init() tea.Cmd {
	return textinput.Blink
}

func (i Input) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return i, tea.Quit
		case "enter":
			return stepper.Next(i.input.Value()), nil
		}
	case errMsg:
		i.err = msg
		return i, nil
	}

	var cmd tea.Cmd
	i.input, cmd = i.input.Update(msg)
	return i, cmd
}

func (i Input) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n",
		stepper.Current().question,
		i.input.View(),
		"Press q to quit.")
}
