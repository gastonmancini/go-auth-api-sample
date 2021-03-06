package controllers

import (
	"go-auth-api-sample/database"
	"go-auth-api-sample/models"
	"go-auth-api-sample/util"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	result := database.DB.Where(&user).First(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		return fiber.ErrUnauthorized
	}

	if err := user.VerifyPassword(loginDto.Password); err != nil {
		return fiber.ErrUnauthorized
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	token, err := util.CreateToken(strconv.Itoa(int(user.Id)), expirationTime)

	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     util.CookieName,
		Value:    token,
		Expires:  expirationTime,
		HTTPOnly: true,
		SameSite: "lax",
	})

	return ctx.JSON(fiber.Map{
		"message": "successfully logged in",
	})
}

func Logout(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:     util.CookieName,
		Expires:  time.Now().Add(-(2 * time.Hour)), // Set expiry date to the past
		HTTPOnly: true,
		SameSite: "lax",
	})

	return ctx.JSON(fiber.Map{
		"message": "successfully logged out",
	})
}

func GetCurrentUser(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies(util.CookieName)
	userId, err := util.ParseToken(cookie)
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

func UpdateCurrentUserInfo(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies(util.CookieName)
	userId, err := util.ParseToken(cookie)
	if err != nil {
		return err
	}
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}
	id, _ := strconv.Atoi(userId)
	user.Id = uint(id)
	result := database.DB.Model(&user).Updates(user)
	if result.Error != nil {
		return err
	}
	return ctx.JSON(user)
}

func UpdateCurrentUserPassword(ctx *fiber.Ctx) error {

	cookie := ctx.Cookies(util.CookieName)
	userId, err := util.ParseToken(cookie)
	if err != nil {
		return err
	}

	updatePasswordDto := struct {
		Password        string `json:"password"`
		PasswordConfirm string `json:"passwordConfirm"`
	}{}

	if updatePasswordDto.Password != updatePasswordDto.PasswordConfirm {
		ctx.Status(fiber.ErrBadRequest.Code)
		return ctx.JSON(fiber.Map{
			"error": "passwords do to match",
		})
	}

	if err := ctx.BodyParser(&updatePasswordDto); err != nil {
		return err
	}

	id, _ := strconv.Atoi(userId)
	user := models.User{
		Id: uint(id),
	}
	user.SetPassword(updatePasswordDto.Password)
	result := database.DB.Model(&user).Updates(user)
	if result.Error != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"message": "password successfully changed",
	})
}

func UpdateCurrentUserProfileImage(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies(util.CookieName)
	userId, err := util.ParseToken(cookie)
	if err != nil {
		return err
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["image"] // This is the key of the file when we send it in the form-data request body
	filename := uuid.New().String()

	// File format = guid.extension
	for _, file := range files {
		filename = filename + filepath.Ext(file.Filename)
		if err := ctx.SaveFile(file, "./uploads/"+filename); err != nil {
			return err
		}
	}

	id, _ := strconv.Atoi(userId)
	user := models.User{
		Id:       uint(id),
		ImageUrl: ctx.BaseURL() + "/api/uploads/" + filename,
	}
	result := database.DB.Model(&user).Updates(user)
	if result.Error != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"url": user.ImageUrl,
	})
}
