package handler

import (
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"github.com/Pixel-DB/Pixel-DB-API/internal/security"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	input := new(LoginInput)

	if err := c.BodyParser(input); err != nil {
		return c.JSON(fiber.Map{"status": "Error"})
	}

	userModel := new(model.Users)

	db := database.DB
	db.Where(model.Users{Email: input.Email}).First(userModel)

	if !security.CheckPasswordHash(userModel.Password, input.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid identity or password", "data": nil})
	}

	AuthResponse := dto.AuthResponse{
		ID:       userModel.ID,
		Email:    userModel.Email,
		Username: userModel.Username,
		Role:     userModel.Role,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Logged in", "data": AuthResponse})
}
