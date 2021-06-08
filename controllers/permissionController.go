package controllers

import (
	"go-auth-api-sample/database"
	"go-auth-api-sample/models"

	"github.com/gofiber/fiber/v2"
)

func AllPermissions(ctx *fiber.Ctx) error {
	var permissions []models.Permission
	database.DB.Find(&permissions)
	return ctx.JSON(permissions)
}
