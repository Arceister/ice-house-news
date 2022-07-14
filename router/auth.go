package router

import (
	"github.com/Arceister/ice-house-news/handler"
	"github.com/Arceister/ice-house-news/server"
	"github.com/go-chi/chi/v5"
)

type AuthRoute struct {
	server      server.Server
	authHandler handler.AuthHandler
}

func NewAuthRouter(
	server server.Server,
	authHandler handler.AuthHandler,
) AuthRoute {
	return AuthRoute{
		server:      server,
		authHandler: authHandler,
	}
}

func (r AuthRoute) Setup(chi *chi.Mux) *chi.Mux {
	return chi
}
