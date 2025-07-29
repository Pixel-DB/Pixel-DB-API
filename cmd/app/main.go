package main

import (
	"log"

	"github.com/Pixel-DB/Pixel-DB-API/config"
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: config.Config("FRONTEND_URL"),
	}))

	router.SetupRouter(app)
	database.ConnectDB()

	log.Fatal(app.Listen(":3000"))
}
