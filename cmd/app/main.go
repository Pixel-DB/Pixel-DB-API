package main

import (
	"github.com/Pixel-DB/Pixel-DB-API/config"
	_ "github.com/Pixel-DB/Pixel-DB-API/docs"
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/middleware"
	"github.com/Pixel-DB/Pixel-DB-API/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

var requestCount int64

// @title PixelDB
// @version 0.1
// @description Pixel-BD is an open-source online platform where anyone can upload, share, and showcase their pixel art creations with the community.
// @contact.name Lukas Haible
// @contact.email lukas.haible@web.de
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	app := fiber.New()

	app.Use(
		cors.New(cors.Config{
			AllowOrigins: config.Config("FRONTEND_URL"),
		}),
		func(c *fiber.Ctx) error {
			middleware.UpdateRequestCount(1) //Stats Counter
			return c.Next()
		},
	)

	router.SetupRouter(app)
	database.ConnectDB()

	log.Fatal(app.Listen(":3000"))
}
