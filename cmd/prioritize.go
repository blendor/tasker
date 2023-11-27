package cmd

import (
	"fmt"
	"strconv"
	"tasker/storage" // Importing the storage package

	"os"

	"github.com/spf13/cobra"
)

// prioritizeCmd represents the prioritize command
var prioritizeCmd = &cobra.Command{
	Use:   "prioritize [task ID] [u/nu] [i/ni]",
	Short: "Prioritize a task",
	Long:  `Sets or updates the priority of a given task based on urgency (u/nu) and importance (i/ni).`,
	Run: func(cmd *cobra.Command, args []string) {
		// Validate input
		if len(args) < 3 {
			fmt.Println("Error: Task ID, urgency (u/nu), and importance (i/ni) are required")
			return
		}
		taskIDStr, urgency, importance := args[0], args[1], args[2]

		// Convert taskID from string to int
		taskID, err := strconv.Atoi(taskIDStr)
		if err != nil {
			fmt.Println("Error: Invalid task ID")
			return
		}

		// Validate urgency and importance inputs
		if !(urgency == "u" || urgency == "nu") || !(importance == "i" || importance == "ni") {
			fmt.Println("Error: Invalid urgency or importance values")
			return
		}

		// Implement the prioritize functionality
		err = PrioritizeTask(taskID, urgency, importance)
		if err != nil {
			fmt.Printf("Error prioritizing task: %s\n", err)
			return
		}

		fmt.Fprintf(os.Stdout, "Task successfully prioritized: %s\n", []any{taskID}...)
	},
}

func init() {
	rootCmd.AddCommand(prioritizeCmd)
}

// PrioritizeTask function implemented using the storage package
func PrioritizeTask(taskID int, urgency, importance string) error {
	// Get the task, update its priority based on urgency and importance, and save it back

	jstorage := storage.NewJSONStorage()

	task, err := jstorage.GetTask(taskID)
	if err != nil {
		return fmt.Errorf("error getting task: %s", err)
	}

	// Update the task's priority based on urgency and importance
	task.Urgency = urgency
	task.Importance = importance

	// Save the updated task back to storage

	err = jstorage.UpdateTask(task)
	if err != nil {
		return fmt.Errorf("error updating task: %s", err)
	}

	return nil
}
