package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Tasks []*Task

func GetTasksFromJSON() (Tasks, *os.File) {
	file, err := os.OpenFile("_tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Could not open file")
	}

	readed, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Failed to read file")
	}

	var data Tasks
	if len(readed) > 0 {
		err = json.Unmarshal(readed, &data)
		if err != nil {
			log.Fatal("failed to unmarshal JSON")
		}
	}

	return data, file
}

func (tasks *Tasks) WriteToJSON(file *os.File) {
	marshaled, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal("Failed to marshal data")
		return
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal("Failed to seek file")
	}

	err = file.Truncate(0)
	if err != nil {
		log.Fatal("Failed to write to file")
	}

	_, err = file.Write(marshaled)
	if err != nil {
		log.Fatal("Failed to write to file")
	}
}
