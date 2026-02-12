package cmd

import (
	"fmt"
	"time"

	"github.com/GourangaDasSamrat/todo-cli-go/internal/models"
	"github.com/GourangaDasSamrat/todo-cli-go/internal/ui"
	"github.com/GourangaDasSamrat/todo-cli-go/internal/utils"
	"github.com/GourangaDasSamrat/todo-cli-go/pkg/filter"
	sortpkg "github.com/GourangaDasSamrat/todo-cli-go/pkg/sort"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:     "interactive",
	Aliases: []string{"i", "int"},
	Short:   "Start interactive mode",
	Long:    `Start interactive mode with menu-driven interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		runInteractiveMode()
	},
}

func runInteractiveMode() {
	menu := ui.NewInteractiveMenu()

	for {
		choice, err := menu.Show()
		if err != nil {
			ui.PrintError("Menu error: " + err.Error())
			return
		}

		switch choice {
		case 0: // View All Tasks
			viewAllTasks()
		case 1: // Add New Task
			addTaskInteractive()
		case 2: // Edit Task
			editTaskInteractive()
		case 3: // Delete Task
			deleteTaskInteractive()
		case 4: // Mark Complete/Incomplete
			toggleCompleteInteractive()
		case 5: // Filter Tasks
			filterTasksInteractive()
		case 6: // Search Tasks
			searchTasksInteractive()
		case 7: // Backup Data
			backupInteractive()
		case 8: // Restore Data
			restoreInteractive()
		case 9: // Exit
			ui.PrintInfo("Goodbye!")
			return
		}

		fmt.Println() // Add spacing between operations
	}
}

func viewAllTasks() {
	taskList, err := store.Load()
	if err != nil {
		ui.PrintError("Failed to load tasks: " + err.Error())
		return
	}

	ui.PrintHeader("All Tasks")
	ui.PrintTaskList(taskList.Tasks)
}

func addTaskInteractive() {
	ui.PrintHeader("Add New Task")

	task, err := ui.PromptTaskInput()
	if err != nil {
		ui.PrintError("Input cancelled or failed: " + err.Error())
		return
	}

	task.ID = utils.GenerateID()
	task.Status = models.StatusPending
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	taskList, err := store.Load()
	if err != nil {
		ui.PrintError("Failed to load tasks: " + err.Error())
		return
	}

	taskList.Add(task)

	if err := store.Save(taskList); err != nil {
		ui.PrintError("Failed to save task: " + err.Error())
		return
	}

	ui.PrintSuccess("Task added successfully!")
	ui.PrintTask(task)
}

func editTaskInteractive() {
	taskList, err := store.Load()
	if err != nil {
		ui.PrintError("Failed to load tasks: " + err.Error())
		return
	}

	if len(taskList.Tasks) == 0 {
		ui.PrintInfo("No tasks available to edit.")
		return
	}

	ui.PrintHeader("Edit Task")
	task, err := ui.SelectTask(taskList.Tasks)
	if err != nil {
		ui.PrintError("Selection cancelled: " + err.Error())
		return
	}

	// Show current task
	ui.PrintInfo("Current task details:")
	ui.PrintTask(task)

	// Prompt for updates
	newTask, err := ui.PromptTaskInput()
	if err != nil {
		ui.PrintError("Input cancelled: " + err.Error())
		return
	}

	// Update task
	task.Title = newTask.Title
	task.Description = newTask.Description
	task.Priority = newTask.Priority
	task.Project = newTask.Project
	task.Tags = newTask.Tags
	if !newTask.DueDate.IsZero() {
		task.DueDate = newTask.DueDate
	}
	task.UpdatedAt = time.Now()
	task.UpdateStatus()

	if err := store.Save(taskList); err != nil {
		ui.PrintError("Failed to save task: " + err.Error())
		return
	}

	ui.PrintSuccess("Task updated successfully!")
	ui.PrintTask(task)
}

func deleteTaskInteractive() {
	taskList, err := store.Load()
	if err != nil {
		ui.PrintError("Failed to load tasks: " + err.Error())
		return
	}

	if len(taskList.Tasks) == 0 {
		ui.PrintInfo("No tasks available to delete.")
		return
	}

	ui.PrintHeader("Delete Task")
	task, err := ui.SelectTask(taskList.Tasks)
	if err != nil {
		ui.PrintError("Selection cancelled: " + err.Error())
		return
	}

	ui.PrintWarning("About to delete:")
	ui.PrintTask(task)

	if !ui.ConfirmAction("Delete this task?") {
		ui.PrintInfo("Deletion cancelled.")
		return
	}

	taskList.Remove(task.ID)

	if err := store.Save(taskList); err != nil {
		ui.PrintError("Failed to save changes: " + err.Error())
		return
	}

	ui.PrintSuccess("Task deleted successfully!")
}

func toggleCompleteInteractive() {
	taskList, err := store.Load()
	if err != nil {
		ui.PrintError("Failed to load tasks: " + err.Error())
		return
	}

	if len(taskList.Tasks) == 0 {
		ui.PrintInfo("No tasks available.")
		return
	}

	ui.PrintHeader("Toggle Task Completion")
	task, err := ui.SelectTask(taskList.Tasks)
	if err != nil {
		ui.PrintError("Selection cancelled: " + err.Error())
		return
	}

	if task.Status == models.StatusCompleted {
		task.MarkIncomplete()
		ui.PrintSuccess("Task marked as incomplete!")
	} else {
		task.MarkComplete()
		ui.PrintSuccess("Task marked as complete!")
	}

	if err := store.Save(taskList); err != nil {
		ui.PrintError("Failed to save task: " + err.Error())
		return
	}

	ui.PrintTask(task)
}

func filterTasksInteractive() {
	taskList, err := store.Load()
	if err != nil {
		ui.PrintError("Failed to load tasks: " + err.Error())
		return
	}

	ui.PrintHeader("Filter Tasks")

	filterType, err := ui.SelectOption("Filter by", []string{
		"Status",
		"Priority",
		"Project",
		"Cancel",
	})
	if err != nil || filterType == "Cancel" {
		return
	}

	var f *filter.Filter

	switch filterType {
	case "Status":
		status, err := ui.SelectOption("Select status", []string{"pending", "completed", "overdue"})
		if err != nil {
			return
		}
		var s models.Status
		switch status {
		case "pending":
			s = models.StatusPending
		case "completed":
			s = models.StatusCompleted
		case "overdue":
			s = models.StatusOverdue
		}
		f = filter.NewStatusFilter(s)

	case "Priority":
		priority, err := ui.SelectOption("Select priority", []string{"low", "medium", "high"})
		if err != nil {
			return
		}
		p := models.ParsePriority(priority)
		f = filter.NewPriorityFilter(p)

	case "Project":
		project, err := ui.PromptInput("Enter project name", true)
		if err != nil {
			return
		}
		f = filter.NewProjectFilter(project)
	}

	tasks := f.Apply(taskList.Tasks)
	sortpkg.Sort(tasks, sortpkg.SortByPriority, false)

	ui.PrintHeader("Filtered Results")
	ui.PrintTaskList(tasks)
}

func searchTasksInteractive() {
	taskList, err := store.Load()
	if err != nil {
		ui.PrintError("Failed to load tasks: " + err.Error())
		return
	}

	ui.PrintHeader("Search Tasks")
	keyword, err := ui.PromptInput("Enter search keyword", true)
	if err != nil {
		return
	}

	f := filter.NewKeywordFilter(keyword)
	tasks := f.Apply(taskList.Tasks)
	sortpkg.Sort(tasks, sortpkg.SortByCreatedAt, false)

	ui.PrintHeader("Search Results")
	ui.PrintTaskList(tasks)
}

func backupInteractive() {
	ui.PrintHeader("Create Backup")

	if err := store.Backup(); err != nil {
		ui.PrintError("Failed to create backup: " + err.Error())
		return
	}

	ui.PrintSuccess("Backup created successfully!")
}

func restoreInteractive() {
	ui.PrintHeader("Restore from Backup")

	backups, err := store.ListBackups()
	if err != nil {
		ui.PrintError("Failed to list backups: " + err.Error())
		return
	}

	if len(backups) == 0 {
		ui.PrintInfo("No backups found.")
		return
	}

	backup, err := ui.SelectOption("Select backup to restore", backups)
	if err != nil {
		return
	}

	ui.PrintWarning("This will replace all current tasks!")
	if !ui.ConfirmAction("Restore from backup?") {
		ui.PrintInfo("Restore cancelled.")
		return
	}

	if err := store.Restore(backup); err != nil {
		ui.PrintError("Failed to restore backup: " + err.Error())
		return
	}

	ui.PrintSuccess("Backup restored successfully!")
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}
