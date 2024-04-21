package model

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
