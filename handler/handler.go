package handler

import "net/http"

type IUsersHandler interface {
	GetOwnProfile(w http.ResponseWriter, r *http.Request)
	GetOneUserHandler(w http.ResponseWriter, r *http.Request)
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
	UpdateUserHandler(w http.ResponseWriter, r *http.Request)
	DeleteUserHandler(w http.ResponseWriter, r *http.Request)
}

type IAuthHandler interface {
	UserSignInHandler(w http.ResponseWriter, r *http.Request)
	ExtendTokenHandler(w http.ResponseWriter, r *http.Request)
}

type ICategoriesHandler interface {
	GetAllNewsCategoryHandler(w http.ResponseWriter, r *http.Request)
}

type INewsHandler interface {
	GetNewsListHandler(w http.ResponseWriter, r *http.Request)
	GetNewsDetailHandler(w http.ResponseWriter, r *http.Request)
	AddNewNewsHandler(w http.ResponseWriter, r *http.Request)
	UpdateNewsHandler(w http.ResponseWriter, r *http.Request)
	DeleteNewsHandler(w http.ResponseWriter, r *http.Request)
}

type ICommentHandler interface {
	GetCommentsOnNewsHandler(w http.ResponseWriter, r *http.Request)
	InsertCommentHandler(w http.ResponseWriter, r *http.Request)
}
