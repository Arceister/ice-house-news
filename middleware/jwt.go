package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	utils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/golang-jwt/jwt/v4"

	response "github.com/Arceister/ice-house-news/server/response"
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
			response.ErrorResponse(w, utils.NewUnauthorizedError("Malformed token"))
			return
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
				return []byte(m.appConfig.SecretKey), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "JWTProps", claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else if verification, ok := err.(*jwt.ValidationError); ok {
				if verification.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					response.ErrorResponse(w, utils.NewUnauthorizedError("Token expired"))
					return
				}
			} else {
				response.ErrorResponse(w, utils.NewUnauthorizedError(err.Error()))
				return
			}

			if !token.Valid {
				response.ErrorResponse(w, utils.NewUnauthorizedError("Token invalid"))
				return
			}
		}
	})
}

func (m MiddlewareJWT) GenerateNewToken(user entity.User) (*string, time.Time, error) {
	timeExpire := time.Now().AddDate(0, 0, 7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"exp":   timeExpire.Unix(),
	})

	tokenString, err := token.SignedString([]byte(m.appConfig.SecretKey))

	if err != nil {
		return nil, time.Time{}, err
	}

	return &tokenString, timeExpire, nil
}

func (m MiddlewareJWT) ExtractClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(m.appConfig.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, err
	} else {
		return nil, err
	}
}
