package repository

import "github.com/Arceister/ice-house-news/lib"

type NewsRepository struct {
	db lib.DB
}

func NewNewsRepository(db lib.DB) NewsRepository {
	return NewsRepository{
		db: db,
	}
}
