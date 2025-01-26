package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load the .env file
	err := godotenv.Load() // Declare 'err' once
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Buat DSN (Data Source Name) dari variabel .env
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Hubungkan ke database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Reuse the same 'err' variable
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully!")
}
