package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasker",
	Short: "Tasker is a task management CLI tool",
	Long:  `Tasker is a CLI tool designed to help manage and prioritize tasks using the Eisenhower matrix.`,
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// This can include global settings or flags that are applicable to all subcommands
}
