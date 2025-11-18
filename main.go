package main

import (
	"CLITool/clitool"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // skip the program name

	if len(args) == 0 {
		fmt.Println("Usage: go run . [add|list|done|delete] ...")
		return
	}

	command := args[0]

	switch command {
	case "add":
		// example call:
		// add a new task
	case "list":
		// show all tasks
	case "done":
		// mark task done
	case "delete":
		// delete task
	default:
		fmt.Println("Unknown command:", command)
	}

	tasks := clitool.LoadTasks()

	if len(os.Args) > 1 && os.Args[1] == "gui" {
		clitool.LaunchUI(tasks)
		return
	}

}
