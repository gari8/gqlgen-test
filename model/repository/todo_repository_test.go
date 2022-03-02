package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gari8/gqlgen-gorm-tutorial/model/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type TodoRepoTestSuite struct {
	suite.Suite
	todoRepo TodoRepo
	mock     sqlmock.Sqlmock
}

func (suite *TodoRepoTestSuite) SetupTodoRepoTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	todoRepo := TodoRepo{}
	suite.mock = mock
	todoRepo.DB, err = gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	assert.NoError(suite.T(), err)
	suite.todoRepo = *NewTodoRepository(todoRepo.DB)
}

func (suite *TodoRepoTestSuite) TearDownTest() {
	db, _ := suite.todoRepo.DB.DB()
	db.Close()
}

func TestTodoRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TodoRepoTestSuite))
}

func TestNewTodoRepository(t *testing.T) {
	db := &gorm.DB{}
	assert.Equal(t, &TodoRepo{db}, NewTodoRepository(db))
}

func TestTodoRepo_GetTodos(t *testing.T) {
	tts := &TodoRepoTestSuite{
		Suite:    suite.Suite{},
		todoRepo: TodoRepo{},
		mock:     nil,
	}
	tts.SetupTodoRepoTest()

	var todos []*domain.Todo

	err := tts.todoRepo.GetTodos(todos)

	assert.NoError(t, err)

	assert.Equal(t, todos, &domain.Todo{
		ID:   "123456",
		Text: "abcd",
	})

	tts.TearDownTest()
}
