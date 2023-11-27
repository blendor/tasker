// Assuming a struct Task is defined in the model package and AddTask function in the storage package
package cmd

import (
	"fmt"
	"strings"
	"tasker/model"
	"tasker/storage"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task",
	Long:  `Add a new task to the task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: Task description is required")
			return
		}
		taskDescription := strings.Join(args, " ")

		jstorage := storage.NewJSONStorage()

		// Create a new task
		newTask := model.Task{Description: taskDescription}
		err := jstorage.SaveTask(newTask)
		if err != nil {
			fmt.Printf("Error adding task: %s\n", err)
			return
		}

		fmt.Println("Task added successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
