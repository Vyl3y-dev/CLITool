package clitool

import (
	"fmt"
)

type Task struct {
	ID          int
	Description string
	Done        bool
}

func AddTask(tasks []Task, desc string) []Task {
	// append a new Task struct to the slice
	newTask := Task{
		ID:          len(tasks) + 1, // give it a simple ID
		Description: desc,           // what the user typed
		Done:        false,          // new tasks start incomplete
	}

	tasks = append(tasks, newTask)
	return tasks
}

func ListTasks(tasks []Task) {
	// loop through slice and print
	if len(tasks) == 0 {
		fmt.Println("No tasks yet!")
	} else {
		for _, t := range tasks {
			fmt.Printf("[%d] %s (done: %v)\n", t.ID, t.Description, t.Done)
		}
	}
}

func MarkDone(tasks []Task, id int) []Task {
	// find task by ID and set Done = true
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
		}

	}
	return tasks
}

func DeleteTask(tasks []Task, id int) []Task {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	return tasks
}
