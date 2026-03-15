package task

import (
	"errors"
	"task-tracker-cli/internal/model"
	"task-tracker-cli/internal/storage"
	"time"
)

func Addtask(description string) (int, error) {

	if description == "" {
		return 0, errors.New("Description cannot be empty")
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		return 0, err
	}
	newID := generateNextID(tasks)
	now := time.Now()

	newTask := model.Task{
		ID:          newID,
		Description: description,
		Status:      model.StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, newTask)
	err = storage.SaveTasks(tasks)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func UpdateTask(id int, description string) error {

	if description == "" {
		return errors.New("Description cannot be empty")
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return storage.SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

func Deletetask(id int) error {

	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return storage.SaveTasks(tasks)
		}
	}
	return errors.New("task not found")
}

func MarkInProgress(id int) error {
	return updateStatus(id, model.StatusInProgress)
}

// MarkDone sets task status to done.
func MarkDone(id int) error {
	return updateStatus(id, model.StatusDone)
}
func updateStatus(id int, newStatus string) error {

	tasks, err := storage.LoadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now()
			return storage.SaveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

func generateNextID(tasks []model.Task) int {
	maxID := 0

	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	return maxID + 1
}

func ListTasks(status string) ([]model.Task, error) {

	tasks, err := storage.LoadTasks()
	if err != nil {
		return nil, err
	}

	if status == "" {
		return tasks, nil
	}

	if !isValidStatus(status) {
		return nil, errors.New("invalid status filter")
	}

	var filtered []model.Task

	for i := range tasks {
		if tasks[i].Status == status {
			filtered = append(filtered, tasks[i])
		}
	}

	return filtered, nil
}

func isValidStatus(status string) bool {
	return status == model.StatusTodo ||
		status == model.StatusInProgress ||
		status == model.StatusDone
}
