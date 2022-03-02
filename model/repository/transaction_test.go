package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func NewMockGormDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := gorm.Open(mysql.Dialector{
		&mysql.Config{
			DriverName:                "mysql",
			Conn:                      db,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	return gormDB, mock, err
}

func TestTx_DoInTx(t *testing.T) {
	mockDB, _, err := NewMockGormDB()
	assert.NoError(t, err)
	tx := NewTx(mockDB)
	err = tx.DoInTx(context.Background(), func(ctx context.Context) error {
		return nil
	})
	assert.NoError(t, err)
}
