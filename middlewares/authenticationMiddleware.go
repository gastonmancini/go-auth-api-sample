package middlewares

import (
	"go-auth-api-sample/util"

	"github.com/gofiber/fiber/v2"
)

func IsUserAuthenticated(ctx *fiber.Ctx) error {
	token := ctx.Cookies(util.CookieName)
	if _, err := util.ParseToken(token); err != nil {
		return fiber.ErrUnauthorized
	}
	return ctx.Next()
}
