package routes

import (
	"ajnash-ibn-ummer/wewithyou/common"
	"ajnash-ibn-ummer/wewithyou/pkg/dto"
	"ajnash-ibn-ummer/wewithyou/pkg/services"

	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {

	r.POST("/signin", func(c *gin.Context) {

		var dto dto.SignInDto
		common.ParseJsonBody(c, &dto)

		resp, err := services.SignIn(&dto)
		fmt.Print("resp", resp)

		if err != nil {
			fmt.Print(err)
		}
		c.JSON(resp.Status, resp)

	})

}
