package main

import (
	"log"

	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRouter(app)
	database.ConnectDB()

	log.Fatal(app.Listen(":3000"))
}
