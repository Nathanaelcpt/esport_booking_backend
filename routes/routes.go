package routes

import (
	"esport-booking-backend/controllers"
	"esport-booking-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Public routes
	app.Post("/login", controllers.Login)
	app.Post("/register", controllers.Register)

	// Protected routes
	api := app.Group("/api", middleware.JWTMiddleware())
	api.Get("/tournaments", controllers.GetTournaments)
	api.Post("/tournaments", controllers.CreateTournament)
	api.Get("/seats/:tournament_id", controllers.GetSeats)
	api.Post("/book", controllers.BookSeat)
}
