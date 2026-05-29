package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name        string       `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Email       string       `gorm:"column:email;type:varchar(32);not null;unique" json:"email"`
	Phone       string       `gorm:"column:phone;type:varchar(15);not null;unique" json:"phone"`
	Password    string       `gorm:"column:password;type:varchar(150);not null" json:"password"`
	Avatar      string       `gorm:"column:avatar;type:varchar(100)" json:"avatar"`
	Roles       []Role       `gorm:"many2many:admin_roles;" json:"roles"`
	Permissions []Permission `gorm:"many2many:admin_permissions;" json:"permissions"`
}
