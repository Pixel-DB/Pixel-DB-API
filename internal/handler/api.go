package handler

import (
	_ "github.com/Pixel-DB/Pixel-DB-API/docs"
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// Hello godoc
// @Summary Hello
// @Description This is the base route. You can check, if the API is online.
// @Tags Base
// @Success		200	{object}	dto.APIResponse
// @Router / [get]
func Hello(c *fiber.Ctx) error {
	respone := dto.APIResponse{
		Status:  "success",
		Message: "Hello? I'm okay!",
		Data:    "",
	}

	return c.JSON(respone)
}
