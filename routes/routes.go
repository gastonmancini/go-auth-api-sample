package routes

import (
	"go-auth-api-sample/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/ping", controllers.Ping)
}
