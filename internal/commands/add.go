package commands

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/Antonious-Stewart/task-manager-cli/internal/types"
)

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
