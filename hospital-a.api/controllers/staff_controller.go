package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teewijit/agnos-test/config"
	"github.com/teewijit/agnos-test/services"
)

func InsertStaffController(c *gin.Context) {
	var req services.InsertStaff // ใช้ struct ที่ปลอดภัยสำหรับ input เท่านั้น
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateStaff(config.DB, req.Username, req.Password, req.HospitalID)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "username already exists" {
			status = http.StatusConflict // 409
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Staff created successfully",
	})
}
