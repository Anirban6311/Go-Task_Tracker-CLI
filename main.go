package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker-cli/internal/task"
)

func main() {
	// Ensure at least one argument is passed
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {

	case "add":
		handleAdd()
	case "update":
		handleUpdate()
	case "delete":
		handleDelete()
	case "mark-in-progress":
		handleMarkInProgress()
	case "mark-done":
		handleMarkDone()
	case "list":
		handleList()
	default:
		fmt.Println("Unknown command:", command)
		printUsage()

	}
}

// / 0 1 2
// / ./task-cli add "Buy groceries"
func handleAdd() {
	if len(os.Args) < 3 {
		fmt.Printf("Missing task description")
	}

	description := os.Args[2]
	id, err := task.Addtask(description)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Task added successfully (ID : %d)", id)

}

// / 0 1 2 3
// / 1-> ID
// / ./task-cli update 1 "Buy groceries"
func handleUpdate() {
	if len(os.Args) < 4 {
		fmt.Println("Error: Missing arguments. Usage: update <id> <description>")
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid task id:", err)
		return
	}
	description := os.Args[3]
	err = task.UpdateTask(id, description)
	if err != nil {
		fmt.Println("Error while updating task", err)
		return
	}
	fmt.Println("Task updated successfully", id)

}
func handleDelete() {
	if len(os.Args) < 3 {
		fmt.Printf("Missing task ID")
	}

	taskId, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid taskId:", err)
		return
	}
	err = task.Deletetask(taskId)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println("Task deleted successfully")

}

func handleMarkInProgress() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Missing task ID.")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: Invalid task ID.")
		return
	}

	err = task.MarkInProgress(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Task marked as in-progress.")

}

func handleMarkDone() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Missing task ID.")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: Invalid task ID.")
		return
	}

	err = task.MarkDone(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Task marked as done.")
}

func handleList() {
	var status string

	if len(os.Args) >= 3 {
		status = os.Args[2]
	}

	tasks, err := task.ListTasks(status)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, t := range tasks {
		fmt.Printf("ID: %d | %s | %s | Created: %s | Updated: %s\n",
			t.ID,
			t.Description,
			t.Status,
			t.CreatedAt.Format("2006-01-02 15:04:05"),
			t.UpdatedAt.Format("2006-01-02 15:04:05"),
		)
	}
}

func printUsage() {
	fmt.Println("Task Tracker CLI Usage:")
	fmt.Println("  add <description>")
	fmt.Println("  update <id> <description>")
	fmt.Println("  delete <id>")
	fmt.Println("  mark-in-progress <id>")
	fmt.Println("  mark-done <id>")
	fmt.Println("  list")
	fmt.Println("  list <status>")
}
