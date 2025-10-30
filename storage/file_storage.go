package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"minitodo/models"
	"os"
)

type FileStorage struct {
	filename string
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{filename: filename}
}

func (fs *FileStorage) Load() ([]models.Task, error) {
	data, err := os.ReadFile(fs.filename)

	if errors.Is(err, os.ErrNotExist) {
		if err := os.WriteFile(fs.filename, []byte("[]"), 0644); err != nil {
			return nil, fmt.Errorf("error creating file: %w", err)
		}
		return []models.Task{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	if len(data) == 0 {
		return []models.Task{}, nil
	}

	var tasks []models.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("error parsing json: %w", err)
	}

	return tasks, nil
}

func (fs *FileStorage) Save(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return fmt.Errorf("cannot serialize tasks for file %w", err)
	}

	err = os.WriteFile(fs.filename, data, 0644)

	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
