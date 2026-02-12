package cmd

import (
	"github.com/GourangaDasSamrat/todo-cli-go/internal/models"
	"github.com/GourangaDasSamrat/todo-cli-go/internal/ui"
	"github.com/GourangaDasSamrat/todo-cli-go/pkg/filter"
	sortpkg "github.com/GourangaDasSamrat/todo-cli-go/pkg/sort"
	"github.com/spf13/cobra"
)

var (
	listStatus   string
	listPriority string
	listProject  string
	listTags     []string
	listSortBy   string
	listAsc      bool
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "List all tasks",
	Long:    `Display all tasks with optional filtering and sorting.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load tasks
		taskList, err := store.Load()
		if err != nil {
			ui.PrintError("Failed to load tasks: " + err.Error())
			return
		}

		tasks := taskList.Tasks

		// Apply filters
		f := &filter.Filter{}
		if listStatus != "" {
			switch listStatus {
			case "pending":
				status := models.StatusPending
				f.Status = &status
			case "completed":
				status := models.StatusCompleted
				f.Status = &status
			case "overdue":
				status := models.StatusOverdue
				f.Status = &status
			}
		}

		if listPriority != "" {
			priority := models.ParsePriority(listPriority)
			f.Priority = &priority
		}

		if listProject != "" {
			f.Project = listProject
		}

		if len(listTags) > 0 {
			f.Tags = listTags
		}

		tasks = f.Apply(tasks)

		// Apply sorting
		switch listSortBy {
		case "priority":
			sortpkg.Sort(tasks, sortpkg.SortByPriority, listAsc)
		case "date", "due":
			sortpkg.Sort(tasks, sortpkg.SortByDueDate, listAsc)
		case "created":
			sortpkg.Sort(tasks, sortpkg.SortByCreatedAt, listAsc)
		case "title":
			sortpkg.Sort(tasks, sortpkg.SortByTitle, listAsc)
		default:
			sortpkg.Sort(tasks, sortpkg.SortByCreatedAt, false)
		}

		// Display tasks
		ui.PrintTaskList(tasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVar(&listStatus, "status", "", "Filter by status (pending, completed, overdue)")
	listCmd.Flags().StringVar(&listPriority, "priority", "", "Filter by priority (low, medium, high)")
	listCmd.Flags().StringVar(&listProject, "project", "", "Filter by project")
	listCmd.Flags().StringSliceVar(&listTags, "tags", []string{}, "Filter by tags")
	listCmd.Flags().StringVar(&listSortBy, "sort", "created", "Sort by (priority, date, created, title)")
	listCmd.Flags().BoolVar(&listAsc, "asc", false, "Sort in ascending order")
}
