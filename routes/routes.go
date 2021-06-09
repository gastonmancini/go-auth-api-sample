package routes

import (
	"go-auth-api-sample/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	SetupPublicAuthRoutes(app)

	app.Use(middlewares.IsUserAuthenticated) // All the routes defined below this call require the user to be authenticated

	SetupAuthRoutes(app)
	SetupUserRoutes(app)
	SetupRoleRoutes(app)
	SetupPermissionRoutes(app)
}
