package controllers

import (
	"esport-booking-backend/config"
	"esport-booking-backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetSeats(c *fiber.Ctx) error {
	tournamentID := c.Params("tournament_id")

	var seats []models.Seat
	if err := config.DB.Where("tournament_id = ?", tournamentID).Find(&seats).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch seats"})
	}

	return c.JSON(seats)
}
