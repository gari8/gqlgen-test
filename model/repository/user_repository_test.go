package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepo_GetUsers(t *testing.T) {
	db, mock, err := NewMockGormDB()
	assert.NoError(t, err)
	userRepo := NewUserRepo(db)

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("abcdefg", "taro"))
	mock.ExpectCommit()

	_, err = userRepo.GetUsers()
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
