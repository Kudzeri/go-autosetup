package templates

import (
	"fmt"

	"github.com/Kudzeri/go-autosetup/pkg/generator/types"
)

// GetGoModTemplate возвращает содержимое файла go.mod без фиксированных версий зависимостей.
// Зависимости будут разрешаться автоматически при запуске `go mod tidy`.
func GetGoModTemplate(opts types.Options) string {
	moduleName := opts.ProjectName
	template := fmt.Sprintf("module %s\n\ngo 1.24\n\n", moduleName)
	return template
}
