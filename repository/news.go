package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
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

	rows, err := r.db.DB.Query(context.Background(),
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
	`)

	if err != nil {
		return NewsListOutput, err
	}

	defer rows.Close()

	for rows.Next() {
		var news entity.NewsListOutput
		var category entity.NewsCategory
		var counter entity.NewsCounter

		err = rows.Scan(
			&news.Id, &news.Title, &news.SlugUrl, &news.CoverImage, &news.AdditionalImages, &news.Nsfw,
			&category.Id, &category.Name,
			&counter.Upvote, &counter.Downvote, &counter.Comment, &counter.View,
			&news.CreatedAt,
		)

		if err != nil {
			return NewsListOutput, err
		}

		news.Category = category
		news.Counter = counter

		NewsListOutput = append(NewsListOutput, news)
	}

	return NewsListOutput, err
}

func (r NewsRepository) GetNewsDetailRepository(newsId string) (entity.NewsDetail, error) {
	var NewsDetailOutput entity.NewsDetail
	var category entity.NewsCategory
	var counter entity.NewsCounter
	var author entity.NewsAuthor

	tx, err := r.db.DB.Begin(context.Background())
	if err != nil {
		return entity.NewsDetail{}, err
	}

	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(),
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
	`, newsId).Scan(&NewsDetailOutput.Id, &NewsDetailOutput.Title, &NewsDetailOutput.SlugUrl, &NewsDetailOutput.CoverImage,
		&NewsDetailOutput.AdditionalImages, &NewsDetailOutput.Nsfw,
		&category.Id, &category.Name,
		&author.Id, &author.Name, &author.Picture,
		&counter.Upvote, &counter.Downvote, &counter.Comment, &counter.View,
		&NewsDetailOutput.CreatedAt, &NewsDetailOutput.Content)

	if err != nil {
		return entity.NewsDetail{}, err
	}

	commandTag, err := tx.Exec(context.Background(),
		`
	UPDATE news_counter
	SET view = view + 1
	WHERE news_id = $1;
	`, newsId)

	if err != nil {
		return entity.NewsDetail{}, err
	}

	if commandTag.RowsAffected() != 1 {
		return entity.NewsDetail{}, errors.New("news not found")
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return entity.NewsDetail{}, err
	}

	NewsDetailOutput.Author = author
	NewsDetailOutput.Category = category
	NewsDetailOutput.Counter = counter

	return NewsDetailOutput, err
}

func (r NewsRepository) AddNewNewsRepository(news entity.NewsInsert) error {
	tx, err := r.db.DB.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	commandTag, err := tx.Exec(context.Background(),
		`INSERT INTO news(id, users_id, category_id, title, isi, slug_url, cover_image, nsfw, created_at) 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
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

	if commandTag.RowsAffected() != 1 {
		return errors.New("create news failed")
	}

	for _, additionalImagesInput := range news.AdditionalImages {
		commandTag, err := tx.Exec(context.Background(),
			"INSERT INTO news_additional_images(news_id, image) VALUES ($1, $2)", news.Id, additionalImagesInput)

		if err != nil {
			return err
		}

		if commandTag.RowsAffected() != 1 {
			return errors.New("input additional image failed")
		}
	}

	commandTag, err = tx.Exec(context.Background(),
		"INSERT INTO news_counter(news_id) VALUES ($1)", news.Id)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("insert news counter failed")
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return err
}
