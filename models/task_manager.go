package models

import (
	"errors"
	"minitodo/storage"
	"strings"
)

type TaskManager struct {
	tasks      []Task
	pathToFile string
}

func NewTaskManager(filename string) (*TaskManager, error) {
	tasks, err := storage.LoadTasks(filename)

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
	return storage.SaveTasks(tm.pathToFile, tm.tasks)
}

func (tm *TaskManager) Delete(index int) error {
	if index >= len(tm.tasks) || index < 0 {
		return errors.New("invalid index")
	}

	tm.tasks = append(tm.tasks[:index], tm.tasks[index+1:]...)
	return storage.SaveTasks(tm.pathToFile, tm.tasks)
}

func (tm *TaskManager) Add(text string) error {
	if strings.TrimSpace(text) == "" {
		return errors.New("text is empty")
	}

	task := Task{
		Text: strings.TrimSpace(text),
		Done: false,
	}

	tm.tasks = append(tm.tasks, task)

	return storage.SaveTasks(tm.pathToFile, tm.tasks)
}

func (tm *TaskManager) GetAll() []Task {
	newCopy := make([]Task, len(tm.tasks))
	copy(newCopy, tm.tasks)
	return newCopy
}

func (tm *TaskManager) GetCount() int {
	return len(tm.tasks)
}
