package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/server"
	"github.com/golang-jwt/jwt/v4"
)

type MiddlewareJWT struct {
	appConfig lib.App
}

func NewMiddlewareJWT(appConfig lib.App) MiddlewareJWT {
	return MiddlewareJWT{
		appConfig: appConfig,
	}
}

func (m MiddlewareJWT) JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			server.ResponseJSON(w, 401, false, "Malformed token")
			return
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
				return []byte(m.appConfig.SecretKey), nil
			})

			if !token.Valid {
				server.ResponseJSON(w, 500, false, "Invalid secret key!")
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "JWTProps", claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else if verification, ok := err.(*jwt.ValidationError); ok {
				if verification.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					server.ResponseJSON(w, 401, false, "Token expired")
					return
				}
			} else {
				server.ResponseJSON(w, 401, false, err.Error())
				return
			}
		}
	})
}
