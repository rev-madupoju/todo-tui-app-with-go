package main

const dataFile = "todos.json"

func main() {
	todos := Todos{}

	storage := NewStorage[Todos](dataFile)
	storage.Load(&todos)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)

	storage.Save(todos)
}
