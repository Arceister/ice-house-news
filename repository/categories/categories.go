package repository

import (
	"errors"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/google/uuid"
)

type CategoriesRepository struct {
	db lib.DB
}

func NewCategoriesRepository(db lib.DB) CategoriesRepository {
	return CategoriesRepository{
		db: db,
	}
}

func (r CategoriesRepository) GetAllNewsCategoryRepository() ([]entity.Categories, error) {
	var NewsCategories []entity.Categories

	rows, err := r.db.DB.Query("SELECT id, name FROM categories")

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
	commandTag, err := r.db.DB.Exec("INSERT INTO categories(id, name) VALUES($1, $2)",
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

	err := r.db.DB.QueryRow("INSERT INTO categories(id, name) VALUES($1, $2) RETURNING id",
		category.Id,
		category.Name).Scan(&returnedCategoryId)

	if err != nil {
		return uuid.Nil, err
	}

	return returnedCategoryId, nil
}

func (r CategoriesRepository) GetCategoryByNameRepository(categoryName string) (entity.Categories, error) {
	var CategoryDetails entity.Categories

	err := r.db.DB.QueryRow("SELECT id, name FROM categories WHERE name = $1",
		categoryName).Scan(
		&CategoryDetails.Id,
		&CategoryDetails.Name,
	)
	if err != nil {
		return entity.Categories{}, err
	}

	return CategoryDetails, nil
}