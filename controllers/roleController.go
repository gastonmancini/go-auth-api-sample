package controllers

import (
	"go-auth-api-sample/database"
	"go-auth-api-sample/models"

	"github.com/gofiber/fiber/v2"
)

func AllRoles(ctx *fiber.Ctx) error {
	var roles []models.Role
	database.DB.Preload("Permissions").Find(&roles)
	return ctx.JSON(roles)
}

func GetRole(ctx *fiber.Ctx) error {
	roleId, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	var role models.Role
	err = database.DB.Preload("Permissions").First(&role, roleId).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}
	return ctx.JSON(role)
}

type RoleDto struct {
	Name        string `json:"name"`
	Permissions []int  `json:"permissions"`
}

func CreateRole(ctx *fiber.Ctx) error {
	var roleDto RoleDto
	if err := ctx.BodyParser(&roleDto); err != nil {
		return err
	}

	var permissions = make([]models.Permission, len(roleDto.Permissions))
	for index, permissionId := range roleDto.Permissions {
		permissions[index] = models.Permission{
			Id: uint(permissionId),
		}
	}

	var role = models.Role{
		Name:        roleDto.Name,
		Permissions: permissions,
	}

	result := database.DB.Create(&role)
	if result.Error != nil {
		return result.Error
	}
	return ctx.JSON(role)
}

func UpdateRole(ctx *fiber.Ctx) error {
	roleId, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	var roleDto RoleDto
	if err := ctx.BodyParser(&roleDto); err != nil {
		return err
	}
	var permissions = make([]models.Permission, len(roleDto.Permissions))
	for index, permissionId := range roleDto.Permissions {
		permissions[index] = models.Permission{
			Id: uint(permissionId),
		}
	}
	var role = models.Role{
		Id:   uint(roleId),
		Name: roleDto.Name,
	}
	database.DB.Model(&role).Association("Permissions").Replace(&permissions) // Replace existing role_permissions many2many
	result := database.DB.Model(&role).Updates(role)
	if result.Error != nil {
		return err
	}
	return ctx.JSON(role)
}

func DeleteRole(ctx *fiber.Ctx) error {
	roleId, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	role := models.Role{
		Id: uint(roleId),
	}
	database.DB.Model(&role).Association("Permissions").Clear() // Remove existing role_permissions many2many before deleting the role
	database.DB.Delete(&role)
	ctx.Status(fiber.StatusNoContent)
	return nil
}
