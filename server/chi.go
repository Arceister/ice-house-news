package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Chi *chi.Mux
}

func NewServer() Server {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	return Server{
		Chi: router,
	}
}
