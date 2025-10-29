package main

import (
	"fmt"
	"minitodo/config"
	"minitodo/ui"
)

func main() {

	if err := ui.RunInteractiveList(config.DefaultFilename); err != nil {
		fmt.Println("Error running interactive list:", err)
		return
	}
}
