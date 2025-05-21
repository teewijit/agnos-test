package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teewijit/agnos-test/config"
	"github.com/teewijit/agnos-test/models"
	"gorm.io/gorm"
)

func SearchPatient(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is required"})
		return
	}

	hospitalInterface, exists := c.Get("hospital")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// สมมติ hospitalID เป็น uint ที่เซ็ตจาก middleware แล้ว
	hospitalID, ok := hospitalInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid hospital info"})
		return
	}

	var patient models.Patient

	err := config.DB.Where("(national_id = ? OR passport_id = ?) AND hospital_id = ?", id, id, hospitalID).First(&patient).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patient)
}
