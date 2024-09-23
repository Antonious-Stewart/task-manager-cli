package commands

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Antonious-Stewart/task-manager-cli/internal/types"
)

func Delete() {
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

	removed := false

	for i, element := range data {
		if element.ID == id {
			data = append(data[:i], data[i+1:]...)
			removed = true
		}
	}

	if !removed {
		log.Printf("No task found with that ID: %v", id)
		return
	}

	writeBack, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path, writeBack, 0666)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Deleted Successfully")
}
