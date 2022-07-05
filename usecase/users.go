package usecase

import (
	"context"

	"github.com/Arceister/ice-house-news/lib"
	"github.com/jackc/pgx/v4"
)

type UsersUsecase struct {
	db lib.DB
}

func NewUsersUsecase(db lib.DB) UsersUsecase {
	return UsersUsecase{
		db: db,
	}
}

func (u UsersUsecase) GetOneUserUsecase(uuid string) pgx.Row {
	return u.db.DB.QueryRow(context.Background(), "SELECT * FROM users WHERE id = $1", uuid)
}
