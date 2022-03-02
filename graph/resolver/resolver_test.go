package resolver

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getMockITodoRepo(t *testing.T) *MockITodoRepo {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return NewMockITodoRepo(ctrl)
}

func getMockITx(t *testing.T) *MockITx {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return NewMockITx(ctrl)
}

func TestNewResolver(t *testing.T) {
	itx := getMockITx(t)
	itRepo := getMockITodoRepo(t)
	assert.Equal(t, &Resolver{
		ITx:       itx,
		ITodoRepo: itRepo,
	}, NewResolver(itx, itRepo))
}
