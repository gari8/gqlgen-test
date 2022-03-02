package resolver

import (
	"context"
	"github.com/gari8/gqlgen-gorm-tutorial/model/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryResolver_Todos(t *testing.T) {
	todoRepo := getMockITodoRepo(t)
	todoRepo.EXPECT().GetTodos(gomock.Any()).Return(nil).AnyTimes()
	iTx := getMockITx(t)
	iTx.EXPECT().DoInTx(gomock.Any()).Return(nil).AnyTimes()
	qr := &queryResolver{NewResolver(iTx,
		todoRepo)}

	todos, err := qr.Todos(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, todos, []*domain.Todo{
		{
			ID:   "123456",
			Text: "abcd",
		},
	})
}
