package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/nx23/todo-list/internal/models"
)

func Load() ([]models.Task, error) {
	var tasks []byte
	var err error

	tasks, err = os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Failed to find tasks.json file, creating a new one")
		file, err := os.Create("tasks.json")
		if err != nil {
			return nil, fmt.Errorf("failed to create tasks.json file: %w", err)
		}

		defer file.Close()
		
		_, err = file.Write([]byte("[]"))
		if err != nil {
			return nil, fmt.Errorf("failed to write to tasks.json file: %w", err)
		}

		tasks, _ = os.ReadFile("tasks.json")
	}

	var todo []models.Task
	if err := json.NewDecoder(bytes.NewBuffer(tasks)).Decode(&todo); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return todo, nil
}