package router

import (
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/server"
	"github.com/go-chi/chi/v5"

	authHandler "github.com/Arceister/ice-house-news/handler/auth"
)

type AuthRoute struct {
	server        server.Server
	middlewareJWT middleware.MiddlewareJWT
	authHandler   authHandler.AuthHandler
}

func NewAuthRouter(
	server server.Server,
	middlewareJWT middleware.MiddlewareJWT,
	authHandler authHandler.AuthHandler,
) AuthRoute {
	return AuthRoute{
		server:        server,
		middlewareJWT: middlewareJWT,
		authHandler:   authHandler,
	}
}

func (r AuthRoute) Setup(chi *chi.Mux) *chi.Mux {
	chi.Post("/api/auth/login", r.authHandler.UserSignInHandler)
	chi.With(r.middlewareJWT.JwtMiddleware).Get("/api/auth/token", r.authHandler.ExtendTokenHandler)
	return chi
}
