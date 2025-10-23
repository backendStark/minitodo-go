package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

const todosFilename = "todos.json"

func loadTasks(filename string) ([]Task, error) {
	if data, err := os.ReadFile(filename); err == nil {
		if len(data) == 0 {
			return []Task{}, nil
		}

		var tasks []Task

		if err := json.Unmarshal(data, &tasks); err != nil {
			return nil, fmt.Errorf("error parsing json %w", err)
		}

		return tasks, nil
	} else if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(filename)

		if err != nil {
			return nil, fmt.Errorf("error creating json file %w", err)
		}

		defer file.Close()

		file.Write([]byte("[]"))

		return []Task{}, nil
	} else {
		return nil, fmt.Errorf("file has errors %w", err)
	}
}

func saveTasks(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		return fmt.Errorf("cannot serialize tasks for file %w", err)
	}

	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("If you want use program, you need to input arguments. Try again")
		return
	}

	tasks, err := loadTasks(todosFilename)

	if err != nil {
		fmt.Println("error with parsing tasks file:", err)
		return
	}

	act := os.Args[1]

	switch act {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("You need to type task text")
			return
		}

		text := strings.Join(os.Args[2:], " ")
		trimmed := strings.TrimSpace(text)

		if len(trimmed) == 0 {
			fmt.Println("You type empty space, you need to input real task")
			return
		}

		newTask := Task{
			Text: trimmed,
			Done: false,
		}

		tasks = append(tasks, newTask)

		if err := saveTasks(todosFilename, tasks); err != nil {
			fmt.Println("Cannot save the file with tasks")
			return
		}

		fmt.Println("Added task:", trimmed)
	case "list":
		fmt.Println("Show list")
	case "done":
		fmt.Println("Make the task done")
	default:
		fmt.Println("Unknown command")
	}
}
