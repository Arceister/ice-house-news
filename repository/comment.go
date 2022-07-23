package repository

import "github.com/Arceister/ice-house-news/lib"

type CommentRepository struct {
	db lib.DB
}

func NewCommentRepository(db lib.DB) CommentRepository {
	return CommentRepository{
		db: db,
	}
}
