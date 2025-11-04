package storage

import "minitodo/models"

type Storage interface {
	Load() ([]models.Task, error)
	Save(tasks []models.Task) error
}
