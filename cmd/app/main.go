package main

import (
	"log"

	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/router"
	"github.com/Pixel-DB/Pixel-DB-API/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRouter(app)
	database.ConnectDB()
	storage.InitMinioClient()

	log.Fatal(app.Listen(":3000"))
}
