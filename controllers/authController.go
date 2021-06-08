package controllers

import "github.com/gofiber/fiber/v2"

func Ping(ctx *fiber.Ctx) error {
	ctx.WriteString("pong")
	return nil
}
