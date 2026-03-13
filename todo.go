package main

import "time"

// define a basic Todo structure
type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// we can define an array i.e., type Todos [10]Todo to store only 10 Todos
// but let's go with a nil slice "Todos" to store Todo type structs with no limit
// since its dynamically-sized and has flexible capacity
// fmt.Println(Todos, len(Todos), cap(Todos)) outputs "[] 0 0""
// if Todos == nil {fmt.Println("nil slice!")} // outputs "nil!"
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
