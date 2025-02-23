package templates

import "github.com/Kudzeri/go-autosetup/pkg/generator/types"

// GetEnvTemplate возвращает шаблон .env файла в зависимости от выбранной базы данных
func GetEnvTemplate(opts types.Options) string {
	switch opts.Database {
	case "postgres":
		return `APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdb`
	case "mongodb":
		return `APP_PORT=8080
MONGO_URI=mongodb://localhost:27017`
	default: // sqlite
		return `APP_PORT=8080
SQLITE_DB=./app.db`
	}
}
