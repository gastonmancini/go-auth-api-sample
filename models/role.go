package models

type Role struct {
	Id          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"not null"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}
