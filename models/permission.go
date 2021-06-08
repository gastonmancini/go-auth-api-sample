package models

type Permission struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}
