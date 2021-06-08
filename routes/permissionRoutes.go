package routes

import (
	"go-auth-api-sample/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPermissionRoutes(app *fiber.App) {
	app.Get("/api/permissions", controllers.AllPermissions)
}
