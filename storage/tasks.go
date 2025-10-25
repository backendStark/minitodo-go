package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"minitodo/models"
	"os"
)

func LoadTasks(filename string) ([]models.Task, error) {
	if data, err := os.ReadFile(filename); err == nil {
		if len(data) == 0 {
			return []models.Task{}, nil
		}

		var tasks []models.Task

		if err := json.Unmarshal(data, &tasks); err != nil {
			return nil, fmt.Errorf("error parsing json %w", err)
		}

		return tasks, nil
	} else if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(filename)

		if err != nil {
			return nil, fmt.Errorf("error creating json file %w", err)
		}

		defer file.Close()

		file.Write([]byte("[]"))

		return []models.Task{}, nil
	} else {
		return nil, fmt.Errorf("file has errors %w", err)
	}
}

func SaveTasks(filename string, tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return fmt.Errorf("cannot serialize tasks for file %w", err)
	}

	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
