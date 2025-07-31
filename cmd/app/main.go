package main

import (
	"github.com/Pixel-DB/Pixel-DB-API/config"
	_ "github.com/Pixel-DB/Pixel-DB-API/docs"
	"github.com/Pixel-DB/Pixel-DB-API/internal/database"
	"github.com/Pixel-DB/Pixel-DB-API/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

// @title Pixel Schwanz
// @version 1.1
// @description Pixel
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: config.Config("FRONTEND_URL"),
	}))

	router.SetupRouter(app)
	database.ConnectDB()

	log.Fatal(app.Listen(":3000"))
}
