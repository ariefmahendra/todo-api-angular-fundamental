package controller

import (
	"encoding/json"
	"github.com/ariefmahendra/crud-api-article/model"
	"github.com/ariefmahendra/crud-api-article/shared/common"
	"github.com/ariefmahendra/crud-api-article/usecase"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type AuthController struct {
	r  *chi.Mux
	au usecase.AuthUsecase
}

func NewAuthController(r *chi.Mux, authUC usecase.AuthUsecase) *AuthController {
	return &AuthController{r: r, au: authUC}
}

func (ac *AuthController) Routes() *chi.Mux {
	ac.r.Post("/login", ac.Login)
	ac.r.Post("/register", ac.Register)
	return ac.r
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		common.ResponseError(w, http.StatusBadRequest, "BAD REQUEST", "invalid body request")
		return
	}

	token, loginResponse, err := ac.au.Login(user)
	if err != nil {
		common.ResponseError(w, http.StatusUnauthorized, "UNAUTHORIZED", "unauthorized")
		return
	}

	cookie := http.Cookie{
		Name:     "jwt-token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	common.ResponseSuccess(w, http.StatusOK, "OK", "user login successfully", loginResponse)
}

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		common.ResponseError(w, http.StatusBadRequest, "BAD REQUEST", "invalid body request")
		return
	}

	registerResponse, err := ac.au.Register(user)
	if err != nil {
		common.ResponseError(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR", "Error: Internal Server Error")
		return
	}

	common.ResponseSuccess(w, http.StatusOK, "OK", "user register successfully", registerResponse)
}
