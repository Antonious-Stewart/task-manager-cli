package commands

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
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

func parseID() int64 {
	if len(os.Args) < 3 {
		log.Fatal("No id was passed to delete")
	}

	id, err := strconv.ParseInt(os.Args[2], 36, 64)

	if err != nil {
		log.Fatal(err)
	}

	return id
}
