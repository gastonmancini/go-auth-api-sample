package database

import (
	"go-auth-api-sample/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/go-auth-api-sample?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	database.AutoMigrate(&models.Permission{}, &models.Role{}, &models.User{})
	DB = database
}
