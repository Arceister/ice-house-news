package repository

import (
	"context"
	"errors"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/google/uuid"
)

type CategoriesRepository struct {
	db lib.DB
}

func NewCategoriesRepository(db lib.DB) repository.ICategoriesRepository {
	return CategoriesRepository{
		db: db,
	}
}

func (r CategoriesRepository) GetAllNewsCategoryRepository() ([]entity.Categories, error) {
	var NewsCategories []entity.Categories

	stmt, err := r.db.DB.PrepareContext(context.Background(),
		"SELECT id, name FROM categories",
	)
	if err != nil {
		return NewsCategories, err
	}

	rows, err := stmt.QueryContext(context.Background())

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var category entity.Categories
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}

		NewsCategories = append(NewsCategories, category)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return NewsCategories, nil
}

func (r CategoriesRepository) CreateCategoryRepository(categoryData entity.Categories) error {
	stmt, err := r.db.DB.PrepareContext(context.Background(), "INSERT INTO categories(id, name) VALUES($1, $2)")
	if err != nil {
		return err
	}

	commandTag, err := stmt.Exec(context.Background(),
		categoryData.Id,
		categoryData.Name,
	)

	if err != nil {
		return err
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("category not created")
	}

	return nil
}

func (r CategoriesRepository) CreateAndReturnCategoryRepository(category entity.Categories) (uuid.UUID, error) {
	var returnedCategoryId uuid.UUID

	stmt, err := r.db.DB.PrepareContext(context.Background(), "INSERT INTO categories(id, name) VALUES($1, $2) RETURNING id")
	if err != nil {
		return uuid.Nil, err
	}

	err = stmt.QueryRowContext(context.Background(),
		category.Id,
		category.Name).Scan(&returnedCategoryId)

	if err != nil {
		return uuid.Nil, err
	}

	return returnedCategoryId, nil
}

func (r CategoriesRepository) GetCategoryByNameRepository(categoryName string) (entity.Categories, error) {
	var CategoryDetails entity.Categories

	stmt, err := r.db.DB.PrepareContext(context.Background(), "SELECT id, name FROM categories WHERE name = $1")
	if err != nil {
		return CategoryDetails, err
	}

	err = stmt.QueryRowContext(context.Background(),
		categoryName).Scan(
		&CategoryDetails.Id,
		&CategoryDetails.Name,
	)
	if err != nil {
		return entity.Categories{}, err
	}

	return CategoryDetails, nil
}
