package controllers

import (
	"go-auth-api-sample/database"
	"go-auth-api-sample/models"

	"github.com/gofiber/fiber/v2"
)

func Ping(ctx *fiber.Ctx) error {
	var user models.User
	database.DB.Preload("Permissions").Preload("Role").First(&user)
	return ctx.JSON(user)
}
