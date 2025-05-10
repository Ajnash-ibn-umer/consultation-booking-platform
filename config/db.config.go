package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var client *gorm.DB

func InitConfig() *gorm.DB {
	fmt.Print("db configuration started")
	connectionStr := os.Getenv("POSTGRES_DB_CONNECTION")
	if connectionStr == "" {
		log.Fatal("Postgres connection string not found")
	}
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	fmt.Print(err)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	client = db
	fmt.Println("Database connected successfully")
	return db
}

// GetDB returns the database client
func GetDB() *gorm.DB {
	if client == nil {
		client = InitConfig()
	}
	return client
}
