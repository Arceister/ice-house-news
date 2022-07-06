package router

import (
	"github.com/Arceister/ice-house-news/controller"
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/server"
	"github.com/go-chi/chi/v5"
)

type UsersRoute struct {
	server          server.Server
	middlewareJWT   middleware.MiddlewareJWT
	usersController controller.UsersController
}

func NewUsersRouter(
	server server.Server,
	middlewareJWT middleware.MiddlewareJWT,
	usersController controller.UsersController,
) UsersRoute {
	return UsersRoute{
		server:          server,
		middlewareJWT:   middlewareJWT,
		usersController: usersController,
	}
}

func (r UsersRoute) Setup(chi *chi.Mux) *chi.Mux {
	chi.Get("/api/users/{uuid}", r.usersController.GetOneUserController)
	chi.Post("/api/users", r.usersController.CreateUserController)
	chi.Put("/api/users/{uuid}", r.usersController.UpdateUserController)
	chi.Delete("/api/users/{uuid}", r.usersController.DeleteUserController)
	return chi
}
