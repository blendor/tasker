// Assuming the storage package has a function DeleteTask for deleting a task
package cmd

import (
	"fmt"
	"strconv"
	"tasker/storage"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task",
	Long:  `Delete a task from the task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: Task ID is required")
			return
		}
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: Invalid task ID")
			return
		}

		jstorage := storage.NewJSONStorage()

		err = jstorage.DeleteTask(taskID)
		if err != nil {
			fmt.Printf("Error deleting task: %s\n", err)
			return
		}

		fmt.Println("Task deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
