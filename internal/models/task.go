package models

type Task struct {
	Id       int    `json:"id"`
	Task     string `json:"task"`
	Done     bool   `json:"done"`
	CreateAt string `json:"createAt"`
	DoneAt   string `json:"doneAt"`
}
