package ui

import (
	"fmt"
	"minitodo/models"
	"minitodo/storage"
	"strings"

	textinput "github.com/charmbracelet/bubbles/textinput"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor     int
	tasks      []models.Task
	pathToFile string
	textInput  textinput.Model
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.cursor == len(m.tasks) {
		m.textInput.Focus()
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
			if m.cursor == len(m.tasks) {
				m.textInput.Focus()
			} else {
				m.textInput.Blur()
			}
		case "down":
			if m.cursor < len(m.tasks) {
				m.cursor++
			}
			if m.cursor == len(m.tasks) {
				m.textInput.Focus()
			} else {
				m.textInput.Blur()
			}
		case " ":
			if m.cursor < len(m.tasks) {
				m.tasks[m.cursor].Done = !m.tasks[m.cursor].Done
				storage.SaveTasks(m.pathToFile, m.tasks)
			}
		case "delete":
			if m.cursor < len(m.tasks) {
				m.tasks = append(m.tasks[:m.cursor], m.tasks[m.cursor+1:]...)

				if (m.cursor) >= len(m.tasks) && len(m.tasks) > 0 {
					m.cursor = len(m.tasks) - 1
				}

				if len(m.tasks) == 0 {
					m.cursor = 0
				}

				storage.SaveTasks(m.pathToFile, m.tasks)
			}
		case "enter":
			if m.cursor == len(m.tasks) && strings.TrimSpace(m.textInput.Value()) != "" {
				task := models.Task{
					Text: strings.TrimSpace(m.textInput.Value()),
					Done: false,
				}

				m.tasks = append(m.tasks, task)

				storage.SaveTasks(m.pathToFile, m.tasks)

				m.textInput.Reset()

				if len(m.tasks) > 0 {
					m.cursor = len(m.tasks)
					m.textInput.Focus()
				}
			}
		case "esc":
			return m, tea.Quit
		}
	}

	if m.cursor == len(m.tasks) {
		m.textInput, cmd = m.textInput.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	s := "Tasks list:\n\n"

	if len(m.tasks) == 0 {
		s += "(no tasks yet)\n\n"
	} else {
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
	}

	cursor := " "

	if m.cursor == len(m.tasks) {
		cursor = ">"
	}

	s += cursor + " [ ] " + m.textInput.View()

	s += "\n\nControls:\n"
	s += "  SPACE - toggle task\n"
	s += "  DEL   - delete task\n"
	s += "  ENTER - add task (when on input)\n"
	s += "  ESC   - quit"
	return s
}

func RunInteractiveList(tasks []models.Task, filename string) error {
	ti := textinput.New()
	ti.Placeholder = "Enter task text"
	ti.Width = 50
	ti.Prompt = ""

	m := model{
		cursor:     0,
		tasks:      tasks,
		pathToFile: filename,
		textInput:  ti,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
