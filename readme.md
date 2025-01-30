# Todo List CLI Documentation

## Overview
This project is a command-line interface (CLI) application for managing a todo list. It includes functionalities such as adding, listing, and removing tasks. The code is written in Go and adheres to standard Go project structure and practices.

## Features
- **Add Tasks**: Add new tasks to your todo list.
- **List Tasks**: List all tasks, or filter tasks by their status (done, pending) or by ID.
- **Remove Tasks**: Remove tasks from your todo list.
- **Save and Load**: Save your todo list to a JSON file and load it back.

## Project Structure
This project is structured as follows:
todo-list/
├── cmd/
│   ├── root.go # Main entry point of the application
│   ├── addTask.go # Command for adding tasks
│   ├── listTasks.go # Command for listing tasks
│   └── removeTask.go # Command for removing tasks
├── internal/
│   ├── models/
│   │   └── task.go # Definition of the Task struct
│   ├── todo/
│   │   ├── add.go # Logic for adding tasks
│   │   ├── list.go # Logic for listing tasks
│   │   ├── remove.go # Logic for removing tasks
│   │   └── save.go # Logic for saving and loading tasks
├── tests/
│   ├── add_test.go # Unit tests for adding tasks
│   ├── list_test.go # Unit tests for listing tasks
│   └── remove_test.go # Unit tests for removing tasks
└── readme.md # This file


## Installation
To install the necessary dependencies, run:
```sh
go mod tidy
```

## How to Run
To run the application, use the following command:
```sh
go run cmd/root.go
```

## Contributing
Contributions are welcome! Please fork this repository and submit a pull request for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Contact
For any inquiries or feedback, please contact gbotareli@gmail.com.

