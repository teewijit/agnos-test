package services

import (
	"errors"

	"github.com/teewijit/agnos-test/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type InsertStaff struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required,min=4"` // รหัสผ่านอย่างน้อย 6 ตัวอักษร
	HospitalID uint   `json:"hospital_id" binding:"required"`
}

// เข้าระบบ
func LoginStaff(db *gorm.DB, username, password string, hospitalID uint) (*models.Staff, error) {
	var staff models.Staff
	err := db.Where("username = ? AND hospital_id = ?", username, hospitalID).First(&staff).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(password)); err != nil {
		return nil, errors.New("incorrect password")
	}

	return &staff, nil
}
