package main

import (
	"errors"
	"fmt"
	"time"
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

// feat: validateIndex() helper func
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

// feat: delete() method to remove a Todo from Todos collection
func (todos *Todos) delete(index int) error {
	tp := *todos

	if err := tp.validateIndex(index); err != nil {
		return err
	}

	*todos = append(tp[:index], tp[index+1:]...)

	return nil
}
