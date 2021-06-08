package routes

import (
	"go-auth-api-sample/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoleRoutes(app *fiber.App) {
	app.Get("/api/roles", controllers.AllRoles)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Post("/api/roles", controllers.CreateRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)
}
