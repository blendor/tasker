package storage

import (
	"tasker/model"
)

// Storage defines the interface for task storage operations.
type Storage interface {
	SaveTask(task model.Task) error
	LoadTasks() ([]model.Task, error)
	GetTask(id int) (model.Task, error)
	DeleteTask(id int) error
	UpdateTask(task model.Task) error // New method to update a task
}
