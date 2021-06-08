package models

type Role struct {
	Id          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"not null" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions" json:"permissions"`
}
