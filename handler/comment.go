package handler

import "github.com/Arceister/ice-house-news/service"

type CommentHandler struct {
	service service.CommentService
}

func NewCommentHandler(service service.CommentService) CommentHandler {
	return CommentHandler{
		service: service,
	}
}
