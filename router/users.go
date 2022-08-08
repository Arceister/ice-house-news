package router

import (
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/server"
	"github.com/go-chi/chi/v5"

	usersHandler "github.com/Arceister/ice-house-news/handler/users"
)

type UsersRoute struct {
	server        server.Server
	middlewareJWT middleware.MiddlewareJWT
	usersHandler  usersHandler.UsersHandler
}

func NewUsersRouter(
	server server.Server,
	middlewareJWT middleware.MiddlewareJWT,
	usersHandler usersHandler.UsersHandler,
) UsersRoute {
	return UsersRoute{
		server:        server,
		middlewareJWT: middlewareJWT,
		usersHandler:  usersHandler,
	}
}

func (r UsersRoute) Setup(chi *chi.Mux) *chi.Mux {
	chi.Get("/api/users/{uuid}", r.usersHandler.GetOneUserHandler)
	chi.With(r.middlewareJWT.JwtMiddleware).Get("/api/me/profile", r.usersHandler.GetOwnProfile)
	chi.Post("/api/users", r.usersHandler.CreateUserHandler)
	chi.Put("/api/users/{uuid}", r.usersHandler.UpdateUserHandler)
	chi.Delete("/api/users/{uuid}", r.usersHandler.DeleteUserHandler)
	return chi
}
