package controller

import (
	"encoding/json"
	"github.com/ariefmahendra/crud-api-article/model"
	"github.com/ariefmahendra/crud-api-article/shared/common"
	"github.com/ariefmahendra/crud-api-article/usecase"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type TodoController struct {
	r  *chi.Mux
	tu usecase.TodoUsecase
}

func NewTodoController(r *chi.Mux, tu usecase.TodoUsecase) *TodoController {
	return &TodoController{r: r, tu: tu}
}

func (tc *TodoController) Routes() *chi.Mux {
	tc.r.Get("/", tc.GetAll)
	tc.r.Post("/", tc.Create)
	return tc.r
}

func (tc *TodoController) Create(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		common.ResponseError(w, http.StatusBadRequest, "BAD REQUEST", "invalid body request")
		return
	}

	createdTodoResponse, err := tc.tu.Create(todo)
	if err != nil {
		common.ResponseError(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR", "error in creating todo")
		return
	}

	common.ResponseSuccess(w, http.StatusCreated, "CREATED", "todo successfully created", createdTodoResponse)
}

func (tc *TodoController) GetAll(w http.ResponseWriter, _ *http.Request) {
	todos, err := tc.tu.GetAll()
	if err != nil {
		common.ResponseError(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR", "error in getting todos")
		return
	}

	common.ResponseSuccess(w, http.StatusOK, "OK", "get all todos successfully", todos)
}
