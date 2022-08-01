package service

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/Arceister/ice-house-news/service"
	"github.com/google/uuid"
)

type CommentService struct {
	newsRepository    repository.INewsRepository
	commentRepository repository.ICommentRepository
}

func NewCommentService(
	newsRepository repository.INewsRepository,
	commentRepository repository.ICommentRepository,
) service.ICommentService {
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

func (s CommentService) InsertCommentService(
	commentRequest entity.CommentInsertRequest,
	newsId string,
	userId string,
) error {
	var commentInsert entity.CommentInsert

	newCommentUUID := uuid.Must(uuid.NewRandom())
	newsUUID := uuid.MustParse(newsId)
	userUUID := uuid.MustParse(userId)

	commentInsert.Id = newCommentUUID
	commentInsert.NewsId = newsUUID
	commentInsert.UserId = userUUID
	commentInsert.CommentInsertRequest = commentRequest

	err := s.commentRepository.InsertCommentRepository(commentInsert)

	if err != nil {
		return err
	}

	return nil
}
