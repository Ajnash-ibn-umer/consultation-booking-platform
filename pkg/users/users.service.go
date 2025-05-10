package users

import (
	"ajnash-ibn-ummer/wewithyou/config"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context, user *User, uid uint) *gorm.DB {

	db := config.GetDB()

	if user.Name == "" {
		log.Panic("User Name is Required")
		c.JSON(404, gin.H{
			"message": "User Name is Required",
			"status":  404,
		})
		c.Abort()

	}

	userResp := db.Create(User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	})
	return userResp
}
