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

func (todos *Todos) createAndAdd(title string) {
	t := *todos

	newTodo := Todo{
		Title:     title,
		CreatedAt: time.Now(),
	}

	*todos = append(t, newTodo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (todos *Todos) deleteByIndex(i int) error {
	t := *todos

	if err := t.validateIndex(i); err != nil {
		return err
	}

	*todos = append(t[:i], t[i+1:]...)

	return nil
}

func (todos *Todos) toggleCompletedByIndex(i int) error {
	t := *todos

	if err := t.validateIndex(i); err != nil {
		return err
	}

	isCompleted := t[i].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[i].CompletedAt = &completionTime
	}

	t[i].Completed = !isCompleted

	return nil
}

func (todos *Todos) updateTitleByIndex(i int, title string) error {
	t := *todos

	if err := t.validateIndex(i); err != nil {
		return err
	}

	t[i].Title = title

	return nil
}

func (todos *Todos) loadTodos() {
	tbl := table.New(os.Stdout)
	tbl.SetRowLines(false)
	tbl.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for i, todo := range *todos {
		completed := CompletedPending
		completedAt := ""

		if todo.Completed {
			completed = CompletedSuccess
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.DateTime)
			}
		}

		tbl.AddRow(
			strconv.Itoa(i),
			todo.Title,
			completed,
			todo.CreatedAt.Format(time.DateTime),
			completedAt,
		)
	}

	tbl.Render()
}
