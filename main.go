package main

import (
	"net/http"

	"github.com/Arceister/ice-house-news/controller"
	"github.com/Arceister/ice-house-news/lib"
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

	usersUsecase := usecase.NewUsersUsecase(database)
	usersService := service.NewUsersService(usersUsecase)
	usersController := controller.NewUsersController(usersService)
	usersRouter := router.NewUsersRouter(chiRouter, usersController)
	usersRouter.Setup(chiRouter.Chi)

	http.ListenAndServe(app.Port, chiRouter.Chi)
}
