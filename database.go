package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	// Required for postgress initialisation
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// CreateConnection connects to the database via Gorm
func CreateConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s",
			host, user, DBName, password,
		),
	)
}
