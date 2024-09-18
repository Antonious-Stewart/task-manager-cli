package commands

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Antonious-Stewart/task-manager-cli/internal/types"
)

func errFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getPath() (string, error) {
	path, err := filepath.Abs("../../internal/storage/tasks.json")

	if err != nil {
		return "", err
	}

	return path, nil
}

func AddTask() {
	if len(os.Args[1:]) != 2 {
		log.Fatal(errors.New("add requires two arguments"))
	}

	title := os.Args[2]
	path, err := getPath()

	errFatal(err)

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)

	errFatal(err)

	defer file.Close()

	bytes, err := os.ReadFile(path)

	errFatal(err)

	var collection []types.Task

	if string(bytes) == "" {
		collection = []types.Task{
			{
				ID:          1,
				Description: title,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Status:      types.TODO.String(),
			},
		}
		data, err := json.Marshal(collection)

		errFatal(err)

		file.WriteString(string(data))
		return
	}

	err = json.Unmarshal(bytes, &collection)

	errFatal(err)

	collection = append(collection, types.Task{
		ID:          int64(len(collection)) + 1,
		Description: title,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Status:      types.TODO.String(),
	})

	fileData, err := json.Marshal(collection)

	errFatal(err)

	err = os.WriteFile(path, fileData, 0666)
	errFatal(err)

	log.Println("Created Successfully.")
}

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
