package models

type User struct {
	Id        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique"`
	Password  []byte `gorm:"not null"`
	RoleId    uint   `gorm:"not null"`
	Role      Role   `gorm:"foreignKey:RoleId"`
}
