package main

import (
	"fmt"
	"minitodo/config"
	"minitodo/storage"
	"minitodo/ui"
)

func main() {
	tasks, err := storage.LoadTasks(config.DefaultFilename)

	if err != nil {
		fmt.Println("error with parsing tasks file:", err)
		return
	}

	if err := ui.RunInteractiveList(tasks, config.DefaultFilename); err != nil {
		fmt.Println("Error running interactive list:", err)
		return
	}
}
