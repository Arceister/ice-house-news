package router

import (
	"github.com/Arceister/ice-house-news/handler"
	"github.com/Arceister/ice-house-news/server"
	"github.com/go-chi/chi/v5"
)

type AuthRoute struct {
	server      server.Server
	authHandler handler.IAuthHandler
}

func NewAuthRouter(
	server server.Server,
	authHandler handler.IAuthHandler,
) AuthRoute {
	return AuthRoute{
		server:      server,
		authHandler: authHandler,
	}
}

func (r AuthRoute) Setup(chi *chi.Mux) *chi.Mux {
	chi.Post("/api/auth/login", r.authHandler.UserSignInHandler)
	return chi
}
