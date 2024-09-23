package commands

import (
	"encoding/json"
	"log"
	"os"
)

func Delete() {
	id := parseID()

	path, err := getPath()

	errFatal(err)

	data := unmarshalTask(path)

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

	errFatal(err)

	err = os.WriteFile(path, writeBack, 0666)

	errFatal(err)

	log.Println("Deleted Successfully")
}
