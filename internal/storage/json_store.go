package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"task-tracker-cli/internal/model"
)

const fileName = "tasks.json"

// ensureFileExists checks if tasks.json exists.
// If not, it creates the file with an empty JSON array.
func ensureFileExists() error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		emptyData := []byte("[]")
		return os.WriteFile(fileName, emptyData, 0644)
	}
	return nil
}

func LoadTasks() ([]model.Task, error) {
	err := ensureFileExists()
	if err != nil {
		return nil, err
	}
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	// Handle empty file edge case
	if len(data) == 0 {
		return []model.Task{}, nil
	}

	var tasks []model.Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, errors.New("failed to parse tasks.json (possibly corrupted)")
	}

	return tasks, nil

}

func SaveTasks(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
