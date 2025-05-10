package router

import (
	"ajnash-ibn-ummer/wewithyou/pkg/users"

	"github.com/gin-gonic/gin"
)

func RouterRegister(app *gin.Engine) {

	v1 := app.Group("/v1")
	users.UserRouter(v1.Group("/user"))

}
