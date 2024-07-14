package main

import (
	"fmt"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
)

const (
	commitType int8 = iota
	hasBreakingChange
	commitScope
	commitSubject
	hasCommitBody
	commitBody
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
			hasBreakingChange: {
				question: "Is it a breaking change?",
			},
			commitScope: {
				question: "What is the scope of the change?",
			},
			commitSubject: {
				question: "What is the subject of the change?",
			},
			hasCommitBody: {
				question: "Do you want add more information?",
			},
			commitBody: {
				question: "What is the body of the change?",
			},
		},
	}
}

func (s *Stepper) Next(data string) (tea.Model, tea.Cmd) {
	switch s.kind {
	case commitType:
		subject.WriteString(data)
		s.kind++
		return newList([]string{"yes", "no"}), nil

	case hasBreakingChange:
		if data == "yes" {
			subject.WriteString("!")
		}

		s.kind++
		return newList(viper.GetStringSlice("scopes")), nil

	case commitScope:
		subject.WriteString(fmt.Sprintf("(%s): ", data))
		s.kind++
		return newInput(), nil

	case commitSubject:
		if len(data) == 0 {
			return newInput(), nil
		}

		subject.WriteString(data)
		s.kind++
		return newList([]string{"yes", "no"}), nil

	case hasCommitBody:
		if data == "yes" {
			s.kind++
			return newTextarea(), nil
		}

		out, err := exec.Command("git", "commit", "-m", subject.String()).
			CombinedOutput()
		if err != nil {
			return newText(
				fmt.Sprintf("failed to perform commit: %s\n", err.Error()),
			), tea.Quit
		}

		return newText(string(out)), tea.Quit

	case commitBody:
		if len(data) == 0 {
			return newTextarea(), nil
		}

		out, err := exec.Command(
			"git",
			"commit",
			"-m",
			subject.String(),
			"-m",
			data,
		).CombinedOutput()
		if err != nil {
			return newText(
				fmt.Sprintf("failed to perform commit: %s\n", err.Error()),
			), tea.Quit
		}

		return newText(string(out)), tea.Quit
	}

	return nil, tea.Quit
}

func (s *Stepper) Current() Step {
	return s.steps[s.kind]
}
