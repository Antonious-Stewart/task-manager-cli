package commands

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func Update() {
	id := parseID()

	if len(os.Args) < 4 {
		log.Fatal("No data passed to be set for the new value.")
	}

	path, err := getPath()

	errFatal(err)

	data := unmarshalTask(path)
	updated := false

	for _, task := range data {
		if id == task.ID {
			task.Description = os.Args[3]
			task.UpdatedAt = time.Now()
			updated = true
		}
	}

	if !updated {
		log.Fatal("No task was updated")
	}

	writeBack, err := json.Marshal(data)

	errFatal(err)

	err = os.WriteFile(path, writeBack, 0666)

	errFatal(err)
}
