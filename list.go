package main

import (
	"fmt"
	"os"
)

func List() {
	var status string;

	if len(os.Args) > 2 {
		status = os.Args[2]
	}

	tasks, file := GetTasksFromJSON()
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Failed to close file")
		}
	}()

	for _, task := range tasks {
		if status != "" && status != task.Status {
			continue
		}

		fmt.Print(task.ID)
		fmt.Print(" ")
		fmt.Print(task.Description)
		fmt.Print(" ")
		fmt.Print(task.CreatedAt)
		fmt.Println()
	}
}
