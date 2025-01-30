package task

import (
	"fmt"

	"github.com/nx23/todo-list/internal/models"
)

func Show(task models.Task) {
	fmt.Printf("ID: %d\tTask: %s\tDone? %t\tCreate At: %s\tDone At: %s\n", task.Id, task.Task, task.Done, task.CreateAt, task.DoneAt)
}