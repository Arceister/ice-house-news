package utils

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RequestHandler struct {
	Chi *chi.Mux
}

func NewRequestHandler() RequestHandler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	return RequestHandler{
		Chi: router,
	}
}
