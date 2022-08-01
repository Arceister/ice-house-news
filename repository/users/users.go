package repository

import (
	"context"
	"errors"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/repository"
	"github.com/google/uuid"
)

type UsersRepository struct {
	db lib.DB
}

func NewUsersRepository(db lib.DB) repository.IUsersRepository {
	return UsersRepository{
		db: db,
	}
}

func (u UsersRepository) GetOneUserRepository(id string) (entity.User, error) {
	userData := entity.User{}

	stmt, err := u.db.DB.PrepareContext(context.Background(),
		"SELECT id, email, password, name, bio, web, picture FROM users WHERE id = $1",
	)
	if err != nil {
		return entity.User{}, err
	}

	err = stmt.QueryRowContext(context.Background(), id).
		Scan(
			&userData.Id,
			&userData.Email,
			&userData.Password,
			&userData.Name,
			&userData.Bio,
			&userData.Web,
			&userData.Picture,
		)

	if err != nil {
		return entity.User{}, err
	}

	return userData, nil
}

func (u UsersRepository) GetUserByEmailRepository(email string) (entity.User, error) {
	userData := entity.User{}

	stmt, err := u.db.DB.PrepareContext(context.Background(),
		"SELECT id, email, password, name, bio, web, picture FROM users WHERE email = $1",
	)
	if err != nil {
		return entity.User{}, err
	}

	err = stmt.QueryRowContext(context.Background(), email).
		Scan(&userData.Id,
			&userData.Email,
			&userData.Password,
			&userData.Name,
			&userData.Bio,
			&userData.Web,
			&userData.Picture,
		)

	if err != nil {
		return entity.User{}, err
	}

	return userData, nil
}

func (u UsersRepository) CreateUserRepository(id uuid.UUID, userData entity.User) error {
	stmt, err := u.db.DB.PrepareContext(context.Background(),
		"INSERT INTO users VALUES($1, $2, $3, $4, $5, $6, $7)",
	)
	if err != nil {
		return err
	}

	commandTag, err := stmt.ExecContext(context.Background(),
		id,
		userData.Email,
		userData.Password,
		userData.Name,
		userData.Bio,
		userData.Web,
		userData.Picture,
	)
	if err != nil {
		return err
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("user not created")
	}

	return nil
}

func (u UsersRepository) UpdateUserRepository(id string, userData entity.User) error {
	stmt, err := u.db.DB.PrepareContext(context.Background(),
		"UPDATE users SET email = $1, password = $2, name = $3, bio = $4, web = $5, picture = $6 WHERE id = $7",
	)
	if err != nil {
		return err
	}

	commandTag, err := stmt.ExecContext(context.Background(),
		userData.Email,
		userData.Password,
		userData.Name,
		userData.Bio,
		userData.Web,
		userData.Picture,
		id,
	)

	if err != nil {
		return err
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("user not updated")
	}

	return nil
}

func (u UsersRepository) DeleteUserRepository(id string) error {
	stmt, err := u.db.DB.PrepareContext(context.Background(),
		"DELETE FROM users WHERE id = $1",
	)
	if err != nil {
		return err
	}

	commandTag, err := stmt.ExecContext(context.Background(), id)

	if err != nil {
		return err
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("user not deleted")
	}

	return nil
}
