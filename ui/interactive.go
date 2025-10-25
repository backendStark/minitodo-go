package ui

import (
	"github.com/charmbracelet/bubbletea"
	"minitodo/models"
)
type model struct {
	cursor int
	tasks []models.Task
	pathToFile string
}

func (m model) Init() tea.Cmd {
	return nil
}