package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/ping", ping)

	app.Listen(":8000")
}

func ping(ctx *fiber.Ctx) error {
	ctx.WriteString("pong")
	return nil
}
