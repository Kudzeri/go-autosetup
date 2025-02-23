package templates

import "github.com/Kudzeri/go-autosetup/pkg/generator/types"

// GetDBTemplate возвращает содержимое config/db.go в зависимости от выбранной базы данных
func GetDBTemplate(opts types.Options) string {
	switch opts.Database {
	case "postgres":
		return getPostgresTemplate()
	case "mongodb":
		return getMongoDBTemplate()
	default: // sqlite
		return getSQLiteTemplate()
	}
}

func getPostgresTemplate() string {
	return `package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "user=youruser password=yourpassword dbname=yourdb sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to PostgreSQL:", err)
		return
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Connection check error:", err)
		return
	}
	fmt.Println("Connected to PostgreSQL")
}
`
}

func getMongoDBTemplate() string {
	return `package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Ping error MongoDB:", err)
		return
	}
	Client = client
	fmt.Println("Connected to MongoDB")
}
`
}

func getSQLiteTemplate() string {
	return `package config

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		fmt.Println("Error connecting to SQLite:", err)
		return
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Error checking connection to SQLite:", err)
		return
	}
	fmt.Println("Connected to SQLite")
}
`
}
