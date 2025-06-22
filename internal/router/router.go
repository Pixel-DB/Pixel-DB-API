package router

import (
	"github.com/Pixel-DB/Pixel-DB-API/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {

	api := app.Group("/") //Main Route
	api.Get("/", handler.Hello)

}
