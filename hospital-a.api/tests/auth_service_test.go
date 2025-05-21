package services_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/teewijit/agnos-test/services"
)

func TestLoginStaff(t *testing.T) {
	gormDB, mock := NewMockDB()

	t.Run("Valid Credentials", func(t *testing.T) {
		hashedPassword, err := services.HashPassword("testpass123")
		require.NoError(t, err)

		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "staffs" WHERE (username = $1 AND hospital_id = $2) AND "staffs"."deleted_at" IS NULL ORDER BY "staffs"."id" LIMIT $3`,
		)).WithArgs("testuser", 1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "hospital_id"}).
				AddRow(1, "testuser", hashedPassword, 1))

		staff, err := services.LoginStaff(gormDB, "testuser", "testpass123", 1)

		require.NoError(t, err)
		require.NotNil(t, staff)
		assert.Equal(t, "testuser", staff.Username)
	})
}
