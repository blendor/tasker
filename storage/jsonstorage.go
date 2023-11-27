package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"tasker/model"
)

// JSONStorage implements Storage interface using a JSON file.
type JSONStorage struct {
	FilePath string
}

// NewJSONStorage creates a new instance of JSONStorage.
func NewJSONStorage() *JSONStorage {
	// Assuming the user's desktop path is used to store the JSON file
	desktopPath, _ := os.UserHomeDir()
	filePath := filepath.Join(desktopPath, "Desktop", "tasks.json")
	return &JSONStorage{FilePath: filePath}
}

// SaveTask saves a task to the JSON file.
func (s *JSONStorage) SaveTask(task model.Task) error {
	tasks, err := s.LoadTasks()
	if err != nil {
		return err
	}

	// Assign a new ID to the task if it doesn't have one
	if task.ID == 0 {
		newID := len(tasks) + 1
		task.ID = newID
	}

	// Check if task already exists and update it; otherwise, add as a new task
	updated := false
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			updated = true
			break
		}
	}
	if !updated {
		tasks = append(tasks, task)
	}

	jsonData, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.FilePath, jsonData, 0644)
}

// LoadTasks loads all tasks from the JSON file.
func (s *JSONStorage) LoadTasks() ([]model.Task, error) {
	file, err := ioutil.ReadFile(s.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, return an empty task list
			return []model.Task{}, nil
		}
		return nil, err
	}

	var tasks []model.Task
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

// GetTask retrieves a task by its ID.
func (s *JSONStorage) GetTask(id int) (model.Task, error) {
	tasks, err := s.LoadTasks()
	if err != nil {
		return model.Task{}, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return model.Task{}, fmt.Errorf("task not found")
}

// DeleteTask deletes a task by its ID.
func (s *JSONStorage) DeleteTask(id int) error {
	tasks, err := s.LoadTasks()
	if err != nil {
		return err
	}

	// Remove the task with the specified ID
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			jsonData, err := json.Marshal(tasks)
			if err != nil {
				return err
			}

			return ioutil.WriteFile(s.FilePath, jsonData, 0644) // Corrected: Added closing bracket
		}
	}

	return fmt.Errorf("task not found")
}

// UpdateTask updates a task in the JSON file.
func (s *JSONStorage) UpdateTask(updatedTask model.Task) error {
	tasks, err := s.LoadTasks()
	if err != nil {
		return err
	}

	// Find and update the task
	found := false
	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			tasks[i] = updatedTask
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task not found")
	}

	jsonData, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.FilePath, jsonData, 0644)
}
