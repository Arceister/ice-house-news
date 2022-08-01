package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type NewsRepository struct {
	db lib.DB
}

func NewNewsRepository(db lib.DB) NewsRepository {
	return NewsRepository{
		db: db,
	}
}

func (r NewsRepository) GetNewsListRepository() ([]entity.NewsListOutput, error) {
	var NewsListOutput []entity.NewsListOutput

	stmt, err := r.db.DB.PrepareContext(context.Background(),
		`
		SELECT n.id, n.title, n.slug_url, n.cover_image, (
				SELECT array_agg(image) FROM news_additional_images WHERE news_id = n.id
				) as additional_images, n.nsfw, 
				c.id, c.name, 
				nc.upvote, nc.downvote, nc.comment, nc.view, 
				n.created_at
		FROM news n
		JOIN categories c ON c.id = n.category_id
		JOIN news_counter nc on n.id = nc.news_id;
	`,
	)
	if err != nil {
		return NewsListOutput, err
	}

	rows, err := stmt.QueryContext(context.Background())

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var news entity.NewsListOutput
		var category entity.NewsCategory
		var counter entity.NewsCounter
		var NewsAdditionalImages sql.NullString

		err = rows.Scan(
			&news.Id, &news.Title, &news.SlugUrl, &news.CoverImage, &NewsAdditionalImages, &news.Nsfw,
			&category.Id, &category.Name,
			&counter.Upvote, &counter.Downvote, &counter.Comment, &counter.View,
			&news.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		if len(NewsAdditionalImages.String) != 0 {
			newsAdditionalImagesExtract := NewsAdditionalImages.String[1 : len(NewsAdditionalImages.String)-1]
			newsAdditionalImages := strings.Split(newsAdditionalImagesExtract, ",")
			news.AdditionalImages = newsAdditionalImages
		}

		news.Category = category
		news.Counter = counter

		NewsListOutput = append(NewsListOutput, news)
	}

	return NewsListOutput, nil
}

func (r NewsRepository) GetNewsDetailRepository(newsId string) (entity.NewsDetail, error) {
	var NewsDetailOutput entity.NewsDetail
	var category entity.NewsCategory
	var counter entity.NewsCounter
	var author entity.NewsAuthor
	var NewsAdditionalImages sql.NullString

	tx, err := r.db.DB.Begin()
	if err != nil {
		return entity.NewsDetail{}, err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(context.Background(),
		`
		SELECT n.id, n.title, n.slug_url, n.cover_image, (
			SELECT array_agg(image) FROM news_additional_images WHERE news_id = n.id
			) as additional_images, n.nsfw,
			c.id, c.name,
			u.id, u.name, u.picture,
			nc.upvote, nc.downvote, nc.comment, nc.view,
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

	if len(NewsAdditionalImages.String) != 0 {
		newsAdditionalImagesExtract := NewsAdditionalImages.String[1 : len(NewsAdditionalImages.String)-1]
		newsAdditionalImages := strings.Split(newsAdditionalImagesExtract, ",")
		NewsDetailOutput.AdditionalImages = newsAdditionalImages
	}

	if err != nil {
		return entity.NewsDetail{}, err
	}

	stmt, err = tx.PrepareContext(
		context.Background(),
		`
		UPDATE news_counter
		SET view = view + 1
		WHERE news_id = $1;
		`)
	if err != nil {
		return entity.NewsDetail{}, err
	}

	defer stmt.Close()

	commandTag, err := stmt.ExecContext(context.Background(),
		newsId,
	)

	if err != nil {
		return entity.NewsDetail{}, err
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return entity.NewsDetail{}, err
	}

	if rows != 1 {
		return entity.NewsDetail{}, errors.New("view not updated")
	}

	err = tx.Commit()
	if err != nil {
		return entity.NewsDetail{}, err
	}

	NewsDetailOutput.Author = author
	NewsDetailOutput.Category = category
	NewsDetailOutput.Counter = counter

	return NewsDetailOutput, nil
}

func (r NewsRepository) GetNewsUserRepository(newsId string) (string, error) {
	var newsUUID string
	stmt, err := r.db.DB.PrepareContext(context.Background(),
		`SELECT users_id FROM news WHERE id = $1`,
	)
	if err != nil {
		return newsUUID, nil
	}

	err = stmt.QueryRow(context.Background(),
		newsId).Scan(&newsUUID)

	if err != nil {
		return newsUUID, err
	}

	return newsUUID, nil
}

func (r NewsRepository) AddNewNewsRepository(news entity.NewsInsert) error {
	tx, err := r.db.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(context.Background(),
		`INSERT INTO news(id, users_id, category_id, title, isi, slug_url, cover_image, nsfw, created_at) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
	)
	if err != nil {
		return err
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
		return err
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("news not created")
	}

	for _, additionalImagesInput := range news.AdditionalImages {
		stmt, err := tx.PrepareContext(context.Background(), "INSERT INTO news_additional_images(news_id, image) VALUES ($1, $2)")
		if err != nil {
			return err
		}

		defer stmt.Close()

		commandTag, err := stmt.ExecContext(context.Background(),
			news.Id, additionalImagesInput)

		if err != nil {
			return err
		}

		rows, err := commandTag.RowsAffected()
		if err != nil {
			return err
		}

		if rows != 1 {
			return errors.New("input additional image failed")
		}
	}
	stmt, err = tx.PrepareContext(context.Background(), "INSERT INTO news_counter(news_id) VALUES ($1)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	commandTag, err = stmt.ExecContext(context.Background(),
		news.Id)

	if err != nil {
		return err
	}

	rowsAffected, err := commandTag.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("input news counter failed")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r NewsRepository) UpdateNewsRepository(news entity.NewsInsert) error {
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
		return err
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
		return err
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("news not updated")
	}

	return err
}

func (r NewsRepository) DeleteNewsRepository(newsId string) error {
	commandTag, err := r.db.DB.Exec(
		"DELETE FROM news WHERE id = $1", newsId)

	if err != nil {
		return err
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("news not deleted")
	}

	return nil
}
