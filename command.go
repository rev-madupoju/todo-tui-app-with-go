package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	CreateAndAddTodo      string
	ReadTodos             bool
	UpdateTodoTitle       string
	DeleteTodo            int
	ToggleTodoAsCompleted int
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.CreateAndAddTodo, "ca", "", "Create a new todo specifying the title")
	flag.BoolVar(&cf.ReadTodos, "ls", false, "List/Read the todos collection")
	flag.StringVar(&cf.UpdateTodoTitle, "ut", "", "Update the title of an existing todo by index <id:new_title>")
	flag.IntVar(&cf.DeleteTodo, "rm", -1, "Delete an existing todo by index")
	flag.IntVar(&cf.ToggleTodoAsCompleted, "tc", -1, "Toggle a todo as completed by index")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(t *Todos) {
	switch {
	case cf.CreateAndAddTodo != "":
		t.createAndAdd(cf.CreateAndAddTodo)

	case cf.ReadTodos:
		t.loadTodos()

	case cf.UpdateTodoTitle != "":
		parts := strings.SplitN(cf.UpdateTodoTitle, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		t.updateTitleByIndex(index, parts[1])

	case cf.ToggleTodoAsCompleted != -1:
		t.toggleCompletedByIndex(cf.ToggleTodoAsCompleted)

	case cf.DeleteTodo != -1:
		t.deleteByIndex(cf.DeleteTodo)

	default:
		fmt.Println("Invalid command")
	}
}
