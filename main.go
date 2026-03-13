package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	todos.add("Clean my desk")
	todos.add("Write a blog")
	todos.add("Read the Go docs")
	todos.add("Check out awesome-rust projs on gh")

	todos.toggle(0) // mark "Clean my desk as completed"

	todos.delete(2) // delete "Read the Go docs"

	todos.edit(2, "Check out awesome-rust github projects ASAP")

	todos.print()

	storage.Save(todos)
}
