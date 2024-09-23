package commands

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"

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

func unmarshalTask(path string) []*types.Task {
	bytes, err := os.ReadFile(path)

	if err != nil {
		log.Fatal()
	}

	var data []*types.Task
	err = json.Unmarshal(bytes, &data)

	if err != nil {
		log.Fatal(err)
	}

	return data
}
