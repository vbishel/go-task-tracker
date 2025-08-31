package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("At least one argument is expected")
		return
	}

	action := os.Args[1]

	switch action {
	case "add":
		Add()
	case "update":
		Update()
	case "delete":
		Delete()
	case "mark-in-progress":
		Mark("in-progress")
	case "mark-done":
		Mark("done")
	case "list":
		List()
	default:
		log.Fatal("Action not found")
	}
}
