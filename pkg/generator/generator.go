package generator

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/Kudzeri/go-autosetup/pkg/generator/actions"
	"github.com/Kudzeri/go-autosetup/pkg/generator/templates"
	"github.com/Kudzeri/go-autosetup/pkg/generator/types"
)

var folders = []string{
	"model",
	"repository",
	"service",
	"controller",
	"route",
	"middleware",
	"utils",
	"config",
}

func GenerateProject(opts types.Options) error {
	// Создаём корневую папку и подпапки
	if err := actions.CreateFolders(opts.ProjectName, folders); err != nil {
		return err
	}

	// Создаём main.go
	mainContent := templates.GetMainTemplate(opts)
	if err := actions.WriteFile(opts.ProjectName, "main.go", mainContent); err != nil {
		return err
	}

	// Создаём config/db.go
	dbContent := templates.GetDBTemplate(opts)
	if err := actions.WriteFile(opts.ProjectName, filepath.Join("config", "db.go"), dbContent); err != nil {
		return err
	}

	// Создаём .env файл
	envContent := templates.GetEnvTemplate(opts)
	if err := actions.WriteFile(opts.ProjectName, ".env", envContent); err != nil {
		return err
	}

	// Создаём go.mod
	goModContent := templates.GetGoModTemplate(opts)
	if err := actions.WriteFile(opts.ProjectName, "go.mod", goModContent); err != nil {
		return err
	}

	// Создаём пустой go.sum (будет заполнен после `go mod tidy`)
	if err := actions.WriteFile(opts.ProjectName, "go.sum", ""); err != nil {
		return err
	}

	// Если выбрана SQLite, создаём пустой файл базы данных
	if opts.Database == "sqlite" {
		if err := actions.WriteFile(opts.ProjectName, "app.db", ""); err != nil {
			return err
		}
	}

	// Если выбран Swagger, создаём папку docs с placeholder‑файлом
	if opts.Swagger {
		if err := actions.CreateFolders(opts.ProjectName, []string{"docs"}); err != nil {
			return err
		}
		if err := actions.WriteFile(opts.ProjectName, filepath.Join("docs", "swagger.go"), "// Swagger docs will be generated here\n"); err != nil {
			return err
		}
	}

	// Если выбран CORS, генерируем middleware для CORS
	if opts.Cors {
		corsContent := templates.GetCORSTemplate(opts)
		if err := actions.WriteFile(opts.ProjectName, filepath.Join("middleware", "cors.go"), corsContent); err != nil {
			return err
		}
	}

	// Автоматически запускаем команду `go mod tidy` в директории проекта
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = opts.ProjectName
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ошибка при выполнении 'go mod tidy': %v, вывод: %s", err, output)
	}

	return nil
}
