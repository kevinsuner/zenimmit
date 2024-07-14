package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
)

const (
	subjectLength = 50
	lineLength    = 72
)

var (
	subject strings.Builder
	stepper = newStepper()
)

func main() {
	viper.SetConfigName("zenimmit")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("failed to read config file: %s\n", err.Error())
		os.Exit(1)
	}

	program := tea.NewProgram(newList(viper.GetStringSlice("types")))
	if _, err := program.Run(); err != nil {
		fmt.Printf("failed to run program: %s\n", err.Error())
		os.Exit(1)
	}
}
