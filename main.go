package main

import (
	"github.com/RionDsilvaCS/zeroHunger/database"
	"github.com/RionDsilvaCS/zeroHunger/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connection()
	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowCredentials: true}))
	routes.Setup(app)
	app.Listen(":3000")
}
