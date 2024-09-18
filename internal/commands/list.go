package commands

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Antonious-Stewart/task-manager-cli/internal/types"
)

func List() {
	path, err := getPath()

	errFatal(err)

	bytes, err := os.ReadFile(path)

	errFatal(err)

	if string(bytes) == "" {
		log.Println("No task found. Try add some tasks in order to see them in a list.")
		return
	}

	var data []types.Task
	err = json.Unmarshal(bytes, &data)

	errFatal(err)

	var positionalArg string

	if len(os.Args) > 2 {
		positionalArg = os.Args[2]
	}

	switch positionalArg {
	case types.TODO.String():
		printTasks(&data, types.TODO)
	case types.IN_PROGRESS.String():
		printTasks(&data, types.IN_PROGRESS)
	case types.DONE.String():
		printTasks(&data, types.DONE)
	default:
		if positionalArg != "" {
			log.Fatalf("Unknown flag '%s' please try again", positionalArg)
		}
		printAllTasks(&data)
	}
	log.Println("Done printing all tasks....")
}

func printAllTasks(data *[]types.Task) {
	for _, chunk := range *data {
		log.Println(chunk)
	}
}

func printTasks(data *[]types.Task, flag types.Status) {
	for _, chunk := range *data {
		if chunk.Status == flag.String() {
			log.Println(chunk)
		}
	}
}
