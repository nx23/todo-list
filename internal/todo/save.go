package todo

import (
	"encoding/json"
	"log"
	"os"

	"github.com/nx23/todo-list/internal/models"
)

func Save(todoList *[]models.Task) {
	file, err := os.OpenFile("tasks.json", os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0777)

	if err != nil {
		log.Fatalf("error while opening the file. %v", err)
	}

	defer file.Close()

	task, err := json.Marshal(todoList)
	if err != nil {
		log.Fatalf("error while marshaling the todo. %v", err)
	}

	if _, err := file.Write([]byte(task)); err != nil {
		log.Fatalf("error while writing the file. %v", err)
	}
}