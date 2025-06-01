package models

import (
	"fmt"

	"gorm.io/gorm"
)

const (
	EmailAndPassword = iota
	GoogleAuth
)

var AuthTypes = map[int]string{
	EmailAndPassword: "email_and_pass",
	GoogleAuth:       "google_auth",
}

type UserRepository interface {
	InsertOne(user *User) *User
}

type User struct {
	gorm.Model
	Name         string `gorm:"index;not null"`
	Email        string `gorm:"uniqueIndex;not null"`
	UserName     string `gorm:"uniqueIndex;not null";json:"user_name"`
	Password     string `not null`
	AuthMethod   int    `not null`
	IsActive     bool   `not null`
	RoleID       *int   `gorm:"null"`
	Role         Role   `gorm:"foreignKey:RoleID"`
	UserType     int    `gorm:"index;not null"`
	ProfileImage string `not null`
	Age          int    `not null`
	Gender       string `not null`
}

func main() {
	val := AuthTypes[EmailAndPassword]
	fmt.Print(val)
}

func (c User) InsertOne(db *gorm.DB, u *User) (*User, error) {
	err := db.Save(&u).Error

	if err != nil {
		return nil, err
	}
	return u, nil
}
