package commands

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/Antonious-Stewart/task-manager-cli/internal/types"
)

func Mark(status types.Status) {
	id := parseID()
	path, err := getPath()

	errFatal(err)

	data := unmarshalTask(path)

	updated := false
	for _, task := range data {
		if task.ID == id {
			task.Status = status.String()
			task.UpdatedAt = time.Now()
			updated = true
		}
	}

	if !updated {
		log.Fatal("No task was updated")
	}

	writeback, err := json.Marshal(data)

	errFatal(err)

	err = os.WriteFile(path, writeback, 0666)

	errFatal(err)

	log.Printf("Task was updated")
}
