package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func Update() {
	if len(os.Args) < 4 {
		log.Fatal("Expected task ID as first argument and task description as second argument")
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
			task.Description = os.Args[3]
		}

		newTasks = append(newTasks, task)
	}

	newTasks.WriteToJSON(file)
	fmt.Println("Task updated successfully")
}