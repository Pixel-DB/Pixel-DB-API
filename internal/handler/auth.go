package handler

import (
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/Pixel-DB/Pixel-DB-API/internal/security"
	"github.com/Pixel-DB/Pixel-DB-API/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	input := new(LoginInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request", "data": nil})
	}

	userModel, err := utils.GetUserEmail(input.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error finding user", "data": nil})
	}

	if userModel == nil {
		dummyHash := "$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK" //Hash something for Timing Attacks protection
		security.CheckPasswordHash(dummyHash, input.Password)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid Email or Password", "data": nil})
	}

	if !security.CheckPasswordHash(userModel.Password, input.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid Email or Password", "data": nil})
	}

	token, err := utils.GenerateToken(userModel.ID, userModel.Email, userModel.Username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error Signing Token", "data": err.Error()})
	}

	AuthResponse := dto.AuthResponse{
		ID:       userModel.ID,
		Email:    userModel.Email,
		Username: userModel.Username,
		Role:     userModel.Role,
		Token:    token,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Logged in", "data": AuthResponse})
}
