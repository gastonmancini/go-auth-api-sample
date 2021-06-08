package routes

import (
	"go-auth-api-sample/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	app.Get("/api/users", controllers.AllUsers)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Post("/api/users", controllers.CreateUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)
}
