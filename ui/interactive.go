package ui

import (
	"fmt"
	"minitodo/models"
	"minitodo/storage"

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
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case " ":
			m.tasks[m.cursor].Done = !m.tasks[m.cursor].Done
			storage.SaveTasks(m.pathToFile, m.tasks)
		case "esc":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Tasks list:\n\n"

	for i, task := range m.tasks {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if task.Done {
			checked = "X"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, task.Text)
	}

	s += "\n Use SPACE for toggle done, ESC for quit"
	return s
}

func RunInteractiveList(tasks []models.Task, filename string) error {
	m := model{
		cursor:     0,
		tasks:      tasks,
		pathToFile: filename,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
