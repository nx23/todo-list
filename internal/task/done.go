package task

import "github.com/nx23/todo-list/internal/models"

func Done(todoList *[]models.Task, taskID int) {
	(*todoList)[taskID - 1].Done = true
}