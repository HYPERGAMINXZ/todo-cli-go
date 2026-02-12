package cmd

import (
	"github.com/GourangaDasSamrat/todo-cli-go/internal/ui"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:     "backup",
	Aliases: []string{"bak"},
	Short:   "Create a backup of tasks",
	Long:    `Create a timestamped backup of all tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := store.Backup(); err != nil {
			ui.PrintError("Failed to create backup: " + err.Error())
			return
		}

		ui.PrintSuccess("Backup created successfully!")
	},
}

var restoreCmd = &cobra.Command{
	Use:     "restore [backup-file]",
	Aliases: []string{"rst"},
	Short:   "Restore tasks from a backup",
	Long:    `Restore tasks from a backup file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// List available backups
			backups, err := store.ListBackups()
			if err != nil {
				ui.PrintError("Failed to list backups: " + err.Error())
				return
			}

			if len(backups) == 0 {
				ui.PrintInfo("No backups found.")
				return
			}

			ui.PrintHeader("Available Backups")
			for i, backup := range backups {
				ui.InfoColor.Printf("%d. %s\n", i+1, backup)
			}
			ui.PrintInfo("\nTo restore a backup, run: todo restore <backup-file>")
			return
		}

		backupFile := args[0]

		// Confirm restore
		ui.PrintWarning("This will replace all current tasks with the backup.")
		if !ui.ConfirmAction("Are you sure you want to restore from backup?") {
			ui.PrintInfo("Restore cancelled.")
			return
		}

		if err := store.Restore(backupFile); err != nil {
			ui.PrintError("Failed to restore backup: " + err.Error())
			return
		}

		ui.PrintSuccess("Backup restored successfully!")
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	rootCmd.AddCommand(restoreCmd)
}
