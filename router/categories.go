package router

import (
	"github.com/Arceister/ice-house-news/handler"
	"github.com/Arceister/ice-house-news/server"
	"github.com/go-chi/chi/v5"
)

type CategoriesRoute struct {
	server            server.Server
	categoriesHandler handler.ICategoriesHandler
}

func NewCategoriesRouter(
	server server.Server,
	categoriesHandler handler.ICategoriesHandler,
) CategoriesRoute {
	return CategoriesRoute{
		server:            server,
		categoriesHandler: categoriesHandler,
	}
}

func (r CategoriesRoute) Setup(chi *chi.Mux) *chi.Mux {
	chi.Get("/api/news/category", r.categoriesHandler.GetAllNewsCategoryHandler)
	return chi
}
