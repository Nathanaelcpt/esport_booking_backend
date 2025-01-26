package main

import (
	"log"
	"os"

	"esport-booking-backend/config"
	"esport-booking-backend/models"
	"esport-booking-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found. Ensure environment variables are set.")
	}

	// Initialize Fiber app
	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // atau tentukan domain frontend Anda
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))
	app.Use(logger.New()) // Request logging

	// Database connection
	config.ConnectDB()

	// Database migrations
	if err := config.DB.AutoMigrate(&models.Users{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	log.Println("Database migrated successfully!")

	// Register routes
	routes.SetupRoutes(app)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3001" // Default port if APP_PORT is not set
	}
	log.Printf("Server is running on port %s...", port)
	log.Fatal(app.Listen(":" + port))
}
