package router

import (
	"github.com/Arceister/ice-house-news/controller"
	"github.com/Arceister/ice-house-news/server"
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
