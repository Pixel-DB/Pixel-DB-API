package handler

import (
	"fmt"

	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"github.com/Pixel-DB/Pixel-DB-API/internal/security"
	"github.com/Pixel-DB/Pixel-DB-API/internal/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateUser(c *fiber.Ctx) error {
	u := new(model.Users)
	db := database.DB

	if err := c.BodyParser(u); err != nil {
		return c.JSON(fiber.Map{"status": "Error"})
	}
	u.Role = "user"

	hashedPassword, err := security.HashPassword(u.Password)
	if err != nil {
		return c.JSON(fiber.Map{"status": "Error hashing password", "error": err.Error()})
	}
	u.Password = hashedPassword

	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body", "error": err.Error()})
	}
	if err := db.Create(u).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "error": err.Error()})
	}

	UserResponse := dto.UserResponse{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		Username:  u.Username,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Role:      u.Role,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created User", "data": UserResponse})

}

func GetUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	userID := utils.GetUserIDFromToken(token)
	user, err := utils.GetUser(userID)
	if err != nil {
		fmt.Println(err.Error())
	}
	data := dto.UserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Get User", "data": data})
}
