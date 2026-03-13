package main

import "fmt"

func main() {
	fmt.Println("todo-cli")
	todos := Todos{}

	todos.add("Clean my desk")
	todos.add("Write a blog")
	todos.add("Read the Go docs")
	todos.add("Check out awesome-rust projs on gh")

	todos.toggle(0)

	todos.delete(2)

	todos.edit(2, "Check out awesome-rust github projects ASAP")

	for i, t := range todos {
		fmt.Printf("| %d | %v | %v | %v | %v |\n", i, t.Title, t.Completed, t.CompletedAt, t.CreatedAt)
	}
}
