package commands

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/Antonious-Stewart/task-manager-cli/internal/types"
)

func MarkInProgress() {
	id := parseID()

	path, err := getPath()

	if err != nil {
		log.Fatal(err)
	}

	var data []*types.Task

	file, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &data)

	if err != nil {
		log.Fatal(err)
	}

	updated := false
	for _, task := range data {
		if task.ID == id {
			task.Status = types.IN_PROGRESS.String()
			task.UpdatedAt = time.Now()
			updated = true
		}
	}

	if !updated {
		log.Fatal("No task was updated")
	}

	writeback, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path, writeback, 0666)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Task was updated")
}
