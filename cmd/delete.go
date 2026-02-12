package cmd

import (
	"github.com/GourangaDasSamrat/todo-cli-go/internal/ui"
	"github.com/spf13/cobra"
)

var deleteID string

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "rm", "remove"},
	Short:   "Delete a task",
	Long:    `Delete a task by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load tasks
		taskList, err := store.Load()
		if err != nil {
			ui.PrintError("Failed to load tasks: " + err.Error())
			return
		}

		// Find and display task before deletion
		task := taskList.GetByID(deleteID)
		if task == nil {
			ui.PrintError("Task not found with ID: " + deleteID)
			return
		}

		// Confirm deletion
		ui.PrintWarning("About to delete the following task:")
		ui.PrintTask(task)

		if !ui.ConfirmAction("Are you sure you want to delete this task?") {
			ui.PrintInfo("Deletion cancelled.")
			return
		}

		// Remove task
		if !taskList.Remove(deleteID) {
			ui.PrintError("Failed to remove task")
			return
		}

		// Save
		if err := store.Save(taskList); err != nil {
			ui.PrintError("Failed to save changes: " + err.Error())
			return
		}

		ui.PrintSuccess("Task deleted successfully!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&deleteID, "id", "i", "", "Task ID (required)")
	deleteCmd.MarkFlagRequired("id")
}
