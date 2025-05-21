package services

import (
	"errors"

	"github.com/teewijit/agnos-test/models"
	"gorm.io/gorm"
)

// สร้าง staff
func CreateStaff(db *gorm.DB, username, password string, hospitalID uint) error {

	// เข้ารหัสรหัสผ่าน
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return errors.New("hashing password failed")
	}

	staff := models.Staff{Username: username, Password: hashedPassword, HospitalID: hospitalID}

	// ตรวจสอบว่าชื่อผู้ใช้นี้มีอยู่แล้วหรือไม่
	var count int64
	db.Model(&models.Staff{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return errors.New("username already exists")
	}

	// Save the new user
	result := db.Create(&staff)
	return result.Error
}
