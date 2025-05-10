package users

import (
	"ajnash-ibn-ummer/wewithyou/pkg/auth"
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

type User struct {
	gorm.Model
	Name         string  `gorm:"index;not null"`
	Email        string  `gorm:"index:unique;not null"`
	UserName     string  `gorm:"index:unique;not null"`
	Password     *string `not null`
	AuthMethod   int     `not null`
	IsActive     bool    `not null`
	RoleID       int     `not null`
	Role         auth.Role
	UserType     int    `gorm:"index;not null"`
	ProfileImage string `not null`
	Age          int    `not null`
	Gender       string `not null`
}

func main() {
	val := AuthTypes[EmailAndPassword]
	fmt.Print(val)
}
