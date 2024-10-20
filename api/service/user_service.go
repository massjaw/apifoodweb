package service

import (
	repository "apifoodweb/api/repository"
	"apifoodweb/pkg/database/model"
	"database/sql"
	"errors"
)

type UserService interface {
	LoginUser(username string) (model.Users, error)
	RegisterUser(userData *model.Users) (*model.Users, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u userService) LoginUser(username string) (model.Users, error) {

	userDB, errGetUser := u.userRepo.FindUserByUsername(username)
	if errGetUser != nil {
		if errGetUser == sql.ErrNoRows {
			return model.Users{}, errors.New("user not found")
		}
		return model.Users{}, errGetUser
	}

	return userDB, nil
}

func (u *userService) RegisterUser(userData *model.Users) (*model.Users, error) {

	if errCreateUser := u.userRepo.CreateUser(userData); errCreateUser != nil {
		return userData, errCreateUser
	}

	return userData, nil
}
