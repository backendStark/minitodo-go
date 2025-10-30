package storage

import "minitodo/models"

type MockStorage struct {
	tasks     []models.Task
	loadError error
	saveError error
	saveCalls int
}

func NewMockStorage(initialTasks []models.Task) *MockStorage {
	return &MockStorage{
		tasks: initialTasks,
	}
}

func (m *MockStorage) Load() ([]models.Task, error) {
	if m.loadError != nil {
		return nil, m.loadError
	}

	return m.tasks, nil
}

func (m *MockStorage) Save(tasks []models.Task) error {
	m.saveCalls++

	if m.saveError != nil {
		return m.saveError
	}

	m.tasks = tasks
	return nil
}

func (m *MockStorage) GetSaveCalls() int {
	return m.saveCalls
}

func (m *MockStorage) SetLoadError(err error) {
	m.loadError = err
}

func (m *MockStorage) SetSaveError(err error) {
	m.saveError = err
}
