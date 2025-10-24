package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
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

func formatTask(task Task) string {
	var done string

	if !task.Done {
		done = "[ ]"
	} else {
		done = "[X]"
	}

	return fmt.Sprintf("%s %s", done, task.Text)
}

func main() {
	tasks, err := loadTasks(todosFilename)

	if err != nil {
		fmt.Println("error with parsing tasks file:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

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

			text := strings.Join(parts[2:], " ")
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
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, formatTask(task))
			}
		case "done":
			if len(parts) < 2 {
				fmt.Println("You need to type task number")
				return
			}

			num, err := strconv.Atoi(parts[1])

			if err != nil {
				fmt.Println("Cannot to parse tasks number")
				return
			}

			if num < 1 || num > len(tasks) {
				fmt.Println("Invalid task number")
				return
			}

			tasks[num-1].Done = true
			doneText := fmt.Sprintf("You completed the task %d: %s", num, tasks[num-1].Text)

			if err := saveTasks(todosFilename, tasks); err != nil {
				fmt.Println("Cannot save the file with tasks")
				return
			}

			fmt.Println(doneText)
		default:
			fmt.Println("Unknown command")
		}
	}
}
