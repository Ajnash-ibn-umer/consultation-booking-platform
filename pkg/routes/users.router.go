package routes

import (
	"ajnash-ibn-ummer/wewithyou/common"
	"ajnash-ibn-ummer/wewithyou/pkg/dto"
	"ajnash-ibn-ummer/wewithyou/pkg/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {

	r.POST("/create", func(c *gin.Context) {

		var dto dto.CreateUserDto
		common.ParseJsonBody(c, &dto)
		fmt.Print(dto)

		resp, err := services.CreateUser(&dto, 11)
		fmt.Print("resp", resp)

		if err != nil {
			fmt.Print(err)
		}
		c.JSON(resp.Status, resp)

	})

}
