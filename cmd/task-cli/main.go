package main

import (
	"fmt"
	"os"
	"task-tracker/internal/service"
	"task-tracker/internal/storage"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <comman> [arguments]")
		fmt.Println("Commands: add, update, delete, mark, list")
		return
	}

	storage := storage.NewJSONStorage("jsonStorage")
	tracker := service.NewTracker(storage)
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <description> ")
			return
		}
		tracker.Add(os.Args[2])
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> <description>")
			return
		}
		tracker.Update(os.Args[2], os.Args[3])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		tracker.Delete(os.Args[2])
	case "mark":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli mark <id> <status>")
			return
		}
		tracker.Mark(os.Args[2], os.Args[3])
	case "list":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli list <status>")
			return
		}
		tracker.List(os.Args[2])
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: add, update, delete, mark, list")
	}
}
