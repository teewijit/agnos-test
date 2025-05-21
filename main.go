package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/teewijit/agnos-test/config"
	"github.com/teewijit/agnos-test/routes"
)

func setup() *gin.Engine {
	config.DB = config.InitDB()
	r := gin.Default()
	routes.SetupRoutes(r)
	return r
}

func main() {
	app := setup()
	app.Run(":" + os.Getenv("PORT"))
}
