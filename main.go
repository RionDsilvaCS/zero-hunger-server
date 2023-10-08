package main

import (
	"github.com/RionDsilvaCS/zeroHunger/database"
	"github.com/RionDsilvaCS/zeroHunger/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
)

func main() {
	database.Connection()
	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowCredentials: true}))
	routes.Setup(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
