package router

import (
	"github.com/Arceister/ice-house-news/handler"
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/server"
	"github.com/go-chi/chi/v5"
)

type NewsRoute struct {
	server        server.Server
	middlewareJWT middleware.MiddlewareJWT
	newsHandler   handler.INewsHandler
}

func NewNewsRouter(
	server server.Server,
	middlewareJWT middleware.MiddlewareJWT,
	newsHandler handler.INewsHandler,
) NewsRoute {
	return NewsRoute{
		server:        server,
		middlewareJWT: middlewareJWT,
		newsHandler:   newsHandler,
	}
}

func (r NewsRoute) Setup(chi *chi.Mux) *chi.Mux {
	chi.With(r.middlewareJWT.JwtMiddleware).Post("/api/news", r.newsHandler.AddNewNewsHandler)
	chi.With(r.middlewareJWT.JwtMiddleware).Put("/api/news/{newsId}", r.newsHandler.UpdateNewsHandler)
	chi.With(r.middlewareJWT.JwtMiddleware).Delete("/api/news/{newsId}", r.newsHandler.DeleteNewsHandler)
	chi.Get("/api/news/{newsId}", r.newsHandler.GetNewsDetailHandler)
	chi.Get("/api/news", r.newsHandler.GetNewsListHandler)
	return chi
}
