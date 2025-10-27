package main

import (
	"fmt"
	"minitodo/storage"
	"minitodo/ui"
)

const todosFilename = "todos.json"

func main() {
	tasks, err := storage.LoadTasks(todosFilename)

	if err != nil {
		fmt.Println("error with parsing tasks file:", err)
		return
	}

	if err := ui.RunInteractiveList(tasks, todosFilename); err != nil {
		fmt.Println("Error running interactive list:", err)
		return
	}
}
