package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/GourangaDasSamrat/todo-cli-go/internal/models"
	"gopkg.in/yaml.v3"
)

const (
	defaultDataDir  = ".todo-cli"
	defaultDataFile = "tasks.json"
	backupDir       = "backups"
)

// Storage interface for task persistence
type Storage interface {
	Load() (*models.TaskList, error)
	Save(taskList *models.TaskList) error
	Backup() error
	Restore(backupFile string) error
	ListBackups() ([]string, error)
}

// JSONStorage implements Storage using JSON files
type JSONStorage struct {
	dataPath   string
	backupPath string
}

// NewJSONStorage creates a new JSON storage instance
func NewJSONStorage() (*JSONStorage, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	dataDir := filepath.Join(homeDir, defaultDataDir)
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	backupPath := filepath.Join(dataDir, backupDir)
	if err := os.MkdirAll(backupPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create backup directory: %w", err)
	}

	return &JSONStorage{
		dataPath:   filepath.Join(dataDir, defaultDataFile),
		backupPath: backupPath,
	}, nil
}

// Load loads tasks from JSON file
func (s *JSONStorage) Load() (*models.TaskList, error) {
	if _, err := os.Stat(s.dataPath); os.IsNotExist(err) {
		return &models.TaskList{Tasks: []*models.Task{}}, nil
	}

	data, err := os.ReadFile(s.dataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read data file: %w", err)
	}

	var taskList models.TaskList
	if err := json.Unmarshal(data, &taskList); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	taskList.UpdateAllStatuses()
	return &taskList, nil
}

// Save saves tasks to JSON file
func (s *JSONStorage) Save(taskList *models.TaskList) error {
	taskList.UpdateAllStatuses()

	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(s.dataPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write data file: %w", err)
	}

	return nil
}

// Backup creates a backup of current tasks
func (s *JSONStorage) Backup() error {
	if _, err := os.Stat(s.dataPath); os.IsNotExist(err) {
		return fmt.Errorf("no data file to backup")
	}

	data, err := os.ReadFile(s.dataPath)
	if err != nil {
		return fmt.Errorf("failed to read data file: %w", err)
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	backupFile := filepath.Join(s.backupPath, fmt.Sprintf("tasks_backup_%s.json", timestamp))

	if err := os.WriteFile(backupFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write backup file: %w", err)
	}

	return nil
}

// Restore restores tasks from a backup file
func (s *JSONStorage) Restore(backupFile string) error {
	backupPath := filepath.Join(s.backupPath, backupFile)
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		return fmt.Errorf("backup file not found: %s", backupFile)
	}

	data, err := os.ReadFile(backupPath)
	if err != nil {
		return fmt.Errorf("failed to read backup file: %w", err)
	}

	if err := os.WriteFile(s.dataPath, data, 0644); err != nil {
		return fmt.Errorf("failed to restore backup: %w", err)
	}

	return nil
}

// ListBackups returns a list of available backup files
func (s *JSONStorage) ListBackups() ([]string, error) {
	files, err := os.ReadDir(s.backupPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read backup directory: %w", err)
	}

	var backups []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			backups = append(backups, file.Name())
		}
	}

	return backups, nil
}

// YAMLStorage implements Storage using YAML files
type YAMLStorage struct {
	dataPath   string
	backupPath string
}

// NewYAMLStorage creates a new YAML storage instance
func NewYAMLStorage() (*YAMLStorage, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	dataDir := filepath.Join(homeDir, defaultDataDir)
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	backupPath := filepath.Join(dataDir, backupDir)
	if err := os.MkdirAll(backupPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create backup directory: %w", err)
	}

	return &YAMLStorage{
		dataPath:   filepath.Join(dataDir, "tasks.yaml"),
		backupPath: backupPath,
	}, nil
}

// Load loads tasks from YAML file
func (s *YAMLStorage) Load() (*models.TaskList, error) {
	if _, err := os.Stat(s.dataPath); os.IsNotExist(err) {
		return &models.TaskList{Tasks: []*models.Task{}}, nil
	}

	data, err := os.ReadFile(s.dataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read data file: %w", err)
	}

	var taskList models.TaskList
	if err := yaml.Unmarshal(data, &taskList); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	taskList.UpdateAllStatuses()
	return &taskList, nil
}

// Save saves tasks to YAML file
func (s *YAMLStorage) Save(taskList *models.TaskList) error {
	taskList.UpdateAllStatuses()

	data, err := yaml.Marshal(taskList)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %w", err)
	}

	if err := os.WriteFile(s.dataPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write data file: %w", err)
	}

	return nil
}

// Backup creates a backup of current tasks
func (s *YAMLStorage) Backup() error {
	if _, err := os.Stat(s.dataPath); os.IsNotExist(err) {
		return fmt.Errorf("no data file to backup")
	}

	data, err := os.ReadFile(s.dataPath)
	if err != nil {
		return fmt.Errorf("failed to read data file: %w", err)
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	backupFile := filepath.Join(s.backupPath, fmt.Sprintf("tasks_backup_%s.yaml", timestamp))

	if err := os.WriteFile(backupFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write backup file: %w", err)
	}

	return nil
}

// Restore restores tasks from a backup file
func (s *YAMLStorage) Restore(backupFile string) error {
	backupPath := filepath.Join(s.backupPath, backupFile)
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		return fmt.Errorf("backup file not found: %s", backupFile)
	}

	data, err := os.ReadFile(backupPath)
	if err != nil {
		return fmt.Errorf("failed to read backup file: %w", err)
	}

	if err := os.WriteFile(s.dataPath, data, 0644); err != nil {
		return fmt.Errorf("failed to restore backup: %w", err)
	}

	return nil
}

// ListBackups returns a list of available backup files
func (s *YAMLStorage) ListBackups() ([]string, error) {
	files, err := os.ReadDir(s.backupPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read backup directory: %w", err)
	}

	var backups []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".yaml" {
			backups = append(backups, file.Name())
		}
	}

	return backups, nil
}
