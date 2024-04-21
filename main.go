package main

import (
	"fmt"
	"github.com/ariefmahendra/crud-api-article/controller"
	"github.com/ariefmahendra/crud-api-article/midleware"
	"github.com/ariefmahendra/crud-api-article/shared/service"
	"github.com/ariefmahendra/crud-api-article/usecase"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	jwt := service.NewJwtServiceImpl()

	middleware := midleware.NewMiddleware(jwt)

	authUsecase := usecase.NewAuthUsecase(jwt)
	todoUsecase := usecase.NewTodoUsecase()

	authController := controller.NewAuthController(r, authUsecase)
	todoController := controller.NewTodoController(r, todoUsecase)

	r.Use(middleware.AuthMiddleware)
	r.Mount("/todos", todoController.Routes())
	r.Mount("/auth", authController.Routes())

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}

	fmt.Println("server running on port 3000")

}
