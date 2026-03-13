package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("todo-cli")

	todos := Todos{}

	todo1 := Todo{
		Title:       "Clean my desk",
		Completed:   true,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	todo2 := Todo{
		Title:       "Write a blog",
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	todos = append(todos, todo1, todo2)
	for i, t := range todos {
		fmt.Printf("%d | %v\n", i, t)
	}
}
