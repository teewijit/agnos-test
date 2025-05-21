package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// เข้ารหัส password
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// ตรวจสอบ password กับ hash ที่ให้มา
func CheckPasswordHash(hash, password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
