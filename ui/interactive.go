package ui

import (
	"fmt"
	"minitodo/models"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor     int
	tasks      []models.Task
	pathToFile string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			fmt.Println("Переводим лист вверх (уменьшим cursor на 1)")
		case "down":
			fmt.Println("Переводим лист вниз (увеличим cursor на 1)")
		case " ":
			fmt.Println("Toggle для чеклиста")
		case "esc":
			fmt.Println("Выходим из режима list")
		}
	}
}
