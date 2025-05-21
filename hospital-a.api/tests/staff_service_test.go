package services_test

import (
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/teewijit/agnos-test/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}

func TestCreateStaff(t *testing.T) {
	// Setup
	gormDB, mock := NewMockDB()

	t.Run("add staff successfully", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "staffs" WHERE username = $1 AND "staffs"."deleted_at" IS NULL`)).
			WithArgs("testuser").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		mock.ExpectBegin()
		mock.ExpectQuery("^INSERT INTO \"staffs\" (.+)$").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		err := services.CreateStaff(gormDB, "testuser", "testpass123", 1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("fail to add staff with existing username", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "staffs" WHERE username = $1 AND "staffs"."deleted_at" IS NULL`)).
			WithArgs("testuser").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

		err := services.CreateStaff(gormDB, "testuser", "testpass123", 1)
		assert.EqualError(t, err, "username already exists")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
