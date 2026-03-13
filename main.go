package main

func main() {
	todos := Todos{}
	
	todos.add("Clean my desk")
	todos.add("Write a blog")
	todos.add("Read the Go docs")
	todos.add("Check out awesome-rust projs on gh")
	todos.add("Go to gym at 5PM")

	todos.toggle(0) // mark "Clean my desk as completed"

	todos.delete(2) // delete "Read the Go docs"

	todos.edit(2, "Check out awesome-rust github projects ASAP")

	todos.print()
}
