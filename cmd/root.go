package cmd

import (
	"fmt"
	"os"

	"github.com/GourangaDasSamrat/todo-cli-go/internal/storage"
	"github.com/GourangaDasSamrat/todo-cli-go/internal/ui"
	"github.com/spf13/cobra"
)

var (
	storageType string
	store       storage.Storage
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A powerful CLI todo application",
	Long: `Todo CLI - A feature-rich command-line todo application

Complete with task management, filtering, sorting, and interactive mode.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		if storageType == "yaml" {
			store, err = storage.NewYAMLStorage()
		} else {
			store, err = storage.NewJSONStorage()
		}
		if err != nil {
			ui.PrintError(fmt.Sprintf("Failed to initialize storage: %v", err))
			os.Exit(1)
		}
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ui.PrintError(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&storageType, "storage", "s", "json", "Storage type (json or yaml)")
}
