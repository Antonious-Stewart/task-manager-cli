package commands

import (
	"log"
	"path/filepath"
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
