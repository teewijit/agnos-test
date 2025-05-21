package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teewijit/agnos-test/services"
)

func TestHashPassword(t *testing.T) {
	t.Run("Valid Password", func(t *testing.T) {
		password := "secure123"
		hashed, err := services.HashPassword(password)
		assert.NoError(t, err)
		assert.NotEmpty(t, hashed)

		err = services.CheckPasswordHash(hashed, password)
		assert.NoError(t, err)
	})

	t.Run("Empty Password", func(t *testing.T) {
		password := ""
		hashed, err := services.HashPassword(password)
		assert.Error(t, err)
		assert.Empty(t, hashed)
	})
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mypassword"
	hashed, err := services.HashPassword(password)
	assert.NoError(t, err)

	t.Run("Correct Password", func(t *testing.T) {
		err := services.CheckPasswordHash(hashed, password)
		assert.NoError(t, err)
	})

	t.Run("Incorrect Password", func(t *testing.T) {
		err := services.CheckPasswordHash(hashed, "wrongpassword")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "hashedPassword")
	})

	t.Run("Empty Password", func(t *testing.T) {
		err := services.CheckPasswordHash(hashed, "")
		assert.Error(t, err)
		assert.EqualError(t, err, "password cannot be empty")
	})
}
