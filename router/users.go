package router

import "github.com/Arceister/ice-house-news/server"

type UsersRoute struct {
	server server.Server
}

func NewUsersRouter(server server.Server) UsersRoute {
	return UsersRoute{
		server: server,
	}
}
