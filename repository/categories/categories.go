package repository

import (
	"context"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/repository"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
	"github.com/google/uuid"
)

var _ repository.ICategoriesRepository = (*CategoriesRepository)(nil)

type CategoriesRepository struct {
	db lib.DB
}

func NewCategoriesRepository(db lib.DB) *CategoriesRepository {
	return &CategoriesRepository{
		db: db,
	}
}

func (r *CategoriesRepository) GetAllNewsCategoryRepository() ([]entity.Categories, errorUtils.IErrorMessage) {
	var NewsCategories []entity.Categories

	stmt, err := r.db.DB.PrepareContext(context.Background(),
		"SELECT id, name FROM categories",
	)
	if err != nil {
		return NewsCategories, errorUtils.NewInternalServerError(err.Error())
	}

	rows, err := stmt.QueryContext(context.Background())

	if err != nil {
		return nil, errorUtils.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var category entity.Categories
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, errorUtils.NewInternalServerError(err.Error())
		}

		NewsCategories = append(NewsCategories, category)
	}

	if rows.Err() != nil {
		return nil, errorUtils.NewInternalServerError(err.Error())
	}

	return NewsCategories, nil
}

func (r *CategoriesRepository) CreateCategoryRepository(categoryData entity.Categories) errorUtils.IErrorMessage {
	stmt, err := r.db.DB.PrepareContext(context.Background(), "INSERT INTO categories(id, name) VALUES($1, $2)")
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	commandTag, err := stmt.Exec(context.Background(),
		categoryData.Id,
		categoryData.Name,
	)

	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	if rows != 1 {
		return errorUtils.NewUnprocessableEntityError("catgory not created")
	}

	return nil
}

func (r *CategoriesRepository) CreateAndReturnCategoryRepository(category entity.Categories) (uuid.UUID, errorUtils.IErrorMessage) {
	var returnedCategoryId uuid.UUID

	stmt, err := r.db.DB.PrepareContext(context.Background(), "INSERT INTO categories(id, name) VALUES($1, $2) RETURNING id")
	if err != nil {
		return uuid.Nil, errorUtils.NewInternalServerError(err.Error())
	}

	err = stmt.QueryRowContext(context.Background(),
		category.Id,
		category.Name).Scan(&returnedCategoryId)

	if err != nil {
		return uuid.Nil, errorUtils.NewInternalServerError(err.Error())
	}

	return returnedCategoryId, nil
}

func (r *CategoriesRepository) GetCategoryByNameRepository(categoryName string) (entity.Categories, errorUtils.IErrorMessage) {
	var CategoryDetails entity.Categories

	stmt, err := r.db.DB.PrepareContext(context.Background(), "SELECT id, name FROM categories WHERE name = $1")
	if err != nil {
		return CategoryDetails, errorUtils.NewInternalServerError(err.Error())
	}

	err = stmt.QueryRowContext(context.Background(),
		categoryName).Scan(
		&CategoryDetails.Id,
		&CategoryDetails.Name,
	)
	if err != nil {
		return entity.Categories{}, errorUtils.NewInternalServerError(err.Error())
	}

	return CategoryDetails, nil
}
