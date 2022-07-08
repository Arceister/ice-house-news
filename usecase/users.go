package usecase

import (
	"context"
	"errors"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

type UsersUsecase struct {
	db lib.DB
}

func NewUsersUsecase(db lib.DB) UsersUsecase {
	return UsersUsecase{
		db: db,
	}
}

func (u UsersUsecase) GetOneUserUsecase(id string) (entity.User, error) {
	userData := entity.User{}

	err := u.db.DB.QueryRow(context.Background(),
		"SELECT id, email, password, name, bio, web, picture FROM users WHERE id = $1",
		id).Scan(&userData.Id,
		&userData.Email,
		&userData.Password,
		&userData.Name,
		&userData.Bio,
		&userData.Web,
		&userData.Picture,
	)

	return userData, err
}

func (u UsersUsecase) CreateUserUsecase(id uuid.UUID, userData entity.User) error {
	commandTag, err := u.db.DB.Exec(context.Background(),
		"INSERT INTO users VALUES($1, $2, $3, $4, $5, $6, $7)",
		id,
		userData.Email,
		userData.Password,
		userData.Name,
		userData.Bio,
		userData.Web,
		userData.Picture,
	)

	if commandTag.RowsAffected() != 1 {
		return errors.New("user not created")
	}

	return err
}

func (u UsersUsecase) UpdateUserUsecase(id string, userData entity.User) (pgconn.CommandTag, error) {
	return u.db.DB.Exec(context.Background(),
		"UPDATE users SET email = $1, password = $2, name = $3, bio = $4, web = $5, picture = $6 WHERE id = $7",
		userData.Email,
		userData.Password,
		userData.Name,
		userData.Bio,
		userData.Web,
		userData.Picture,
		id,
	)
}

func (u UsersUsecase) DeleteUserUsecase(uuid string) (pgconn.CommandTag, error) {
	return u.db.DB.Exec(context.Background(), "DELETE FROM users WHERE id = $1", uuid)
}
