package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null" json:"firstName"`
	LastName  string `gorm:"not null" json:"lastName"`
	Email     string `gorm:"unique" json:"email"`
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
