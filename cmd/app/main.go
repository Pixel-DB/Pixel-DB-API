package main

import (
	"fmt"
	"log"

	"github.com/Pixel-DB/Pixel-DB-API/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	log.Fatal(app.Listen(":3000"))

	fmt.Println(config.Config("PORT"))
}
