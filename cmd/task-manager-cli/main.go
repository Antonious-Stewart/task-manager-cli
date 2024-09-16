package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Antonious-Stewart/task-manager-cli/internal/commands"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No command line args passed:", flag.Args())
		os.Exit(2)
	}

	switch command := os.Args[1]; command {
	case "add":
		commands.AddTask(os.Args[2])
	}
}
