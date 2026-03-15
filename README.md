# Task Tracker CLI

A simple command-line task tracker written in Go. Tasks are stored in a local `tasks.json` file and can be added, updated, deleted, and filtered by status.

## ✅ Features

- Add new tasks
- Update task descriptions
- Delete tasks
- Mark tasks as **in-progress** or **done**
- List tasks (optionally filtered by status)
- Stores data in a lightweight `tasks.json` file in the project directory

## 🔧 Prerequisites

- Go 1.20+ installed (or compatible Go toolchain)

## 🚀 Build & Run

From the project root:

```bash
# Build executable
go build -o task-cli .

# Run directly with `go run`
go run main.go <command> [args...]
```

## 📦 Usage

The CLI supports the following commands:

- `add <description>` — Create a new task
- `update <id> <description>` — Update the description of a task
- `delete <id>` — Delete a task by ID
- `mark-in-progress <id>` — Mark a task as in-progress
- `mark-done <id>` — Mark a task as done
- `list [status]` — List all tasks (optional status filter: `todo`, `in-progress`, `done`)

### Examples

```bash
# Add a task
./task-cli add "Buy groceries"

# Update a task
./task-cli update 1 "Buy groceries and snacks"

# Mark as in-progress
./task-cli mark-in-progress 1

# Mark as done
./task-cli mark-done 1

# List all tasks
./task-cli list

# List only done tasks
./task-cli list done
```

## 🗂️ Data Storage

The application uses `tasks.json` in the working directory to persist tasks. If the file does not exist, it is created automatically.

## 🧩 Project Structure

- `main.go` — CLI entry point and argument parsing
- `internal/task/service.go` — Task logic (add/update/delete/status/list)
- `internal/storage/json_store.go` — Task persistence with `tasks.json`
- `internal/model/task.go` — Task model and status constants

---

If you want to extend the CLI, consider adding support for:

- task priorities or due dates
- search/filter by keyword
- interactive mode
- colored output
