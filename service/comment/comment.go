package service

import (
	"github.com/Arceister/ice-house-news/entity"
	comment "github.com/Arceister/ice-house-news/repository/comment"
	news "github.com/Arceister/ice-house-news/repository/news"
	"github.com/google/uuid"
)

type CommentService struct {
	newsRepository    news.NewsRepository
	commentRepository comment.CommentRepository
}

func NewCommentService(
	newsRepository news.NewsRepository,
	commentRepository comment.CommentRepository,
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
