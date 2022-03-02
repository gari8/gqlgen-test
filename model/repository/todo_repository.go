package repository

import (
	"fmt"
	"github.com/gari8/gqlgen-gorm-tutorial/config"
	"github.com/gari8/gqlgen-gorm-tutorial/model/domain"
	"gorm.io/gorm"
)

type TodoRepo struct {
	*gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepo {
	return &TodoRepo{DB: db}
}

func (t *TodoRepo) GetTodos(todos []*domain.Todo) error {
	cfg := config.Configure("./config")
	fmt.Println(cfg.Env.ID, ":", cfg.Env.Title)
	todos = append(todos, &domain.Todo{
		ID:   cfg.Env.ID,
		Text: cfg.Env.Title,
	})
	return nil
}
