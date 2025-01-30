package task

import (
	"time"

	"github.com/nx23/todo-list/internal/models"
)

type CreateParameters struct {
	Task string
	IsDone bool `default:"false"`
}

func Create(prm CreateParameters) models.Task {
	createAt := time.Now().Format(time.RFC3339)
	doneAt := ""

	if prm.IsDone {
		doneAt = createAt
	}

	return models.Task{
		Task: prm.Task,
		Done: prm.IsDone,
		CreateAt: createAt,
		DoneAt: doneAt,
	}
}