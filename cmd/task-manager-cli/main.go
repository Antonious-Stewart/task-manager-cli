package main

import (
	"flag"
	"log"
	"os"

	"github.com/Antonious-Stewart/task-manager-cli/internal/commands"
	"github.com/Antonious-Stewart/task-manager-cli/internal/types"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No command line args passed:", flag.Args())
	}

	switch command := os.Args[1]; command {
	case "add":
		commands.AddTask()
	case "mark-in-progress":
		commands.Mark(types.IN_PROGRESS)
	case "mark-done":
		commands.Mark(types.DONE)
	case "mark-todo":
		commands.Mark(types.TODO)
	case "list":
		commands.List()
	case "delete":
		commands.Delete()
	case "update":
		commands.Update()
	default:
		log.Printf("Unrecognized command '%v'\n", os.Args[1])
	}
}
