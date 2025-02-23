package actions

import (
	"os"
	"path/filepath"
)

// CreateFolders создаёт корневую папку проекта и все указанные подпапки.
func CreateFolders(projectName string, folders []string) error {
	// Создаём корневую папку
	if err := os.Mkdir(projectName, 0755); err != nil {
		return err
	}

	// Создаём подпапки
	for _, folder := range folders {
		path := filepath.Join(projectName, folder)
		if err := os.Mkdir(path, 0755); err != nil {
			return err
		}
	}
	return nil
}
