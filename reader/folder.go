package reader

import (
	"fmt"
	"os"
	"wizzy/core/model"
)

func listFolders(path string) ([]model.Folder, error) {
	var folders []model.Folder

	// Open the directory
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", path, err)
	}

	// Loop through directory entries
	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, model.Folder{
				Name: entry.Name(),
				Desc: "TODO: yet to implement",
			})
		}
	}

	return folders, nil
}
