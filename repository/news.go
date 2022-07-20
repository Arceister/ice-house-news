package repository

import (
	"context"

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
