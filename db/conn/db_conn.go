package conn

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// TODO : Set a config to open db

	dsn := "host=a user=a password=a dbname=a port=a sslmode=disable TimeZone=Argentina"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	log.Println("Database connection established.")
	return db, nil
}
