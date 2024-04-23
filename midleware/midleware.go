package midleware

import (
	"context"
	"github.com/ariefmahendra/crud-api-article/shared/common"
	"github.com/ariefmahendra/crud-api-article/shared/service"
	"net/http"
	"strings"
)

type Middleware struct {
	jwt service.JwtService
}

func NewMiddleware(jwt service.JwtService) *Middleware {
	return &Middleware{jwt: jwt}
}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() != "/api/v1/auth/login" && r.URL.String() != "/api/v1/auth/register" {
			authHeader := r.Header.Get("Authorization")
			bearerToken := strings.Split(authHeader, " ")

			if len(bearerToken) != 2 {
				common.ResponseError(w, http.StatusUnauthorized, "UNAUTHORIZED", "token is not valid")
				return
			}

			if bearerToken[0] == "" || bearerToken[1] == "" {
				common.ResponseError(w, http.StatusUnauthorized, "UNAUTHORIZED", "token is not valid")
				return
			}

			if bearerToken[0] != "Bearer" {
				common.ResponseError(w, http.StatusUnauthorized, "UNAUTHORIZED", "token type is not valid")
				return
			}

			claims, err := m.jwt.ValidateToken(bearerToken[1])
			if err != nil {
				common.ResponseError(w, http.StatusUnauthorized, "UNAUTHORIZED", "token is not valid")
				return
			}

			ctx := context.WithValue(r.Context(), "jwt-token", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		next.ServeHTTP(w, r)
	})
}
