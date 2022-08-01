package repository

import (
	"context"
	"errors"
	"time"

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

	stmt, err := r.db.DB.PrepareContext(context.Background(),
		`
		SELECT nc.id, nc.description,
					u.id, u.name, u.picture,
					nc.created_at
		FROM news_comment nc
		JOIN users u on u.id = nc.users_id
		WHERE nc.news_id = $1
		`,
	)
	if err != nil {
		return CommentsList, err
	}

	rows, err := stmt.QueryContext(context.Background(),
		newsId)

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

func (r CommentRepository) InsertCommentRepository(commentDetails entity.CommentInsert) error {
	stmt, err := r.db.DB.PrepareContext(context.Background(),
		`
		INSERT INTO news_comment(id, news_id, users_id, description, created_at) 
		VALUES ($1, $2, $3, $4, $5)
		`,
	)
	if err != nil {
		return err
	}

	commandTag, err := stmt.ExecContext(context.Background(),
		commentDetails.Id, commentDetails.NewsId, commentDetails.UserId, commentDetails.Description, time.Now())

	if err != nil {
		return err
	}

	rowsAffected, err := commandTag.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("user insert failed")
	}

	return nil
}
