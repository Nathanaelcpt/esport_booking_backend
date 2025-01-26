package controllers

import (
	"os"
	"time"

	"esport-booking-backend/config"
	"esport-booking-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// RegisterUser handles user registration
func RegisterUser(c *fiber.Ctx) error {
	users := new(models.Users)
	if err := c.BodyParser(users); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing request body")
	}

	// Cek apakah email sudah ada di database
	if err := config.DB.Where("email = ?", users.Email).First(&models.Users{}).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already exists"})
	}

	// Simpan pengguna baru tanpa hash password
	if err := config.DB.Create(users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error registering user")
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// Login handles user login and returns a JWT token
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.Users
	// Cek apakah email ada di database
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Bandingkan password langsung tanpa hash
	if user.Password != input.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect password"})
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{"token": t})
}
