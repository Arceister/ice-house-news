package service

import "github.com/Arceister/ice-house-news/repository"

type CommentService struct {
	newsRepository    repository.NewsRepository
	commentRepository repository.CommentRepository
}

func NewCommentService(
	newsRepository repository.NewsRepository,
	commentRepository repository.CommentRepository,
) CommentService {
	return CommentService{
		newsRepository:    newsRepository,
		commentRepository: commentRepository,
	}
}
