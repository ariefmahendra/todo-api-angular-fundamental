package usecase

import (
	"errors"
	"github.com/ariefmahendra/crud-api-article/model"
	"github.com/ariefmahendra/crud-api-article/model/dto"
	"github.com/google/uuid"
	"time"
)

var todos []model.Todo

type TodoUsecase interface {
	GetAll() ([]model.Todo, error)
	Create(todo model.Todo) (model.Todo, error)
	Delete(payloadId string) (dto.DeleteTodoResponse, error)
	Update(payload dto.TodoUpdateRequest) (model.Todo, error)
	GetById(payloadId string) (model.Todo, error)
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

func (u *TodoUsecaseImpl) Delete(payloadId string) (dto.DeleteTodoResponse, error) {
	for i, todo := range todos {
		if todo.Id == payloadId {
			todos = append(todos[:i], todos[i+1:]...)
			return dto.DeleteTodoResponse{Id: todo.Id}, nil
		}
	}

	return dto.DeleteTodoResponse{}, errors.New("todo not found")
}

func (u *TodoUsecaseImpl) Update(payload dto.TodoUpdateRequest) (model.Todo, error) {
	for i, todo := range todos {
		if todo.Id == payload.Id {
			todos[i].Name = payload.Name
			todos[i].IsCompleted = payload.IsCompleted
			return todo, nil
		}
	}

	return model.Todo{}, errors.New("todo not found")
}

func (u *TodoUsecaseImpl) GetById(payloadId string) (model.Todo, error) {
	for _, todo := range todos {
		if todo.Id == payloadId {
			return todo, nil
		}
	}

	return model.Todo{}, errors.New("todo not found")
}
