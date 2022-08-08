package router

import (
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/server"
	"github.com/go-chi/chi/v5"

	commentHandler "github.com/Arceister/ice-house-news/handler/comment"
)

type CommentRoute struct {
	server         server.Server
	middlewareJWT  middleware.MiddlewareJWT
	commentHandler commentHandler.CommentHandler
}

func NewCommentRoute(
	server server.Server,
	middlewareJWT middleware.MiddlewareJWT,
	commentHandler commentHandler.CommentHandler,
) CommentRoute {
	return CommentRoute{
		server:         server,
		middlewareJWT:  middlewareJWT,
		commentHandler: commentHandler,
	}
}

func (r CommentRoute) Setup(chi *chi.Mux) *chi.Mux {
	chi.With(r.middlewareJWT.JwtMiddleware).Get("/api/news/{newsId}/comment", r.commentHandler.GetCommentsOnNewsHandler)
	chi.With(r.middlewareJWT.JwtMiddleware).Post("/api/news/{newsId}/comment", r.commentHandler.InsertCommentHandler)
	return chi
}
