package todo

import (
	"fmt"

	"github.com/nx23/todo-list/internal/models"
	"github.com/nx23/todo-list/internal/task"
)

type FilterParameters struct {
	Id int
	Done bool
	All bool
}

func List(todoList *[]models.Task, filters FilterParameters) {
	fmt.Println("Your's Tasks are: ")
	
	for _, currentTask := range *todoList {
		if filters.All {
			task.Show(currentTask)
		} else if filters.Id == currentTask.Id {
			task.Show(currentTask)
		} else if filters.Done == currentTask.Done {
			task.Show(currentTask)
		}
	}
}