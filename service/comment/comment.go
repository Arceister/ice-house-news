package service

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/service"
	"github.com/google/uuid"

	commentRepository "github.com/Arceister/ice-house-news/repository/comment"
	newsRepository "github.com/Arceister/ice-house-news/repository/news"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
)

var _ service.ICommentService = (*CommentService)(nil)

type CommentService struct {
	newsRepository    newsRepository.NewsRepository
	commentRepository commentRepository.CommentRepository
}

func NewCommentService(
	newsRepository newsRepository.NewsRepository,
	commentRepository commentRepository.CommentRepository,
) CommentService {
	return CommentService{
		newsRepository:    newsRepository,
		commentRepository: commentRepository,
	}
}

func (s *CommentService) GetCommentsOnNewsService(newsId string) ([]entity.Comment, errorUtils.IErrorMessage) {
	newsComment, err := s.commentRepository.GetCommentsOnNewsRepository(newsId)
	if err != nil {
		return nil, err
	}

	return newsComment, nil
}

func (s *CommentService) InsertCommentService(
	commentRequest entity.CommentInsertRequest,
	newsId string,
	userId string,
) errorUtils.IErrorMessage {
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
