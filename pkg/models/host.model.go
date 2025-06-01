package models

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	Name        string           `gorm:"not null"`
	Permissions []RolePermission `gorm:"foreignKey:RoleID"`
}
