package main

import (
	"flag"
	"log"
	"os"

	"github.com/Antonious-Stewart/task-manager-cli/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No command line args passed:", flag.Args())
	}

	switch command := os.Args[1]; command {
	case "add":
		commands.AddTask()
	case "mark-in-progress":
		commands.MarkInProgress()
	case "mark-done":
		// commands.MarkDone()
	case "list":
		commands.List()
	case "delete":
		commands.Delete()
	}
}
