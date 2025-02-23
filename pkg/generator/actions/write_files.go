package actions

import (
	"os"
	"path/filepath"
)

// WriteFile записывает контент в файл по указанному относительному пути внутри проекта.
func WriteFile(projectName, relativePath, content string) error {
	fullPath := filepath.Join(projectName, relativePath)
	return os.WriteFile(fullPath, []byte(content), 0644)
}
