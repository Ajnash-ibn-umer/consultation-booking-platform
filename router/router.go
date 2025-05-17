package router

import (
	"ajnash-ibn-ummer/wewithyou/pkg/routes"

	"github.com/gin-gonic/gin"
)

func RouterRegister(app *gin.Engine) {

	v1 := app.Group("/v1")
	routes.UserRouter(v1.Group("/user"))
	routes.AuthRouter(v1.Group("/auth"))

}
