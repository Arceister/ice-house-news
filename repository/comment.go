package repository

import (
	"context"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
)

type CommentRepository struct {
	db lib.DB
}

func NewCommentRepository(db lib.DB) CommentRepository {
	return CommentRepository{
		db: db,
	}
}

func (r CommentRepository) GetCommentsOnNewsRepository(newsId string) ([]entity.Comment, error) {
	var CommentsList []entity.Comment

	rows, err := r.db.DB.Query(context.Background(),
		`
	SELECT nc.id, nc.description,
				u.id, u.name, u.picture,
				nc.created_at
	FROM news_comment nc
	JOIN users u on u.id = nc.users_id
	WHERE nc.news_id = $1
	`, newsId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var comment entity.Comment
		var commentator entity.Commentator

		err = rows.Scan(
			&comment.Id, &comment.Description,
			&commentator.Id, &commentator.Name, &commentator.Picture,
			&comment.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		comment.User = commentator

		CommentsList = append(CommentsList, comment)
	}

	return CommentsList, nil
}
