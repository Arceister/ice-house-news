package router

import (
	"github.com/Arceister/ice-house-news/controller"
	"github.com/Arceister/ice-house-news/server"
	"github.com/go-chi/chi/v5"
)

type UsersRoute struct {
	server          server.Server
	usersController controller.UsersController
}

func NewUsersRouter(
	server server.Server,
	usersController controller.UsersController,
) UsersRoute {
	return UsersRoute{
		server:          server,
		usersController: usersController,
	}
}

func (r UsersRoute) Setup(chi *chi.Mux) *chi.Mux {
	chi.Get("/api/users/{uuid}", r.usersController.GetOneUserController)
	chi.Post("/api/users", r.usersController.CreateUserController)
	return chi
}
