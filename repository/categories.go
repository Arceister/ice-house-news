package repository

import (
	"context"
	"errors"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
)

type CategoriesRepository struct {
	db lib.DB
}

func NewCategoriesRepository(db lib.DB) CategoriesRepository {
	return CategoriesRepository{
		db: db,
	}
}

func (r CategoriesRepository) CreateCategoryRepository(categoryData entity.Categories) error {
	commandTag, err := r.db.DB.Exec(context.Background(),
		"INSERT INTO categories(id, name) VALUES($1, $2)",
		categoryData.Id,
		categoryData.Name,
	)

	if commandTag.RowsAffected() != 1 {
		return errors.New("category not created")
	}

	return err
}
