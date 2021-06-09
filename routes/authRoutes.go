package routes

import (
	"go-auth-api-sample/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPublicAuthRoutes(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
}

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/me", controllers.GetCurrentUser)
	app.Put("/api/me", controllers.UpdateCurrentUserInfo)
	app.Put("/api/me/password", controllers.UpdateCurrentUserPassword)
	app.Post("/api/me/image", controllers.UpdateCurrentUserProfileImage)
}
