package repository

import (
	"github.com/gari8/gqlgen-gorm-tutorial/model/domain"
	"gorm.io/gorm"
)

type UserRepo struct {
	*gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db}
}

func (r UserRepo) InsertUser(userId, name string) (*domain.User, error) {
	newUser := &domain.User{
		ID:   userId,
		Name: name,
	}
	return newUser, r.DB.Model(&domain.User{}).Create(newUser).Error
}

func (r UserRepo) GetUsers() ([]*domain.User, error) {
	var users []*domain.User
	err := r.DB.Debug().Find(&users).Error
	return users, err
}
