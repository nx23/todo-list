package cmd

import (
	"log"

	"github.com/nx23/todo-list/internal/todo"
	"github.com/spf13/cobra"
)

var listTasks = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long:  `List tasks from the todo list acording to the filters
Example:
todo-list list -i 1
todo-list list -d true`,
	Run:   list,
}

var Id int
var Done, Pending bool

func init() {
	rootCmd.AddCommand(listTasks)

	listTasks.Flags().IntVarP(&Id, "id", "i", 0, "List tasks by id")
	listTasks.Flags().BoolVarP(&Done, "done", "d", false, "List only done tasks")
	listTasks.Flags().BoolVarP(&Pending, "pending", "p", false, "List only pending tasks")

	listTasks.MarkFlagsMutuallyExclusive("id", "done", "pending")
}

func list(cmd *cobra.Command, args []string) {

	todoList, err := todo.Load()
	if err != nil {
		log.Panicln(err)
	}

	if Id > 0 {
		todo.List(&todoList, todo.FilterParameters{Id: Id})
	} else if Done {
		todo.List(&todoList, todo.FilterParameters{Done: true})
	} else if Pending {
		todo.List(&todoList, todo.FilterParameters{Done: false})
	} else {
		todo.List(&todoList, todo.FilterParameters{All: true})
	}
}