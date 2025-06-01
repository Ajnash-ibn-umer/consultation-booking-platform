package common

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseJsonBody(c *gin.Context, binder interface{}) any {
	if err := c.ShouldBindJSON(&binder); err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, E{
			Message: "Body Parsing failed",
			Status:  http.StatusBadRequest,
		})
	}
	return binder
}

func PasswordHashing() {

}
