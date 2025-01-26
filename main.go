package main

import (
	"log"
	"os"

	"esport-booking-backend/config"
	"esport-booking-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found. Ensure environment variables are set.")
	}

	// Initialize Fiber app
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Database connection
	config.ConnectDB()

	// Register routes
	routes.SetupRoutes(app)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000" // Default port
	}
	log.Printf("Server is running on port %s...", port)
	log.Fatal(app.Listen(":" + port))
}
