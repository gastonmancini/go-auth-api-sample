package middlewares

import (
	"go-auth-api-sample/database"
	"go-auth-api-sample/models"
	"go-auth-api-sample/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Naive way of checking if the logged in user has permissions to access the requested resource
// TODO: Improve the authorization middleware
func IsUserAuthorized(ctx *fiber.Ctx, resource string) error {
	token := ctx.Cookies(util.CookieName)
	id, err := util.ParseToken(token)
	if err != nil {
		return fiber.ErrUnauthorized
	}
	userId, _ := strconv.Atoi(id)
	user := models.User{
		Id: uint(userId),
	}

	database.DB.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}

	database.DB.Preload("Permissions").Find(&role)

	if ctx.Method() == "GET" {
		// A user with view or edit permissions over the resource can permorm a GET request
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+resource || permission.Name == "edit_"+resource {
				return nil
			}
		}
	} else {
		// A user only with edit permissions over the resource can permorm a POST/PUT/DELETE request
		for _, permission := range role.Permissions {
			if permission.Name == "edit_"+resource {
				return nil
			}
		}
	}

	return fiber.ErrUnauthorized
}
