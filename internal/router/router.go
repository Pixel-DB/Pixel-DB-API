package router

import (
	"github.com/Pixel-DB/Pixel-DB-API/internal/handler"
	"github.com/Pixel-DB/Pixel-DB-API/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

func SetupRouter(app *fiber.App) {

	app.Get("/swagger/*", swagger.HandlerDefault)                                  //Docs Route (Swagger-API)
	app.Get("/metrics", monitor.New(monitor.Config{Title: "PixelDB-Server Info"})) //Server Monitor

	api := app.Group("/", logger.New()) //Main Route
	api.Get("/", handler.Hello)

	user := app.Group("/user")                             //User Route
	user.Post("/", handler.CreateUser)                     //Create User
	user.Get("/", middleware.Protected(), handler.GetUser) //Get User by Token
	user.Patch("/", handler.UpdateUser)                    //Update User

	auth := app.Group("/auth")         //Auth Route
	auth.Post("/login", handler.Login) //Login User

	pixelart := app.Group("/pixelart")                                 //Pixel Art Route
	pixelart.Post("/", middleware.Protected(), handler.UploadPixelArt) //Upload a Pixel Art
	pixelart.Get("/", handler.GetAllPixelArts)                         //Get all PixelArts
	pixelart.Get("/:pixelArtID", handler.GetPixelArt)                  //Get one Specific PixelArt by ID
	pixelart.Get("/:pixelArtID/picture", handler.GetPixelArtPicture)   //Get Pixel Art Picture

}
