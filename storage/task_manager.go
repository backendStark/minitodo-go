package storage

import (
	"errors"
	"minitodo/models"
	"strings"
)

type TaskManager struct {
	tasks   []models.Task
	storage Storage
}

func NewTaskManager(filename string) (*TaskManager, error) {
	storage := NewFileStorage(filename)

	return NewTaskManagerWithStorage(storage)
}

func NewTaskManagerWithStorage(storage Storage) (*TaskManager, error) {
	tasks, err := storage.Load()

	if err != nil {
		return nil, err
	}

	return &TaskManager{
		tasks:   tasks,
		storage: storage,
	}, nil
}

func (tm *TaskManager) Toggle(index int) error {
	if index >= len(tm.tasks) || index < 0 {
		return errors.New("invalid index")
	}
	tm.tasks[index].Toggle()
	return tm.storage.Save(tm.tasks)
}

func (tm *TaskManager) Delete(index int) error {
	if index >= len(tm.tasks) || index < 0 {
		return errors.New("invalid index")
	}

	tm.tasks = append(tm.tasks[:index], tm.tasks[index+1:]...)
	return tm.storage.Save(tm.tasks)
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

	return tm.storage.Save(tm.tasks)
}

func (tm *TaskManager) GetAll() []models.Task {
	newCopy := make([]models.Task, len(tm.tasks))
	copy(newCopy, tm.tasks)
	return newCopy
}

func (tm *TaskManager) GetCount() int {
	return len(tm.tasks)
}
