package service

import (
	"github.com/Arceister/ice-house-news/entity"
	"github.com/Arceister/ice-house-news/usecase"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	usecase usecase.UsersUsecase
}

func NewUsersService(usecase usecase.UsersUsecase) UsersService {
	return UsersService{
		usecase: usecase,
	}
}

func (s UsersService) GetOneUserService(id string) (entity.User, error) {
	userData := entity.User{}
	var uid string

	err := s.usecase.GetOneUserUsecase(id).Scan(&uid,
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
	userData.Id = uuid.Must(uuid.Parse(uid))

	return userData, err
}

func (s UsersService) CreateUserService(userData entity.User) (pgconn.CommandTag, error) {
	var uniqueUserId string = uuid.Must(uuid.NewRandom()).String()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)
	userData.Password = string(hashedPassword)

	return s.usecase.CreateUserUsecase(uniqueUserId, userData)
}
