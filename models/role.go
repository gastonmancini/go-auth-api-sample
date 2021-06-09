package models

import "gorm.io/gorm"

type Role struct {
	Id          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"not null" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions" json:"permissions"`
}

func (role Role) Count(db *gorm.DB) int64 {
	var totalRoles int64
	db.Model(&Role{}).Count(&totalRoles)
	return totalRoles
}

func (role Role) Take(db *gorm.DB, limit int, offset int) interface{} {
	var roles []Role
	db.Preload("Permissions").Offset(offset).Limit(limit).Find(&roles)
	return roles
}
