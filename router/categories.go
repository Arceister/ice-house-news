package router

import (
	"github.com/Arceister/ice-house-news/server"
	"github.com/go-chi/chi/v5"

	categoriesHandler "github.com/Arceister/ice-house-news/handler/categories"
)

type CategoriesRoute struct {
	server            server.Server
	categoriesHandler categoriesHandler.CategoriesHandler
}

func NewCategoriesRouter(
	server server.Server,
	categoriesHandler categoriesHandler.CategoriesHandler,
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
