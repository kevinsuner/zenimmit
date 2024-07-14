package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Text struct {
	text string
}

func newText(text string) Text {
	return Text{
		text: text,
	}
}

func (t Text) Init() tea.Cmd {
	return nil
}

func (t Text) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

func (t Text) View() string {
	return t.text
}
