package ui

import (
	"fmt"
	"minitodo/config"
	"minitodo/models"
	"minitodo/storage"

	textinput "github.com/charmbracelet/bubbles/textinput"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	keyUp     = "up"
	keyDown   = "down"
	keySpace  = " "
	keyDelete = "delete"
	keyEnter  = "enter"
	keyEsc    = "esc"
	keySort   = "s"
)

type model struct {
	cursor      int
	textInput   textinput.Model
	err         error
	taskManager *storage.TaskManager
}

func (m *model) updateInputFocus() {
	if m.cursor == m.taskManager.GetCount() {
		m.textInput.Focus()
	} else {
		m.textInput.Blur()
	}
}

func (m *model) normalizeCursor() {
	if (m.cursor) >= m.taskManager.GetCount() && m.taskManager.GetCount() > 0 {
		m.cursor = m.taskManager.GetCount() - 1
	}

	if m.taskManager.GetCount() == 0 {
		m.cursor = 0
	}
}

func (m *model) renderTask(i int, task models.Task) string {
	cursor := " "

	if m.cursor == i {
		cursor = ">"
	}

	checked := " "
	if task.Done {
		checked = "X"
	}

	return fmt.Sprintf("%s [%s] %s\n", cursor, checked, task.Text)
}

func (m *model) addTask() error {
	if m.cursor == m.taskManager.GetCount() {
		if err := m.taskManager.Add(m.textInput.Value()); err != nil {
			return err
		}
		m.cursor = m.taskManager.GetCount()
		m.textInput.Reset()
		m.updateInputFocus()
	}

	return nil
}

func handleKeyPress(m *model, key string) (tea.Model, tea.Cmd) {
	m.err = nil
	switch key {
	case keyUp:
		if m.cursor > 0 {
			m.cursor--
		}
		m.updateInputFocus()
	case keyDown:
		if m.cursor < m.taskManager.GetCount() {
			m.cursor++
		}
		m.updateInputFocus()
	case keySpace:
		if m.IsNotOnInputField() {
			if err := m.taskManager.Toggle(m.cursor); err != nil {
				m.err = err
				return *m, nil
			}
		}
	case keySort:
		if m.IsNotOnInputField() {
			if err := m.taskManager.ToggleSortMode(); err != nil {
				m.err = err
				return *m, nil
			}
			m.cursor = 0
			m.updateInputFocus()
		}
	case keyDelete:
		if m.IsNotOnInputField() {
			if err := m.taskManager.Delete(m.cursor); err != nil {
				m.err = err
				return *m, nil
			}
			m.normalizeCursor()
		}
	case keyEnter:
		if err := m.addTask(); err != nil {
			m.err = err
			return *m, nil
		}
	case keyEsc:
		return *m, tea.Quit
	}
	return *m, nil
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.updateInputFocus()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		updatedModel, keyCmd := handleKeyPress(&m, msg.String())

		if keyCmd != nil {
			return updatedModel, keyCmd
		}

		if m.cursor == m.taskManager.GetCount() {
			m.textInput, cmd = m.textInput.Update(msg)
		}
	}

	return m, cmd
}

func (m model) View() string {
	doneCount := m.taskManager.GetDoneCount()
	totalCount := m.taskManager.GetCount()
	sortMode := m.taskManager.GetSortMode().String()

	s := fmt.Sprintf("Tasks: %d/%d | [S]ort: %s\n\n", doneCount, totalCount, sortMode)

	if m.taskManager.GetCount() == 0 {
		s += "(no tasks yet)\n\n"
	} else {
		for i, task := range m.taskManager.GetAll() {
			s += m.renderTask(i, task)
		}
	}

	cursor := " "

	if m.cursor == m.taskManager.GetCount() {
		cursor = ">"
	}

	s += cursor + " [ ] " + m.textInput.View()

	s += "\n\nControls:\n"
	s += "  SPACE - toggle task\n"
	s += "  S     - sort tasks\n"
	s += "  DEL   - delete task\n"
	s += "  ENTER - add task (when on input)\n"
	s += "  ESC   - quit"

	if m.err != nil {
		s += fmt.Sprintf("\n\nError: %v", m.err)
	}
	return s
}

func RunInteractiveList(filename string) error {
	ti := textinput.New()
	ti.Placeholder = config.InputPlaceholder
	ti.Width = config.InputWidth
	ti.Prompt = ""

	taskManager, err := storage.NewTaskManager(filename)

	if err != nil {
		return fmt.Errorf("error with creating taskManager: %w", err)
	}

	m := model{
		cursor:      0,
		textInput:   ti,
		taskManager: taskManager,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}

func (m model) IsNotOnInputField() bool {
	return m.cursor < m.taskManager.GetCount()
}
