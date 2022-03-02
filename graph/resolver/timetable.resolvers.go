package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gari8/gqlgen-gorm-tutorial/graph/model"
	"github.com/gari8/gqlgen-gorm-tutorial/model/domain"
)

func (r *mutationResolver) CreateTimeTable(ctx context.Context, input model.NewTimeTable) (*domain.TimeTable, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TimeTables(ctx context.Context) ([]*domain.TimeTable, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TimeTable(ctx context.Context, id string) (*domain.TimeTable, error) {
	panic(fmt.Errorf("not implemented"))
}
