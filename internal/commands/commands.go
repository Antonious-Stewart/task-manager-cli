package commands

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Antonious-Stewart/task-manager-cli/internal/types"
)

func AddTask(title string) {
	path, err := filepath.Abs("../../internal/storage/tasks.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	bytes, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

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
		if err != nil {
			log.Fatal(err)
		}

		file.WriteString(string(data))
		return
	}

	err = json.Unmarshal(bytes, &collection)

	if err != nil {
		log.Fatal(err)
	}

	collection = append(collection, types.Task{
		ID:          int64(len(collection)) + 1,
		Description: title,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Status:      types.TODO.String(),
	})

	fileData, err := json.Marshal(collection)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = os.WriteFile(path, fileData, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
}
