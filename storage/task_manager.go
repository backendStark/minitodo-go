package storage

import (
	"errors"
	"minitodo/models"
	"strings"
)

type TaskManager struct {
	tasks      []models.Task
	pathToFile string
}

func NewTaskManager(filename string) (*TaskManager, error) {
	tasks, err := LoadTasks(filename)

	if err != nil {
		return nil, err
	}

	return &TaskManager{tasks: tasks, pathToFile: filename}, nil
}

func (tm *TaskManager) Toggle(index int) error {
	if index >= len(tm.tasks) || index < 0 {
		return errors.New("invalid index")
	}
	tm.tasks[index].Toggle()
	return SaveTasks(tm.pathToFile, tm.tasks)
}

func (tm *TaskManager) Delete(index int) error {
	if index >= len(tm.tasks) || index < 0 {
		return errors.New("invalid index")
	}

	tm.tasks = append(tm.tasks[:index], tm.tasks[index+1:]...)
	return SaveTasks(tm.pathToFile, tm.tasks)
}

func (tm *TaskManager) Add(text string) error {
	if strings.TrimSpace(text) == "" {
		return errors.New("text is empty")
	}

	task := models.Task{
		Text: strings.TrimSpace(text),
		Done: false,
	}

	tm.tasks = append(tm.tasks, task)

	return SaveTasks(tm.pathToFile, tm.tasks)
}

func (tm *TaskManager) GetAll() []models.Task {
	newCopy := make([]models.Task, len(tm.tasks))
	copy(newCopy, tm.tasks)
	return newCopy
}

func (tm *TaskManager) GetCount() int {
	return len(tm.tasks)
}
