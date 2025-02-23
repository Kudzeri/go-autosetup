package types

// Options описывает выбор пользователя для генерации проекта
type Options struct {
	ProjectName string // Название проекта (имя каталога)
	Framework   string // "none" или "gin"
	Database    string // "sqlite", "postgres", "mongodb"
	Swagger     bool   // Подключать ли swagger
	Cors		bool   // Подключать ли CORS
}
