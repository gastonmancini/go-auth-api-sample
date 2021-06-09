package controllers

import (
	"go-auth-api-sample/database"
	"go-auth-api-sample/middlewares"
	"go-auth-api-sample/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllUsers(ctx *fiber.Ctx) error {
	if err := middlewares.IsUserAuthorized(ctx, "users"); err != nil {
		return err
	}
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	return ctx.JSON(
		models.Paginate(database.DB, &models.User{}, page),
	)
}

func GetUser(ctx *fiber.Ctx) error {
	if err := middlewares.IsUserAuthorized(ctx, "users"); err != nil {
		return err
	}
	userId, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	var user models.User
	err = database.DB.Preload("Role").First(&user, userId).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}
	return ctx.JSON(user)
}

func CreateUser(ctx *fiber.Ctx) error {
	if err := middlewares.IsUserAuthorized(ctx, "users"); err != nil {
		return err
	}
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}
	const DefaultPasswordForNewUsers = "1234"
	user.SetPassword(DefaultPasswordForNewUsers)
	result := database.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return ctx.JSON(user)
}

func UpdateUser(ctx *fiber.Ctx) error {
	if err := middlewares.IsUserAuthorized(ctx, "users"); err != nil {
		return err
	}
	userId, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}
	user.Id = uint(userId)
	result := database.DB.Model(&user).Updates(user)
	if result.Error != nil {
		return err
	}
	return ctx.JSON(user)
}

func DeleteUser(ctx *fiber.Ctx) error {
	if err := middlewares.IsUserAuthorized(ctx, "users"); err != nil {
		return err
	}
	userId, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	result := database.DB.Delete(&models.User{}, userId)
	if result.Error != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound)
	}
	ctx.Status(fiber.StatusNoContent)
	return nil
}
