package main

import (
	"fmt"
	"log"

	"github.com/Pixel-DB/Pixel-DB-API/config"
	"github.com/Pixel-DB/Pixel-DB-API/internal/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRouter(app)
	log.Fatal(app.Listen(":3000"))

	fmt.Println(config.Config("PORT"))
}
