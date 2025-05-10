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
	router.RouterRegister(app)
	app.GET("/ping", ping)

	port := os.Getenv("PORT")
	app.Run(":" + port)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, common.Response{
		Message: "Ping",
		Status:  200,
		Data:    nil,
	})
}
