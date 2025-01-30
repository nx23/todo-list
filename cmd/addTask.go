package cmd

import (
	"log"

	"github.com/nx23/todo-list/internal/task"
	"github.com/nx23/todo-list/internal/todo"
	"github.com/spf13/cobra"
)

var addTask = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to the todo list
Example:
todo-list add -t "Buy milk" -d false`,
	Run:   addNewTask,
}

func init() {
	rootCmd.AddCommand(addTask)

	addTask.Flags().StringP("task", "t", "", "Task to be added")
	addTask.Flags().BoolP("isDone", "d", false, "Is the task done?")
}

func addNewTask(cmd *cobra.Command, args []string) {
	todoList, err := todo.Load()
	if err != nil {
		log.Panicln(err)
		return
	}

	taskName, _ := cmd.Flags().GetString("task")
	isDone, _ := cmd.Flags().GetBool("isDone")

	if taskName == "" {
		log.Println("Task name is required")
		return
	}

	newTask := task.Create(task.CreateParameters{Task: taskName, IsDone: isDone})

	todo.Add(&todoList, newTask)

	todo.Save(&todoList)

	log.Println("Task added successfully")
}