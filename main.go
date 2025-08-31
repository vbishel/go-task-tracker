package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("At least one argument is expected")
		return
	}

	action := os.Args[1]

	switch action {
	case "add":
		add()
	case "update":
		update()
	case "delete":
		delete()
	case "mark-in-progress":
		mark("in-progress")
	case "mark-done":
		mark("done")
	case "list":
		list()
	default:
		log.Fatal("Action not found")
	}
}

func add() {
	if len(os.Args) < 3 {
		log.Fatal("No tasks added")
	}

	description := os.Args[2]
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Could not open file")
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal("Failed to close file")
		}
	}()

	readed, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Failed to read file")
	}

	var data []Task
	if len(readed) > 0 {
		err = json.Unmarshal(readed, &data)
		if err != nil {
			log.Fatal("failed to unmarshal JSON")
		}
	}

	var nextId int

	if len(data) == 0 {
		nextId = 0
	} else {
		nextId = data[len(data) - 1].ID + 1
	}

	data = append(data, Task{
		ID:          nextId,
		Description: description,
		Status:      "in-progress",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	})

	marshaled, err := json.Marshal(data)
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

	fmt.Printf("Task added successfully: (ID: %d)", nextId)
}

func update() {

}

func delete() {

}

func list() {

}

func mark(status string) {

}
