package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gari8/gqlgen-gorm-tutorial/config"
	"github.com/gari8/gqlgen-gorm-tutorial/graph/model"
	"github.com/gari8/gqlgen-gorm-tutorial/model/domain"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*domain.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*domain.Todo, error) {
	var todos []*domain.Todo
	if err := r.ITx.DoInTx(ctx, func(ctx context.Context) error {
		return r.ITodoRepo.GetTodos(todos)
	}); err != nil {
		return nil, err
	}

	cfg := config.Configure("../../config")
	todos = append(todos, &domain.Todo{
		ID:   cfg.Env.ID,
		Text: cfg.Env.Title})
	return todos, nil
}
