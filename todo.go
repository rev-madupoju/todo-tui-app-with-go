package main

import (
	"errors"
	"fmt"
	"github.com/aquasecurity/table"
	"os"
	"strconv"
	"time"
)

const (
	CompletedPending = "❌"
	CompletedSuccess = "✅"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	newTodo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, newTodo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	tp := *todos
	if err := tp.validateIndex(index); err != nil {
		return err
	}

	*todos = append(tp[:index], tp[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	if err := (*todos).validateIndex(index); err != nil {
		return err
	}

	isCompleted := (*todos)[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		(*todos)[index].CompletedAt = &completionTime
	}

	(*todos)[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	if err := (*todos).validateIndex(index); err != nil {
		return err
	}

	(*todos)[index].Title = title
	return nil
}

// feat: print() for Todos to have a tabular view
func (todos *Todos) print() {
	tbl := table.New(os.Stdout)
	tbl.SetRowLines(false)
	tbl.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, todo := range *todos {
		completed := CompletedPending
		completedAt := ""

		if todo.Completed {
			completed = CompletedSuccess
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.DateTime)
			}
		}

		tbl.AddRow(
			strconv.Itoa(index),
			todo.Title,
			completed,
			todo.CreatedAt.Format(time.DateTime),
			completedAt,
		)

		tbl.Render()
	}
}
