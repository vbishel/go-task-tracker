package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Add() {
	if len(os.Args) < 3 {
		log.Fatal("Expected task description as second argument")
	}

	description := os.Args[2]

	data, file := GetTasksFromJSON()

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal("Failed to close file")
		}
	}()

	var nextId int

	if len(data) == 0 {
		nextId = 0
	} else {
		nextId = data[len(data)-1].ID + 1
	}

	data = append(data, &Task{
		ID:          nextId,
		Description: description,
		Status:      "in-progress",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	})

	data.WriteToJSON(file)

	fmt.Printf("Task added successfully: (ID: %d)", nextId)
}
