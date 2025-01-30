package todo

import (
	"github.com/nx23/todo-list/internal/models"
)

func Reset(todoList *[]models.Task) {
	*todoList = []models.Task{}
}