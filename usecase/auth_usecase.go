package usecase

import (
	"fmt"
	"github.com/ariefmahendra/crud-api-article/model"
	"github.com/ariefmahendra/crud-api-article/model/dto"
	"github.com/ariefmahendra/crud-api-article/shared/service"
)

var datas = []model.User{
	{
		Email:    "ariefmahendra@gmail.com",
		Password: "123456",
	},
}

type AuthUsecase interface {
	Login(payload model.User) (string, dto.LoginResponse, error)
	Register(payload model.User) (dto.LoginResponse, error)
}

type AuthUsecaseImpl struct {
	jwt service.JwtService
}

func NewAuthUsecase(jwt service.JwtService) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{jwt: jwt}
}

func (a *AuthUsecaseImpl) Login(payload model.User) (string, dto.LoginResponse, error) {
	var userModel model.User
	for _, user := range datas {
		if user.Email == payload.Email && user.Password == payload.Password {
			userModel = user
		}
	}

	if (userModel == model.User{}) {
		return "", dto.LoginResponse{}, fmt.Errorf("user not found")
	}

	token, err := a.jwt.GenerateToken(userModel.Email)
	if err != nil {
		return "", dto.LoginResponse{}, err
	}

	return token, dto.LoginResponse{Email: userModel.Email, AccessToken: token}, nil
}

func (a *AuthUsecaseImpl) Register(payload model.User) (dto.LoginResponse, error) {
	datas = append(datas, payload)
	return dto.LoginResponse{Email: payload.Email}, nil
}
