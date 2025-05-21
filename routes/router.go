package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/teewijit/agnos-test/controllers"
	"github.com/teewijit/agnos-test/middlewares"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "Respose from",
			"hostname": hostname,
		})
	})

	// Auth Routes
	r.POST("/staff/create", controllers.InsertStaffController)
	r.POST("/staff/login", controllers.LoginStaffController)

	// Protected Routes	// Protected Routes
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	// Patient Routes
	protected.GET("/patient/search/:id", controllers.SearchPatient)

}
