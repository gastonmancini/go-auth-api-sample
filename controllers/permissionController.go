package controllers

import (
	"go-auth-api-sample/database"
	"go-auth-api-sample/middlewares"
	"go-auth-api-sample/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllPermissions(ctx *fiber.Ctx) error {
	if err := middlewares.IsUserAuthorized(ctx, "permissions"); err != nil {
		return err
	}
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	return ctx.JSON(
		models.Paginate(database.DB, &models.Permission{}, page),
	)
}
