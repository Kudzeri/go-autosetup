package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kudzeri/go-autosetup/pkg/generator"
	"github.com/Kudzeri/go-autosetup/pkg/generator/types"
	"github.com/Kudzeri/go-autosetup/pkg/utils"
)

func main() {
	// Выводим красивый ASCII‑арт с подписью
	utils.PrintASCII()

	reader := bufio.NewReader(os.Stdin)

	// Ввод названия проекта
	fmt.Print("Enter the name of the project: ")
	projectName, _ := reader.ReadString('\n')
	projectName = strings.TrimSpace(projectName)

	// Выбор фреймворка
	fmt.Println("Choose a framework:")
	fmt.Println("0) Without")
	fmt.Println("1) Gin Gonic")
	fmt.Println("2) Fiber")
	fmt.Print("Enter the number: ")
	frameworkChoiceStr, _ := reader.ReadString('\n')
	frameworkChoiceStr = strings.TrimSpace(frameworkChoiceStr)

	// Выбор базы данных
	fmt.Println("Select a database:")
	fmt.Println("0) Sqlite")
	fmt.Println("1) PostgreSQL")
	fmt.Println("2) MongoDB")
	fmt.Print("Enter the number: ")
	dbChoiceStr, _ := reader.ReadString('\n')
	dbChoiceStr = strings.TrimSpace(dbChoiceStr)

	// Swagger: нужен или нет
	fmt.Print("Add Swagger? (yes/no): ")
	swaggerChoiceStr, _ := reader.ReadString('\n')
	swaggerChoiceStr = strings.TrimSpace(swaggerChoiceStr)

	// CORS: нужен или нет
	fmt.Print("Add CORS? (yes/no): ")
	corsChoiceStr, _ := reader.ReadString('\n')
	corsChoiceStr = strings.TrimSpace(corsChoiceStr)

	// Отображение выбора в переменные
	var framework string
	switch frameworkChoiceStr {
	case "1":
		framework = "gin"
	case "2":
		framework = "fiber"
	default:
		framework = "none"
	}

	var db string
	switch dbChoiceStr {
	case "1":
		db = "postgres"
	case "2":
		db = "mongodb"
	default:
		db = "sqlite"
	}

	swagger := false
	if strings.ToLower(swaggerChoiceStr) == "yes" || strings.ToLower(swaggerChoiceStr) == "y" {
		swagger = true
	}

	cors := false
	if strings.ToLower(corsChoiceStr) == "yes" || strings.ToLower(corsChoiceStr) == "y" {
		cors = true
	}

	// Формируем структуру опций
	options := types.Options{
		ProjectName: projectName,
		Framework:   framework,
		Database:    db,
		Swagger:     swagger,
		Cors:        cors,
	}

	// Генерируем проект
	err := generator.GenerateProject(options)
	if err != nil {
		fmt.Println("Error during project generation:", err)
		return
	}
	fmt.Println("The project is successfully created!")
}
