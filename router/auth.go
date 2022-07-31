package router

import (
	auth "github.com/Arceister/ice-house-news/handler/auth"
	"github.com/Arceister/ice-house-news/server"
	"github.com/go-chi/chi/v5"
)

type AuthRoute struct {
	server      server.Server
	authHandler auth.AuthHandler
}

func NewAuthRouter(
	server server.Server,
	authHandler auth.AuthHandler,
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
