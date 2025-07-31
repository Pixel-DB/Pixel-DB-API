package handler

import (
	_ "github.com/Pixel-DB/Pixel-DB-API/docs"
	"github.com/gofiber/fiber/v2"
)

// Hello godoc
// @Summary Hello
// @Description This is the base route. You can check, if the API is online.
// @Tags Base
// @Router / [get]
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
