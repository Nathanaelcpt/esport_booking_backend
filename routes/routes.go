package routes

import (
	"esport-booking-backend/controllers"
	"esport-booking-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Protected routes
	api := app.Group("/api", middleware.JWTMiddleware())
	api.Get("/tournaments", controllers.GetTournaments) // Fetch tournaments
	api.Post("/tournaments", controllers.CreateTournament)
	api.Get("/seats/:tournament_id", controllers.GetSeats)
	api.Post("/book", controllers.BookSeat)

	app.Get("/protected", middleware.Protect(), func(c *fiber.Ctx) error {
		return c.SendString("This is a protected route")
	})

	// Public routes
	app.Post("/login", controllers.Login)
	app.Post("/register", controllers.RegisterUser)
}
