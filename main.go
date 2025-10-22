package main

import (
	"fmt"
	"os"
)

type Task struct {
	text string
	done bool
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("If you want use program, you need to input arguments. Try again")
		return
	}

	act := os.Args[1]

	switch act {
	case "add":
		fmt.Println("Add task")
	case "list":
		fmt.Println("Show list")
	case "done":
		fmt.Println("Make the task done")
	default:
		fmt.Println("Unknown command")
	}
}
