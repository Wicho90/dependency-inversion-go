package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wicho90/dependency-inversion-go/infrastructure"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := fiber.New()

	infrastructure.Routers(app)
	app.Listen(":" + port)
}
