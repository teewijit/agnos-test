package models

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	Username   string   `json:"username" gorm:"unique"`
	Password   string   `json:"password" binding:"required,min=4"` // รหัสผ่านอย่างน้อย 6 ตัวอักษร
	HospitalID uint     `json:"hospital_id" binding:"required"`
	Hospital   Hospital `json:"hospital"`
}
