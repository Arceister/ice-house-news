package router

import (
	"github.com/Arceister/ice-house-news/handler"
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/server"
)

type CommentRoute struct {
	server         server.Server
	middlewareJWT  middleware.MiddlewareJWT
	commentHandler handler.CommentHandler
}

func NewCommentRoute(
	server server.Server,
	middlewareJWT middleware.MiddlewareJWT,
	commentHandler handler.CommentHandler,
) CommentRoute {
	return CommentRoute{
		server:         server,
		middlewareJWT:  middlewareJWT,
		commentHandler: commentHandler,
	}
}
