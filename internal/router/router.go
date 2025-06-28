package router

import (
	"github.com/Pixel-DB/Pixel-DB-API/internal/handler"
	"github.com/Pixel-DB/Pixel-DB-API/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouter(app *fiber.App) {

	api := app.Group("/", logger.New()) //Main Route
	api.Get("/", handler.Hello)
	api.Post("/", middleware.Protected(), handler.UploadPixelArt) //Upload a Pixel Art

	user := app.Group("/user")         //User Route
	user.Post("/", handler.CreateUser) //Create User

	auth := app.Group("/auth")         //Auth Route
	auth.Post("/login", handler.Login) //Login User

}
