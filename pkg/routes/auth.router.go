package routes

import (
	"ajnash-ibn-ummer/wewithyou/common"
	"ajnash-ibn-ummer/wewithyou/pkg/dto"
	"ajnash-ibn-ummer/wewithyou/pkg/services"

	"fmt"

	"github.com/gin-gonic/gin"
)

type createRoleDto struct {
	Name string `json:"name`
}

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

	r.POST("/role/create", func(c *gin.Context) {

		var dto createRoleDto
		common.ParseJsonBody(c, &dto)

		resp, err := services.RoleCreate(&dto.Name)
		fmt.Print("resp", resp)

		if err != nil {
			fmt.Print(err)
		}
		c.JSON(resp.Status, resp)

	})

	r.GET("/role", func(c *gin.Context) {

		resp, err := services.RoleList()
		fmt.Print("resp", resp)

		if err != nil {
			fmt.Print(err)
		}
		c.JSON(resp.Status, resp)

	})

}
