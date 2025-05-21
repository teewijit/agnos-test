package models

import "gorm.io/gorm"

type Hospital struct {
	gorm.Model
	Name     string    `json:"name" gorm:"unique"`
	Staffs   []Staff   `gorm:"foreignKey:HospitalID"`
	Patients []Patient `gorm:"foreignKey:HospitalID"`
}
