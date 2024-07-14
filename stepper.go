package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
)

const (
	commitType int8 = iota
	breakingChange
	commitScope
)

type Stepper struct {
	kind  int8
	steps map[int8]Step
}

type Step struct {
	question string
}

func newStepper() Stepper {
	return Stepper{
		kind: commitType,
		steps: map[int8]Step{
			commitType: {
				question: "What type of change are you performing?",
			},
			breakingChange: {
				question: "Is it a breaking change?",
			},
			commitScope: {
				question: "What is the scope of the change?",
			},
		},
	}
}

func (s *Stepper) Next() tea.Model {
	switch s.kind {
	case commitType:
		s.kind++
		return newList([]string{"yes", "no"})
	case breakingChange:
		s.kind++
		return newList(viper.GetStringSlice("scopes"))
	case commitScope:
		s.kind++
		return nil
	default:
		return nil
	}
}

func (s *Stepper) Current() Step {
	return s.steps[s.kind]
}
