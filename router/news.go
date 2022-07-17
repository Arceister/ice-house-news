package router

import (
	"github.com/Arceister/ice-house-news/handler"
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/server"
)

type NewsRoute struct {
	server        server.Server
	middlewareJWT middleware.MiddlewareJWT
	newsHandler   handler.NewsHandler
}

func NewNewsRouter(
	server server.Server,
	middlewareJWT middleware.MiddlewareJWT,
	newsHandler handler.NewsHandler,
) NewsRoute {
	return NewsRoute{
		server:        server,
		middlewareJWT: middlewareJWT,
		newsHandler:   newsHandler,
	}
}
