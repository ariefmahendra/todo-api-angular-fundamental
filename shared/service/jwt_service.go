package service

import (
	"errors"
	"github.com/ariefmahendra/crud-api-article/model"
	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(email string) (string, error)
	ValidateToken(tokenString string) (model.CustomClaims, error)
}

type JwtServiceImpl struct {
}

func NewJwtServiceImpl() *JwtServiceImpl {
	return &JwtServiceImpl{}
}

func (js *JwtServiceImpl) GenerateToken(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, model.CustomClaims{
		Email: email,
	})

	signedString, err := claims.SignedString([]byte("key-secret"))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func (js *JwtServiceImpl) ValidateToken(tokenString string) (model.CustomClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("key-secret"), nil
	})

	if err != nil {
		return model.CustomClaims{}, err
	}

	valid := token.Valid
	if !valid {
		return model.CustomClaims{}, errors.New("invalid token")
	}

	return model.CustomClaims{}, nil
}
