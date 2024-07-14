package main

import (
	"log/slog"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
)

var stepper = newStepper()

func main() {
	viper.SetConfigName("zenimmit")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		slog.Error("failed to read config", "error", err)
		os.Exit(1)
	}

	program := tea.NewProgram(newList(viper.GetStringSlice("types")))
	if _, err := program.Run(); err != nil {
		slog.Error("failed to run program", "error", err)
		os.Exit(1)
	}
}
