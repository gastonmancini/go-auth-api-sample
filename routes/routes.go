package routes

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	app.Get("/api/ping", ping)
}

func ping(ctx *fiber.Ctx) error {
	ctx.WriteString("pong")
	return nil
}
