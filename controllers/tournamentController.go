package controllers

import (
	"esport-booking-backend/config"
	"esport-booking-backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetTournaments(c *fiber.Ctx) error {
	var tournaments []models.Tournament
	if err := config.DB.Find(&tournaments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch tournaments"})
	}

	return c.JSON(tournaments)
}

func CreateTournament(c *fiber.Ctx) error {
	var tournament models.Tournament
	if err := c.BodyParser(&tournament); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := config.DB.Create(&tournament).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create tournament"})
	}

	return c.Status(fiber.StatusCreated).JSON(tournament)
}
