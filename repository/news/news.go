package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/repository"

	errorUtils "github.com/Arceister/ice-house-news/utils/error"
)

var _ repository.INewsRepository = (*NewsRepository)(nil)

type NewsRepository struct {
	db lib.DB
}

func NewNewsRepository(db lib.DB) NewsRepository {
	return NewsRepository{
		db: db,
	}
}

func (r *NewsRepository) GetNewsListRepository(scope int, category string) ([]entity.NewsListOutput, errorUtils.IErrorMessage) {
	var NewsListOutput []entity.NewsListOutput
	var scopeQuery string
	var categoryQuery string

	if scope > 0 {
		scopeQuery = fmt.Sprintf("LIMIT %d", scope)
	}

	if len(category) > 0 {
		categoryQuery = fmt.Sprintf("WHERE c.name = '%s'", category)
	}

	query := `
	SELECT n.id, n.title, n.slug_url, n.cover_image, (
			SELECT array_to_json(array_agg(row_to_json(t))) FROM ( SELECT image FROM news_additional_images WHERE news_id = n.id ) t
			) as additional_images, n.nsfw,
			c.id, c.name,
			u.id, u.name, u.picture,
			nc.upvote, nc.downvote, (SELECT COUNT(*) FROM news_comment WHERE news_id = n.id) as comment, nc.view,
			n.created_at
	FROM news n
	JOIN categories c ON c.id = n.category_id
	JOIN news_counter nc on n.id = nc.news_id
	JOIN users u on n.users_id = u.id
` + categoryQuery + scopeQuery

	stmt, err := r.db.DB.PrepareContext(context.Background(),
		query,
	)
	if err != nil {
		return NewsListOutput, errorUtils.NewInternalServerError(err.Error())
	}

	rows, err := stmt.QueryContext(context.Background())

	if err != nil {
		return nil, errorUtils.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	if rows.Err() != nil {
		return nil, errorUtils.NewInternalServerError(rows.Err().Error())
	}

	for rows.Next() {
		var news entity.NewsListOutput
		var category entity.NewsCategory
		var author entity.NewsAuthor
		var counter entity.NewsCounter
		var NewsAdditionalImages sql.NullString

		err = rows.Scan(
			&news.Id, &news.Title, &news.SlugUrl, &news.CoverImage, &NewsAdditionalImages, &news.Nsfw,
			&category.Id, &category.Name,
			&author.Id, &author.Name, &author.Picture,
			&counter.Upvote, &counter.Downvote, &counter.Comment, &counter.View,
			&news.CreatedAt,
		)

		if err != nil {
			return nil, errorUtils.NewInternalServerError(err.Error())
		}

		var additionalImages []string

		if NewsAdditionalImages.Valid {
			var imagesJson []map[string]interface{}
			if err := json.Unmarshal([]byte(NewsAdditionalImages.String), &imagesJson); err != nil {
				return NewsListOutput, errorUtils.NewInternalServerError(err.Error())
			}

			for _, images := range imagesJson {
				additionalImages = append(additionalImages, images["image"].(string))
			}
		}

		news.Category = category
		news.Counter = counter
		news.Author = author
		news.AdditionalImages = additionalImages

		NewsListOutput = append(NewsListOutput, news)
	}

	if NewsListOutput == nil {
		return nil, errorUtils.NewNotFoundError("specified search not found")
	}

	return NewsListOutput, nil
}

func (r *NewsRepository) GetNewsDetailRepository(newsId string) (entity.NewsDetail, errorUtils.IErrorMessage) {
	var NewsDetailOutput entity.NewsDetail
	var category entity.NewsCategory
	var counter entity.NewsCounter
	var author entity.NewsAuthor
	var NewsAdditionalImages sql.NullString

	tx, err := r.db.DB.Begin()
	if err != nil {
		return entity.NewsDetail{}, errorUtils.NewInternalServerError(err.Error())
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(context.Background(),
		`
		SELECT n.id, n.title, n.slug_url, n.cover_image, (
				SELECT array_to_json(array_agg(row_to_json(t))) FROM ( SELECT image FROM news_additional_images WHERE news_id = n.id ) t
				) as additional_images, n.nsfw,
				c.id, c.name,
				u.id, u.name, u.picture,
				nc.upvote, nc.downvote, (SELECT COUNT(*) FROM news_comment WHERE news_id = n.id) as comment, nc.view,
				n.created_at, n.isi
		FROM news n
		JOIN categories c ON c.id = n.category_id
		JOIN news_counter nc on n.id = nc.news_id
		JOIN users u on n.users_id = u.id
		WHERE n.id = $1;
		`,
	)
	if err != nil {
		tx.Rollback()
	}

	defer stmt.Close()

	err = stmt.QueryRowContext(
		context.Background(), newsId).
		Scan(&NewsDetailOutput.Id, &NewsDetailOutput.Title, &NewsDetailOutput.SlugUrl, &NewsDetailOutput.CoverImage,
			&NewsAdditionalImages, &NewsDetailOutput.Nsfw,
			&category.Id, &category.Name,
			&author.Id, &author.Name, &author.Picture,
			&counter.Upvote, &counter.Downvote, &counter.Comment, &counter.View,
			&NewsDetailOutput.CreatedAt, &NewsDetailOutput.Content)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() || strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			return NewsDetailOutput, errorUtils.NewNotFoundError("news not found")
		}
		return NewsDetailOutput, errorUtils.NewInternalServerError(err.Error())
	}

	var additionalImages []string

	if NewsAdditionalImages.Valid {
		var imagesJson []map[string]interface{}
		if err := json.Unmarshal([]byte(NewsAdditionalImages.String), &imagesJson); err != nil {
			return NewsDetailOutput, errorUtils.NewInternalServerError(err.Error())
		}

		for _, images := range imagesJson {
			additionalImages = append(additionalImages, images["image"].(string))
		}
	}

	NewsDetailOutput.AdditionalImages = additionalImages

	if err != nil {
		return entity.NewsDetail{}, errorUtils.NewInternalServerError(err.Error())
	}

	stmt, err = tx.PrepareContext(
		context.Background(),
		`
		UPDATE news_counter
		SET view = view + 1
		WHERE news_id = $1;
		`)
	if err != nil {
		return entity.NewsDetail{}, errorUtils.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	commandTag, err := stmt.ExecContext(context.Background(),
		newsId,
	)

	if err != nil {
		return entity.NewsDetail{}, errorUtils.NewInternalServerError(err.Error())
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return entity.NewsDetail{}, errorUtils.NewInternalServerError(err.Error())
	}

	if rows != 1 {
		return entity.NewsDetail{}, errorUtils.NewUnprocessableEntityError("view not updated")
	}

	err = tx.Commit()
	if err != nil {
		return entity.NewsDetail{}, errorUtils.NewInternalServerError(err.Error())
	}

	NewsDetailOutput.Author = author
	NewsDetailOutput.Category = category
	NewsDetailOutput.Counter = counter

	return NewsDetailOutput, nil
}

func (r *NewsRepository) GetNewsUserRepository(newsId string) (string, errorUtils.IErrorMessage) {
	var newsUUID string
	stmt, err := r.db.DB.PrepareContext(context.Background(),
		`SELECT users_id FROM news WHERE id = $1`,
	)
	if err != nil {
		return newsUUID, errorUtils.NewInternalServerError(err.Error())
	}

	err = stmt.QueryRow(context.Background(),
		newsId).Scan(&newsUUID)

	if err != nil {
		return newsUUID, errorUtils.NewInternalServerError(err.Error())
	}

	return newsUUID, nil
}

func (r *NewsRepository) AddNewNewsRepository(news entity.NewsInsert) errorUtils.IErrorMessage {
	tx, err := r.db.DB.Begin()
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(context.Background(),
		`INSERT INTO news(id, users_id, category_id, title, isi, slug_url, cover_image, nsfw, created_at) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
	)
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	commandTag, err := stmt.ExecContext(context.Background(),
		news.Id,
		news.UserId,
		news.CategoryId,
		news.Title,
		news.Content,
		news.SlugUrl,
		news.CoverImage,
		news.Nsfw,
		time.Now())

	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	if rows != 1 {
		return errorUtils.NewUnprocessableEntityError("news not created")
	}

	for _, additionalImagesInput := range news.AdditionalImages {
		stmt, err := tx.PrepareContext(context.Background(), "INSERT INTO news_additional_images(news_id, image) VALUES ($1, $2)")
		if err != nil {
			return errorUtils.NewInternalServerError(err.Error())
		}

		defer stmt.Close()

		commandTag, err := stmt.ExecContext(context.Background(),
			news.Id, additionalImagesInput)

		if err != nil {
			return errorUtils.NewInternalServerError(err.Error())
		}

		rows, err := commandTag.RowsAffected()
		if err != nil {
			return errorUtils.NewInternalServerError(err.Error())
		}

		if rows != 1 {
			return errorUtils.NewUnprocessableEntityError("input additional image failed")
		}
	}
	stmt, err = tx.PrepareContext(context.Background(), "INSERT INTO news_counter(news_id) VALUES ($1)")
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	commandTag, err = stmt.ExecContext(context.Background(),
		news.Id)

	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	rowsAffected, err := commandTag.RowsAffected()
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	if rowsAffected != 1 {
		return errorUtils.NewUnprocessableEntityError("input news counter failed")
	}

	err = tx.Commit()
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *NewsRepository) UpdateNewsRepository(news entity.NewsInsert) errorUtils.IErrorMessage {
	stmt, err := r.db.DB.PrepareContext(context.Background(),
		`UPDATE news SET
		category_id = $1,
		title = $2,
		isi = $3,
		slug_url = $4,
		cover_image = $5,
		nsfw = $6
		WHERE id = $7
		`,
	)
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	commandTag, err := stmt.ExecContext(
		context.Background(),
		news.CategoryId,
		news.Title,
		news.Content,
		news.SlugUrl,
		news.CoverImage,
		news.Nsfw,
		news.Id)

	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	if rows != 1 {
		return errorUtils.NewUnprocessableEntityError("news not updated")
	}

	return nil
}

func (r *NewsRepository) DeleteNewsRepository(newsId string) errorUtils.IErrorMessage {
	commandTag, err := r.db.DB.Exec(
		"DELETE FROM news WHERE id = $1", newsId)

	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	if rows != 1 {
		return errorUtils.NewUnprocessableEntityError("news not deleted")
	}

	return nil
}
