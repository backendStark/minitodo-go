package storage

import (
	"errors"
	"minitodo/models"
	"sort"
	"strings"
)

type TaskManager struct {
	tasks    []models.Task
	storage  Storage
	sortMode models.SortMode
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
		tasks:    tasks,
		storage:  storage,
		sortMode: models.SortByStatus,
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

func (tm *TaskManager) Sort() error {
	comparator := func(i, j int) bool {
		switch tm.sortMode {
		case models.SortByStatus:
			if tm.tasks[i].Done != tm.tasks[j].Done {
				return !tm.tasks[i].Done
			}
		case models.SortByStatusReverse:
			if tm.tasks[i].Done != tm.tasks[j].Done {
				return tm.tasks[i].Done
			}
		}
		return false
	}

	sort.SliceStable(tm.tasks, comparator)
	return tm.storage.Save(tm.tasks)
}

func (tm *TaskManager) ToggleSortMode() error {
	if tm.sortMode == models.SortByStatus {
		tm.sortMode = models.SortByStatusReverse
	} else {
		tm.sortMode = models.SortByStatus
	}

	return tm.Sort()
}

func (tm *TaskManager) GetSortMode() models.SortMode {
	return tm.sortMode
}

func (tm *TaskManager) GetDoneCount() int {
	doneCount := 0
	for _, task := range tm.GetAll() {
		if task.Done {
			doneCount++
		}
	}

	return doneCount
}
