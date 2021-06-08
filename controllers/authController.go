package controllers

import (
	"go-auth-api-sample/models"

	"github.com/gofiber/fiber/v2"
)

func Ping(ctx *fiber.Ctx) error {
	user := models.User{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Password:  []byte{1},
		Email:     "john@doe.com",
		RoleId:    1,
		Role: models.Role{
			Id:   1,
			Name: "Admin",
			Permissions: []models.Permission{
				{
					Id:   1,
					Name: "Edit",
				}},
		},
	}
	return ctx.JSON(user)
}
