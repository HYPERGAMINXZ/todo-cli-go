package cmd

import (
	"time"

	"github.com/GourangaDasSamrat/todo-cli-go/internal/models"
	"github.com/GourangaDasSamrat/todo-cli-go/internal/ui"
	"github.com/GourangaDasSamrat/todo-cli-go/internal/utils"
	"github.com/spf13/cobra"
)

var (
	editID          string
	editTitle       string
	editDescription string
	editPriority    string
	editProject     string
	editTags        []string
	editDueDate     string
)

var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e", "update", "modify"},
	Short:   "Edit an existing task",
	Long:    `Edit an existing task by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load tasks
		taskList, err := store.Load()
		if err != nil {
			ui.PrintError("Failed to load tasks: " + err.Error())
			return
		}

		// Find task
		task := taskList.GetByID(editID)
		if task == nil {
			ui.PrintError("Task not found with ID: " + editID)
			return
		}

		// Update fields if provided
		if cmd.Flags().Changed("title") {
			task.Title = editTitle
		}
		if cmd.Flags().Changed("description") {
			task.Description = editDescription
		}
		if cmd.Flags().Changed("priority") {
			task.Priority = models.ParsePriority(editPriority)
		}
		if cmd.Flags().Changed("project") {
			task.Project = editProject
		}
		if cmd.Flags().Changed("tags") {
			task.Tags = editTags
		}
		if cmd.Flags().Changed("due") {
			if editDueDate == "" {
				task.DueDate = time.Time{}
			} else {
				dueDate, err := utils.ParseDate(editDueDate)
				if err != nil {
					ui.PrintError("Invalid due date format: " + err.Error())
					return
				}
				task.DueDate = dueDate
			}
		}

		task.UpdatedAt = time.Now()
		task.UpdateStatus()

		// Save
		if err := store.Save(taskList); err != nil {
			ui.PrintError("Failed to save task: " + err.Error())
			return
		}

		ui.PrintSuccess("Task updated successfully!")
		ui.PrintTask(task)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringVarP(&editID, "id", "i", "", "Task ID (required)")
	editCmd.Flags().StringVarP(&editTitle, "title", "t", "", "New task title")
	editCmd.Flags().StringVarP(&editDescription, "description", "d", "", "New task description")
	editCmd.Flags().StringVarP(&editPriority, "priority", "p", "", "New priority (low, medium, high)")
	editCmd.Flags().StringVar(&editProject, "project", "", "New project name")
	editCmd.Flags().StringSliceVar(&editTags, "tags", []string{}, "New tags")
	editCmd.Flags().StringVar(&editDueDate, "due", "", "New due date (empty to clear)")

	editCmd.MarkFlagRequired("id")
}
