package models

type User struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null" json:"firstName"`
	LastName  string `gorm:"not null" json:"lastName"`
	Email     string `gorm:"unique" json:"email"`
	Password  []byte `gorm:"not null" json:"-"`
	RoleId    uint   `gorm:"not null" json:"roleId"`
	Role      Role   `gorm:"foreignKey:RoleId"`
}
