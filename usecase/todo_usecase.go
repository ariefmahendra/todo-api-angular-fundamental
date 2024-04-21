package usecase

import (
	"github.com/ariefmahendra/crud-api-article/model"
	"github.com/google/uuid"
	"time"
)

var todos []model.Todo

type TodoUsecase interface {
	GetAll() ([]model.Todo, error)
	Create(todo model.Todo) (model.Todo, error)
}

type TodoUsecaseImpl struct {
}

func NewTodoUsecase() *TodoUsecaseImpl {
	return &TodoUsecaseImpl{}
}

func (u *TodoUsecaseImpl) GetAll() ([]model.Todo, error) {
	return todos, nil
}

func (u *TodoUsecaseImpl) Create(todo model.Todo) (model.Todo, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return model.Todo{}, err
	}

	todo.Id = newUUID.String()
	todo.CreatedAt = time.Now().Format(time.DateOnly)
	todos = append(todos, todo)

	return todo, nil
}
