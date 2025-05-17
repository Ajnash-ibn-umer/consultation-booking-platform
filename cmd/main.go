//   Product Api:
//    version: 0.1
//    title: Consultation Booking
//   Schemes: http, https
//   Host:
//   BasePath: /api/v1
//      Consumes:
//      - application/json
//   Produces:
//   - application/json
//   SecurityDefinitions:
//    Bearer:
//     type: apiKey
//     name: Authorization
//     in: header
//   swagger:meta

package main

import (
	"ajnash-ibn-ummer/wewithyou/common"
	"ajnash-ibn-ummer/wewithyou/config"
	"ajnash-ibn-ummer/wewithyou/migration"
	"ajnash-ibn-ummer/wewithyou/router"
	"fmt"

	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env not found")
	}
	// db configuration
	config.InitConfig()
	db := config.GetDB()
	migration.MigrateEntities()

	fmt.Print((db))
	// routers
	app := gin.Default()
	app.Use(gin.Recovery())
	router.RouterRegister(app)
	app.GET("/ping", ping)

	port := os.Getenv("PORT")
	app.Run(":" + port)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, common.Resp{
		Message: "Ping",
		Status:  200,
		Data:    nil,
	})
}
