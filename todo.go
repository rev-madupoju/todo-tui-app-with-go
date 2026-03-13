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
	tp := *todos

	if err := tp.validateIndex(index); err != nil {
		return err
	}

	isCompleted := tp[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		tp[index].CompletedAt = &completionTime
	}

	tp[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	tp := *todos

	if err := tp.validateIndex(index); err != nil {
		return err
	}

	tp[index].Title = title

	return nil
}
