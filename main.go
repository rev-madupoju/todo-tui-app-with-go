package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("todo-cli")

	todos := Todos{
		{
			Title:       "Clean my desk",
			Completed:   true,
			CompletedAt: nil,
			CreatedAt:   time.Now(),
		},
	}

	todos.add("Write a blog")
	todos.delete(0)

	fmt.Printf("%+v\n\n", todos)
}
