// Assuming the storage package has a function LoadTasks for listing tasks
package cmd

import (
	"fmt"
	"tasker/storage"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks in the task list.`,
	Run: func(cmd *cobra.Command, args []string) {

		jstorage := storage.NewJSONStorage()

		tasks, err := jstorage.LoadTasks()
		if err != nil {
			fmt.Printf("Error loading tasks: %s\n", err)
			return
		}

		for _, task := range tasks {
			fmt.Println(task.String()) // Assuming Task has a String method for display
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
