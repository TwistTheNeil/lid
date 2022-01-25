package disk_inventory

import (
	"lid/models"
	"lid/services/logger"
	"os"
	"path/filepath"
)

func List(dir string) models.NodeList {
	log := logger.CreateLogger("disk_inventory.List")
	var nl models.NodeList

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Error("something went wrong walking the directory", err, "path", path)
			return err
		}

		if info.IsDir() {
			return nil
		}

		newNode := models.Node{Name: path, Size: info.Size()}
		newNode.Hash()
		nl = nl.Append(newNode)
		return nil
	})

	if err != nil {
		panic(err)
	}

	return nl
}
