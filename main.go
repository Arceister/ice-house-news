package main

import (
	"net/http"

	"github.com/Arceister/ice-house-news/handler"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/Arceister/ice-house-news/router"
	server "github.com/Arceister/ice-house-news/server"
	"github.com/Arceister/ice-house-news/service"
)

func main() {
	chiRouter := server.NewServer()

	config := lib.NewConfig()
	app := config.App
	db := config.Database

	database := lib.NewDB(db)

	jwtMiddleware := middleware.NewMiddlewareJWT(app)

	usersRepository := repository.NewUsersRepository(database)
	newsRepository := repository.NewNewsRepository(database)
  categoriesRepository := repository.NewCategoriesRepository(database)

	usersService := service.NewUsersService(usersRepository, jwtMiddleware)
	newsService := service.NewNewsService(newsRepository, usersRepository)
  categoriesService := service.NewCategoriesService(categoriesRepository)

	usersHandler := handler.NewUsersHandler(usersService)
	authHandler := handler.NewAuthHandler(usersService)
	newsHandler := handler.NewNewsHandler(newsService)
  categoriesHandler := handler.NewCategoriesHandler(categoriesService)

	usersRouter := router.NewUsersRouter(chiRouter, jwtMiddleware, usersHandler)
	authRouter := router.NewAuthRouter(chiRouter, authHandler)
	newsRouter := router.NewNewsRouter(chiRouter, jwtMiddleware, newsHandler)
  categoriesRouter := router.NewCategoriesRouter(chiRouter, categoriesHandler)

	usersRouter.Setup(chiRouter.Chi)
	authRouter.Setup(chiRouter.Chi)
	newsRouter.Setup(chiRouter.Chi)
  categoriesRouter.Setup(chiRouter.Chi)

	http.ListenAndServe(app.Port, chiRouter.Chi)
}
