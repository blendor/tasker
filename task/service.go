package task

import (
	"tasker/model"
	"tasker/storage"
)

// Service provides task management functionalities.
type Service struct {
	storage storage.Storage
}

// NewService creates a new task service.
func NewService(storage storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}

// AddTask adds a new task.
func (s *Service) AddTask(description string, priority int, urgency string, importance string) error {
	task := model.NewTask(0, description, priority, urgency, importance) // Assuming ID is auto-generated or managed by storage
	return s.storage.SaveTask(task)
}

// ListTasks lists all tasks.
func (s *Service) ListTasks() ([]model.Task, error) {
	return s.storage.LoadTasks()
}

// UpdateTask updates an existing task.
func (s *Service) UpdateTask(id int, description string, priority int, urgency string, importance string) error {
	task, err := s.storage.GetTask(id)
	if err != nil {
		return err
	}

	task.Description = description
	task.Priority = priority
	task.Urgency = urgency
	task.Importance = importance

	return s.storage.SaveTask(task) // Using SaveTask to handle the update
}

// DeleteTask deletes a task.
func (s *Service) DeleteTask(id int) error {
	return s.storage.DeleteTask(id)
}
