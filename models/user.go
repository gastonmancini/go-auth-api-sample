package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null" json:"firstName"`
	LastName  string `gorm:"not null" json:"lastName"`
	Email     string `gorm:"unique" json:"email"`
	ImageUrl  string `json:"image" json:"imageUrl"`
	Password  []byte `gorm:"not null" json:"-"`
	RoleId    uint   `gorm:"not null" json:"roleId"`
	Role      Role   `gorm:"foreignKey:RoleId"`
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return nil
}

func (user *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

func (user *User) Count(db *gorm.DB) int64 {
	var totalUsers int64
	db.Model(&User{}).Count(&totalUsers)
	return totalUsers
}

func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User
	db.Preload("Role").Offset(offset).Limit(limit).Find(&users)
	return users
}
