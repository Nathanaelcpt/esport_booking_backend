package controllers

import (
	"esport-booking-backend/config"
	"esport-booking-backend/models"

	"github.com/gofiber/fiber/v2"
)

func BookSeat(c *fiber.Ctx) error {
	type BookingInput struct {
		UsersID uint `json:"user_id"`
		SeatID  uint `json:"seat_id"`
	}

	var input BookingInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Check if seat is already booked
	var seat models.Seat
	if err := config.DB.First(&seat, input.SeatID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Seat not found"})
	}

	if seat.IsBooked {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Seat already booked"})
	}

	// Book the seat
	seat.IsBooked = true
	config.DB.Save(&seat)

	booking := models.Booking{UsersID: input.UsersID, SeatID: input.SeatID}
	config.DB.Create(&booking)

	// Optionally: Emit real-time update
	return c.JSON(fiber.Map{"message": "Seat booked successfully", "booking": booking})
}
