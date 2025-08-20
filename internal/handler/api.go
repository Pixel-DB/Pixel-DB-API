package handler

import (
	_ "github.com/Pixel-DB/Pixel-DB-API/docs"
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/dto"
	"github.com/Pixel-DB/Pixel-DB-API/internal/model"
	"github.com/gofiber/fiber/v2"
)

// Hello godoc
// @Summary Hello
// @Description This is the base route. You can check, if the API is online.
// @Tags Base
// @Success		200	{object}	dto.APIResponse
// @Router / [get]
func Hello(c *fiber.Ctx) error {
	db := database.DB
	s := new(model.Stats)

	db.Where(&model.Stats{ID: 1}).First(s)

	respone := dto.APIResponse{
		Status:  "success",
		Message: "Hello? I'm okay!",
		Data: dto.APIResponseData{
			TotalRequests:    s.RequestCount,
			TotalUsers:       0,
			TotalPixelArts:   0,
			TotalGithubStars: 0,
		},
	}

	return c.JSON(respone)
}
