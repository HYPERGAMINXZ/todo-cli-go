package cmd

import (
	"time"

	"github.com/GourangaDasSamrat/todo-cli-go/internal/models"
	"github.com/GourangaDasSamrat/todo-cli-go/internal/ui"
	"github.com/GourangaDasSamrat/todo-cli-go/internal/utils"
	"github.com/spf13/cobra"
)

var (
	addTitle       string
	addDescription string
	addPriority    string
	addProject     string
	addTags        []string
	addDueDate     string
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a", "new", "create"},
	Short:   "Add a new task",
	Long:    `Add a new task to your todo list with optional metadata.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load existing tasks
		taskList, err := store.Load()
		if err != nil {
			ui.PrintError("Failed to load tasks: " + err.Error())
			return
		}

		// Create new task
		task := &models.Task{
			ID:          utils.GenerateID(),
			Title:       addTitle,
			Description: addDescription,
			Priority:    models.ParsePriority(addPriority),
			Status:      models.StatusPending,
			Project:     addProject,
			Tags:        addTags,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// Parse due date if provided
		if addDueDate != "" {
			dueDate, err := utils.ParseDate(addDueDate)
			if err != nil {
				ui.PrintError("Invalid due date format: " + err.Error())
				return
			}
			task.DueDate = dueDate
		}

		// Add task to list
		taskList.Add(task)

		// Save
		if err := store.Save(taskList); err != nil {
			ui.PrintError("Failed to save task: " + err.Error())
			return
		}

		ui.PrintSuccess("Task added successfully!")
		ui.PrintTask(task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addTitle, "title", "t", "", "Task title (required)")
	addCmd.Flags().StringVarP(&addDescription, "description", "d", "", "Task description")
	addCmd.Flags().StringVarP(&addPriority, "priority", "p", "low", "Task priority (low, medium, high)")
	addCmd.Flags().StringVar(&addProject, "project", "", "Project name")
	addCmd.Flags().StringSliceVar(&addTags, "tags", []string{}, "Task tags (comma-separated)")
	addCmd.Flags().StringVar(&addDueDate, "due", "", "Due date (YYYY-MM-DD or YYYY-MM-DD HH:MM)")

	addCmd.MarkFlagRequired("title")
}
