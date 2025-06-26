package handler

import (
	"log"

	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	u := new(model.Users)
	db := database.DB

	if err := c.BodyParser(u); err != nil {
		return c.JSON(fiber.Map{"status": "Error"})
	}
	log.Println(u)

	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body", "errors": err.Error()})
	}
	if err := db.Create(&u).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "errors": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Create User", "data": nil})

}
