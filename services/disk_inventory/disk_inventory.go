package disk_inventory

import (
	"lid/services/logger"
	"os"
	"path/filepath"
)

func listRecursively(dir string) []string {
	log := logger.CreateLogger("disk_inventory.listRecursively")
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Error("something went wrong walking the directory", err, "path", path)
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		log.Error("Error listing "+dir, err)
	}

	return files
}

func BuildFileList(paths ...string) []string {
	log := logger.CreateLogger("disk_inventory.BuildFileList")
	var files []string

	for _, path := range paths {
		fileStat, err := os.Stat(path)
		if err != nil {
			log.Error("Can't read file", err, "file", path)
			continue
		}

		if fileStat.IsDir() {
			subDirList := listRecursively(path)
			files = append(files, subDirList...)
		} else {
			files = append(files, path)
		}
	}

	return files
}
