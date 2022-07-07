package router

import (
	"github.com/Arceister/ice-house-news/controller"
	"github.com/Arceister/ice-house-news/server"
	"github.com/go-chi/chi/v5"
)

type AuthRoute struct {
	server         server.Server
	authController controller.AuthController
}

func NewAuthRouter(
	server server.Server,
	authController controller.AuthController,
) AuthRoute {
	return AuthRoute{
		server:         server,
		authController: authController,
	}
}

func (r AuthRoute) Setup(chi *chi.Mux) *chi.Mux {
	return chi
}
