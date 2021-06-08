package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	SetupPublicAuthRoutes(app)
	SetupAuthRoutes(app)
	SetupUserRoutes(app)
	SetupRoleRoutes(app)
	SetupPermissionRoutes(app)
}
