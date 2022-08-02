package main

import (
	"net/http"

	_authHandler "github.com/Arceister/ice-house-news/handler/auth"
	_categoriesHandler "github.com/Arceister/ice-house-news/handler/categories"
	_commentHandler "github.com/Arceister/ice-house-news/handler/comment"
	_newsHandler "github.com/Arceister/ice-house-news/handler/news"
	_usersHandler "github.com/Arceister/ice-house-news/handler/users"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/middleware"
	_categoriesRepository "github.com/Arceister/ice-house-news/repository/categories"
	_commentRepository "github.com/Arceister/ice-house-news/repository/comment"
	_newsRepository "github.com/Arceister/ice-house-news/repository/news"
	_usersRepository "github.com/Arceister/ice-house-news/repository/users"
	"github.com/Arceister/ice-house-news/router"
	server "github.com/Arceister/ice-house-news/server"
	_categoriesService "github.com/Arceister/ice-house-news/service/categories"
	_commentService "github.com/Arceister/ice-house-news/service/comment"
	_newsService "github.com/Arceister/ice-house-news/service/news"
	_usersService "github.com/Arceister/ice-house-news/service/users"
)

func main() {
	chiRouter := server.NewServer()

	config := lib.NewConfig()
	app := config.App
	db := config.Database

	database := lib.NewDB(db)

	jwtMiddleware := middleware.NewMiddlewareJWT(app)

	usersRepository := _usersRepository.NewUsersRepository(database)
	newsRepository := _newsRepository.NewNewsRepository(database)
	categoriesRepository := _categoriesRepository.NewCategoriesRepository(database)
	commentRepository := _commentRepository.NewCommentRepository(database)

	usersService := _usersService.NewUsersService(usersRepository, jwtMiddleware)
	newsService := _newsService.NewNewsService(newsRepository, usersRepository, categoriesRepository)
	categoriesService := _categoriesService.NewCategoriesService(categoriesRepository)
	commentService := _commentService.NewCommentService(newsRepository, commentRepository)

	usersHandler := _usersHandler.NewUsersHandler(usersService)
	authHandler := _authHandler.NewAuthHandler(usersService)
	newsHandler := _newsHandler.NewNewsHandler(newsService)
	categoriesHandler := _categoriesHandler.NewCategoriesHandler(categoriesService)
	commentHandler := _commentHandler.NewCommentHandler(commentService)

	usersRouter := router.NewUsersRouter(chiRouter, jwtMiddleware, usersHandler)
	authRouter := router.NewAuthRouter(chiRouter, jwtMiddleware, authHandler)
	newsRouter := router.NewNewsRouter(chiRouter, jwtMiddleware, newsHandler)
	categoriesRouter := router.NewCategoriesRouter(chiRouter, categoriesHandler)
	commentRouter := router.NewCommentRoute(chiRouter, jwtMiddleware, commentHandler)

	usersRouter.Setup(chiRouter.Chi)
	authRouter.Setup(chiRouter.Chi)
	newsRouter.Setup(chiRouter.Chi)
	categoriesRouter.Setup(chiRouter.Chi)
	commentRouter.Setup(chiRouter.Chi)

	http.ListenAndServe(app.Port, chiRouter.Chi)
}
