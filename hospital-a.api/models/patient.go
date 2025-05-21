package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	FirstNameTh  string `json:"first_name_th" binding:"required"`
	MiddleNameTh string `json:"middle_name_th"`
	LastNameTh   string `json:"last_name_th" binding:"required"`
	FirstNameEn  string `json:"first_name_en"`
	MiddleNameEn string `json:"middle_name_en"`
	LastNameEn   string `json:"last_name_en"`
	DateOfBirth  string `json:"date_of_birth" binding:"required,datetime=2006-01-02"` // รูปแบบวันที่: YYYY-MM-DD
	PatientHN    string `json:"patient_hn" binding:"required"`
	NationalID   string `json:"national_id" binding:"required_without=PassportID,len=13,numeric"` // อย่างน้อยต้องมี NationalID หรือ PassportID
	PassportID   string `json:"passport_id" binding:"required_without=NationalID"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email" binding:"omitempty,email"`
	Gender       string `json:"gender" binding:"required,oneof=male female other"`
	HospitalID   uint   `json:"hospital_id" binding:"required"`
}
