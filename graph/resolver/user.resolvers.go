package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/dgryski/trifles/uuid"

	"github.com/gari8/gqlgen-gorm-tutorial/graph/generated"
	"github.com/gari8/gqlgen-gorm-tutorial/model/domain"
)

func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*domain.User, error) {
	u := uuid.UUIDv4()
	return r.IUserRepo.InsertUser(u, name)
}

func (r *queryResolver) Users(ctx context.Context) ([]*domain.User, error) {
	users, err := r.GetUsers()
	if err != nil {
		return nil, message.build(RegisterTokenInternalServerError, nil)
	}
	return users, nil
}

func (r *userResolver) Todos(ctx context.Context, obj *domain.User) ([]*domain.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
