package main

import (
	"bufio"
	"fmt"
	"minitodo/models"
	"minitodo/storage"
	"minitodo/ui"
	"os"
	"strings"
)

const todosFilename = "todos.json"

func main() {
	tasks, err := storage.LoadTasks(todosFilename)

	if err != nil {
		fmt.Println("error with parsing tasks file:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(`Input one of commands:
    add <text> - add a new task
    list       - view interactive task list
    quit       - exit program`)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		input = strings.TrimSpace(input)

		if input == "quit" {
			break
		}

		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		act := parts[0]

		switch act {
		case "add":
			if len(parts) < 3 {
				fmt.Println("You need to type task text")
				return
			}

			text := strings.Join(parts[1:], " ")
			trimmed := strings.TrimSpace(text)

			if len(trimmed) == 0 {
				fmt.Println("You type empty space, you need to input real task")
				return
			}

			newTask := models.Task{
				Text: trimmed,
				Done: false,
			}

			tasks = append(tasks, newTask)

			if err := storage.SaveTasks(todosFilename, tasks); err != nil {
				fmt.Println("Cannot save the file with tasks")
				return
			}

			fmt.Println("Added task:", trimmed)
		case "list":
			if err := ui.RunInteractiveList(tasks, todosFilename); err != nil {
				fmt.Println("Error running interactive list:", err)
				return
			}

			tasks, err = storage.LoadTasks(todosFilename)

			if err != nil {
				fmt.Println("Error reloading tasks:", err)
				return
			}
		default:
			fmt.Println("Unknown command")
		}
	}
}
