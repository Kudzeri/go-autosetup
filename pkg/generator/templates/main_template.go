package templates

import (
	"strings"

	"github.com/Kudzeri/go-autosetup/pkg/generator/types"
) // Обратите внимание на корректный импорт вашего модуля

// GetMainTemplate возвращает содержимое main.go с учётом выбранного фреймворка
func GetMainTemplate(opts types.Options) string {
	var sb strings.Builder
	sb.WriteString("package main\n\n")

	switch opts.Framework {
	case "gin":
		sb.WriteString(getGinMainTemplate())
	case "fiber":
		sb.WriteString(getFiberMainTemplate())
	default:
		sb.WriteString(getDefaultMainTemplate())
	}

	return sb.String()
}

func getGinMainTemplate() string {
	return `import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello from Gin!")
	})
	r.GET("/health", func(c *gin.Context) {
    c.JSON(200, gin.H{"status": "ok"})
	})
	
	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}
`
}

func getDefaultMainTemplate() string {
	return `import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(` + "`" + `{"status": "ok"}` + "`" + `))
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
`
}

func getFiberMainTemplate() string {
	return `import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Fiber!")
	})
	app.Get("/health", func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"status": "ok"})
	})

	fmt.Println("Server running on port 8080")
	app.Listen(":8080")
}
`
}
