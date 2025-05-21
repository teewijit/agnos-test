package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teewijit/agnos-test/config"
	"github.com/teewijit/agnos-test/services"
)

func LoginStaffController(c *gin.Context) {
	var input struct {
		Username   string `json:"username" binding:"required"`
		Password   string `json:"password" binding:"required"`
		HospitalID uint   `json:"hospital_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staff, err := services.LoginStaff(config.DB, input.Username, input.Password, input.HospitalID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := services.GenerateToken(staff.ID, staff.HospitalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"staff": gin.H{
			"id":         staff.ID,
			"username":   staff.Username,
			"hospital":   staff.Hospital.Name,
			"hospitalID": staff.HospitalID,
		},
	})
}
