package repository

import (
	"context"
	"database/sql"

	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/lib"
	"github.com/Arceister/ice-house-news/repository"
	errorUtils "github.com/Arceister/ice-house-news/utils/error"
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

func (u UsersRepository) GetOneUserRepository(id string) (entity.User, errorUtils.IErrorMessage) {
	userData := entity.User{}

	stmt, err := u.db.DB.PrepareContext(context.Background(),
		"SELECT id, email, password, name, bio, web, picture FROM users WHERE id = $1",
	)
	if err != nil {
		return entity.User{}, errorUtils.NewInternalServerError(err.Error())
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
		if err == sql.ErrNoRows {
			return entity.User{}, errorUtils.NewNotFoundError("user not found")
		}
		return entity.User{}, errorUtils.NewInternalServerError(err.Error())
	}

	return userData, nil
}

func (u UsersRepository) GetUserByEmailRepository(email string) (entity.User, errorUtils.IErrorMessage) {
	userData := entity.User{}

	stmt, err := u.db.DB.PrepareContext(context.Background(),
		"SELECT id, email, password, name, bio, web, picture FROM users WHERE email = $1",
	)
	if err != nil {
		return entity.User{}, errorUtils.NewInternalServerError(err.Error())
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
		return entity.User{}, errorUtils.NewInternalServerError(err.Error())
	}

	return userData, nil
}

func (u UsersRepository) CreateUserRepository(id uuid.UUID, userData entity.User) errorUtils.IErrorMessage {
	stmt, err := u.db.DB.PrepareContext(context.Background(),
		"INSERT INTO users VALUES($1, $2, $3, $4, $5, $6, $7)",
	)
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
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
		return errorUtils.NewInternalServerError(err.Error())
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	if rows != 1 {
		return errorUtils.NewUnprocessableEntityError("user not created")
	}

	return nil
}

func (u UsersRepository) UpdateUserRepository(id string, userData entity.User) errorUtils.IErrorMessage {
	stmt, err := u.db.DB.PrepareContext(context.Background(),
		"UPDATE users SET email = $1, password = $2, name = $3, bio = $4, web = $5, picture = $6 WHERE id = $7",
	)
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
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
		return errorUtils.NewInternalServerError(err.Error())
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	if rows != 1 {
		return errorUtils.NewUnprocessableEntityError("user not updated")
	}

	return nil
}

func (u UsersRepository) DeleteUserRepository(id string) errorUtils.IErrorMessage {
	stmt, err := u.db.DB.PrepareContext(context.Background(),
		"DELETE FROM users WHERE id = $1",
	)
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	commandTag, err := stmt.ExecContext(context.Background(), id)

	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	rows, err := commandTag.RowsAffected()
	if err != nil {
		return errorUtils.NewInternalServerError(err.Error())
	}

	if rows != 1 {
		return errorUtils.NewUnprocessableEntityError("user not deleted")
	}

	return nil
}
