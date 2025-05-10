package users

import (
	"ajnash-ibn-ummer/wewithyou/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {

	r.GET("/create", func(c *gin.Context) {

		userData := User{}
		resp := CreateUser(c, &userData, 11)
		c.JSON(http.StatusOK, common.Response{
			Data:    resp,
			Message: "New User created",
			Status:  http.StatusOK,
		})
	})

}
