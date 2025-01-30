package todo

import (
	"github.com/nx23/todo-list/internal/models"
)

func Add(todoList *[]models.Task, newTask models.Task) {
	newTask.Id = len(*todoList) + 1
	*todoList = append(*todoList, newTask)
}