package usecase

import (
	"context"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/jackc/pgconn"
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

func (u UsersUsecase) CreateUserUsecase(uuid string, userData entity.User) (pgconn.CommandTag, error) {
	return u.db.DB.Exec(context.Background(),
		"INSERT INTO users VALUES($1, $2, $3, $4, $5, $6, $7)",
		uuid,
		userData.Email,
		userData.Password,
		userData.Name,
		userData.Bio,
		userData.Web,
		userData.Picture,
	)
}

func (u UsersUsecase) UpdateUserUsecase(uuid string, userData entity.User) (pgconn.CommandTag, error) {
	return u.db.DB.Exec(context.Background(),
		"UPDATE users SET username = $1, password = $2, name = $3, bio = $4, web = $5, picture = $6 WHERE id = $7",
		userData.Email,
		userData.Password,
		userData.Name,
		userData.Bio,
		userData.Web,
		userData.Picture,
		uuid,
	)
}

func (u UsersUsecase) DeleteUserUsecase(uuid string) (pgconn.CommandTag, error) {
	return u.db.DB.Exec(context.Background(), "DELETE FROM users WHERE id = $1", uuid)
}
