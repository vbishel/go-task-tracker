package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func Mark(status string) {
	if len(os.Args) < 3 {
		log.Fatal("Expected to have task ID as 2nd parameter")
	}

	taskID, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatal("Incorrect task ID")
	}

	tasks, file := GetTasksFromJSON()

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal("Failed to close file")
		}
	}()

	var newTasks Tasks = make(Tasks, 0, len(tasks))

	for _, task := range tasks {
		if task.ID == taskID {
			task.Status = status
		}

		newTasks = append(newTasks, task)
	}

	newTasks.WriteToJSON(file)
	fmt.Println("Task marked successfully")
}
