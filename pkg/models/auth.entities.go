package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string           `gorm:"not null"`
	Permissions []RolePermission `gorm:"foreignKey:RoleID"`
}

type RolePermission struct {
	gorm.Model
	RoleID   int    `gorm:"not null"`
	Resource string `gorm:"not null"`
	Actions  string `gorm:"not null"`
	Path     string `gorm:"not null"`
	Role     Role
}
