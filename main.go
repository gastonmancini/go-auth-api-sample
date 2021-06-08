package main

import (
	"go-auth-api-sample/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
