package main

import (
	"log"
	"os"
	"strconv"
)

func Delete() {
	if len(os.Args) < 3 {
		log.Fatal("Expected task ID as second argument")
	}

	id, err := strconv.Atoi(os.Args[2])

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

	var newTasks Tasks = make(Tasks, 0, len(tasks) - 1);

	for _, task := range tasks {
		if task.ID == id {
			continue;
		}

		newTasks = append(newTasks, task)
	}

	newTasks.WriteToJSON(file)
}