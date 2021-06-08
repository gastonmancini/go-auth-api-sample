package controllers

import (
	"go-auth-api-sample/database"
	"go-auth-api-sample/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// This method allows a user to self-register.
func Register(ctx *fiber.Ctx) error {
	registerDto := struct {
		Id              uint   `json:"id"`
		FirstName       string `json:"firstName"`
		LastName        string `json:"lastName"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		PasswordConfirm string `json:"passwordConfirm"`
		RoleId          uint   `json:"roleId"`
	}{}

	if err := ctx.BodyParser(&registerDto); err != nil {
		return err
	}

	if registerDto.Password != registerDto.PasswordConfirm {
		ctx.Status(fiber.ErrBadRequest.Code)
		return ctx.JSON(fiber.Map{
			"error": "passwords do not match",
		})
	}

	user := models.User{
		FirstName: registerDto.FirstName,
		LastName:  registerDto.LastName,
		Email:     registerDto.Email,
		RoleId:    registerDto.RoleId,
	}
	user.SetPassword(registerDto.Password)

	err := database.DB.Create(&user).Error
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "successfully registered",
	})
}

const cookieName = "go-auth-api-sample-token"

func Login(ctx *fiber.Ctx) error {
	loginDto := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := ctx.BodyParser(&loginDto); err != nil {
		return err
	}

	user := models.User{
		Email: loginDto.Email,
	}

	if err := user.VerifyPassword(loginDto.Password); err != nil {
		return fiber.ErrUnauthorized
	}

	result := database.DB.Where(&user).First(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		return fiber.ErrUnauthorized
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     cookieName,
		Value:    "some-token", // TODO
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		SameSite: "lax",
	})

	return ctx.JSON(fiber.Map{
		"message": "successfully logged in",
	})
}

func Logout(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:     cookieName,
		Expires:  time.Now().Add(-(22 * time.Hour)), // Set expiry date to the past
		HTTPOnly: true,
		SameSite: "lax",
	})

	return ctx.JSON(fiber.Map{
		"message": "successfully logged out",
	})
}

func GetCurrentUser(ctx *fiber.Ctx) error {
	return nil
}

func UpdateCurrentUserInfo(ctx *fiber.Ctx) error {
	return nil
}

func UpdateCurrentUserPassword(ctx *fiber.Ctx) error {
	return nil
}
