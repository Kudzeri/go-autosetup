package templates

import "github.com/Kudzeri/go-autosetup/pkg/generator/types"

// GetCORSTemplate возвращает шаблон для CORS middleware.
func GetCORSTemplate(opts types.Options) string {
	switch opts.Framework {
	case "gin":
		return `package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORSMiddleware customizes CORS for Gin.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
`
	case "fiber":
		return `package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// CORSMiddleware customizes CORS for Fiber.
func CORSMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(204)
		}
		return c.Next()
	}
}
`
	default:
		return `package middleware

// CORSMiddleware - stub for CORS, customize for your framework.
func CORSMiddleware() {}
`
	}
}
