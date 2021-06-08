package controllers

import (
	"go-auth-api-sample/database"
	"go-auth-api-sample/models"

	"github.com/gofiber/fiber/v2"
)

func AllUsers(ctx *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return ctx.JSON(users)
}

func GetUser(ctx *fiber.Ctx) error {
	userId, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	var user models.User
	err = database.DB.First(&user, userId).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}
	return ctx.JSON(user)
}
