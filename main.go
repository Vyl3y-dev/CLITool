package clitool

import (
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
}
