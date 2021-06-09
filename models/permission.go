package models

import "gorm.io/gorm"

type Permission struct {
	Id   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

func (permission Permission) Count(db *gorm.DB) int64 {
	var totalPermissions int64
	db.Model(&Permission{}).Count(&totalPermissions)
	return totalPermissions
}

func (permission Permission) Take(db *gorm.DB, limit int, offset int) interface{} {
	var permissions []Permission
	db.Offset(offset).Limit(limit).Find(&permissions)
	return permissions
}
