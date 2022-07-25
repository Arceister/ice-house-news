package service

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/repository"
)

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

func (s CommentService) GetCommentsOnNewsService(newsId string) ([]entity.Comment, error) {
	newsComment, err := s.commentRepository.GetCommentsOnNewsRepository(newsId)
	if err != nil {
		return nil, err
	}

	return newsComment, nil
}
