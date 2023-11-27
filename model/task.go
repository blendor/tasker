package model

import (
	"fmt"
	"time"
)

// Task represents a task with a description, priority, urgency, importance, and creation time.
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`   // This field can be kept or removed depending on how priority is determined
	Urgency     string    `json:"urgency"`    // New field for urgency ("u" for urgent, "nu" for non-urgent)
	Importance  string    `json:"importance"` // New field for importance ("i" for important, "ni" for not important)
	CreatedAt   time.Time `json:"created_at"`
}

func (t *Task) String() string {
	return fmt.Sprintf("%d. \"%s\"", t.ID, t.Description)
}

// NewTask creates and returns a new Task with urgency and importance.
func NewTask(id int, description string, priority int, urgency string, importance string) Task {
	return Task{
		ID:          id,
		Description: description,
		Priority:    priority, // The priority field can be calculated based on urgency and importance if needed
		Urgency:     urgency,
		Importance:  importance,
		CreatedAt:   time.Now(),
	}
}
