package services_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teewijit/agnos-test/services"
)

func init() {
	// ตั้งค่า ENV สำหรับ JWT_SECRET ก่อนรัน test
	os.Setenv("JWT_SECRET", "test-secret-key")
}

func TestGenerateToken(t *testing.T) {
	staffID := uint(1)
	hospitalID := uint(2)

	token, err := services.GenerateToken(staffID, hospitalID)

	assert.NoError(t, err)
	assert.NotEmpty(t, token, "Token should not be empty")
}
func TestValidateToken(t *testing.T) {
	staffID := uint(1)
	hospitalID := uint(2)

	validToken, err := services.GenerateToken(staffID, hospitalID)
	assert.NoError(t, err)

	t.Run("Valid Token", func(t *testing.T) {
		claims, err := services.ValidateToken(validToken)
		assert.NoError(t, err)
		assert.Equal(t, staffID, claims.StaffID)
		assert.Equal(t, hospitalID, claims.HospitalID)
	})

	t.Run("Invalid Token Format", func(t *testing.T) {
		_, err := services.ValidateToken("invalid.token.format")
		assert.Error(t, err)
	})

	t.Run("Expired Token", func(t *testing.T) {
		// สร้าง token ที่หมดอายุแล้ว
		expiredToken := services.CreateExpiredTestToken(staffID, hospitalID)
		_, err := services.ValidateToken(expiredToken)
		assert.Error(t, err)
	})
}
