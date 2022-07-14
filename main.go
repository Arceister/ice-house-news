package main

import (
	"net/http"

	"github.com/Arceister/ice-house-news/handler"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/middleware"
	"github.com/Arceister/ice-house-news/router"
	server "github.com/Arceister/ice-house-news/server"
	"github.com/Arceister/ice-house-news/service"
	"github.com/Arceister/ice-house-news/usecase"
)

func main() {
	chiRouter := server.NewServer()

	config := lib.NewConfig()
	app := config.App
	db := config.Database

	database := lib.NewDB(db)

	jwtMiddleware := middleware.NewMiddlewareJWT(app)

	usersUsecase := usecase.NewUsersUsecase(database)
	usersService := service.NewUsersService(usersUsecase)

	usersHandler := handler.NewUsersHandler(usersService)
	authHandler := handler.NewAuthHandler(usersService)

	usersRouter := router.NewUsersRouter(chiRouter, jwtMiddleware, usersHandler)
	authRouter := router.NewAuthRouter(chiRouter, authHandler)

	usersRouter.Setup(chiRouter.Chi)
	authRouter.Setup(chiRouter.Chi)

	http.ListenAndServe(app.Port, chiRouter.Chi)
}
