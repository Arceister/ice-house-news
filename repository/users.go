package repository

import (
	"context"
	"errors"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/google/uuid"
)

type UsersRepository struct {
	db lib.DB
}

func NewUsersRepository(db lib.DB) UsersRepository {
	return UsersRepository{
		db: db,
	}
}

func (u UsersRepository) GetOneUserRepository(id string) (entity.User, error) {
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

func (u UsersRepository) GetUserByEmailRepository(email string) (entity.User, error) {
	userData := entity.User{}

	err := u.db.DB.QueryRow(context.Background(),
		"SELECT id, email, password, name, bio, web, picture FROM users WHERE email = $1",
		email).Scan(&userData.Id,
		&userData.Email,
		&userData.Password,
		&userData.Name,
		&userData.Bio,
		&userData.Web,
		&userData.Picture,
	)

	return userData, err
}

func (u UsersRepository) CreateUserRepository(id uuid.UUID, userData entity.User) error {
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

func (u UsersRepository) UpdateUserRepository(id string, userData entity.User) error {
	commandTag, err := u.db.DB.Exec(context.Background(),
		"UPDATE users SET email = $1, password = $2, name = $3, bio = $4, web = $5, picture = $6 WHERE id = $7",
		userData.Email,
		userData.Password,
		userData.Name,
		userData.Bio,
		userData.Web,
		userData.Picture,
		id,
	)

	if commandTag.RowsAffected() != 1 {
		return errors.New("user not found")
	}

	return err
}

func (u UsersRepository) DeleteUserRepository(id string) error {
	commandTag, err := u.db.DB.Exec(context.Background(), "DELETE FROM users WHERE id = $1", id)

	if commandTag.RowsAffected() != 1 {
		return errors.New("user not found")
	}

	return err
}
