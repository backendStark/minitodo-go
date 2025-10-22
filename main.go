package main

import (
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Text string
	Done bool
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("If you want use program, you need to input arguments. Try again")
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

		fmt.Println(trimmed)
	case "list":
		fmt.Println("Show list")
	case "done":
		fmt.Println("Make the task done")
	default:
		fmt.Println("Unknown command")
	}
}
