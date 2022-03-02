package resolver

//go:generate mockgen -source=$GOFILE -destination=./mock.go -package=$GOPACKAGE
//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"
	"github.com/gari8/gqlgen-gorm-tutorial/model/domain"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ITx
	ITodoRepo
	IUserRepo
}

func NewResolver(
	tx ITx,
	todoRepo ITodoRepo,
	userRepo IUserRepo) *Resolver {
	return &Resolver{
		tx,
		todoRepo,
		userRepo,
	}
}

type ITodoRepo interface {
	GetTodos(todos []*domain.Todo) error
}

type ITx interface {
	DoInTx(ctx context.Context, f func(ctx context.Context) error) error
}

type IUserRepo interface {
	InsertUser(userId, name string) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
}
