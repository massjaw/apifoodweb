package service

import (
	"apifoodweb/api/dto/req"
	"apifoodweb/api/dto/resp"
	"apifoodweb/pkg/database/model"
	repository "apifoodweb/pkg/database/repos"
	"apifoodweb/pkg/util"
	"database/sql"
	"errors"
)

type UserService interface {
	LoginUser(loginUser *req.UsersRequestAPI) (*resp.CommonResponse, error)
	RegisterUser(regisUser *req.UsersRequestAPI) (*resp.CommonResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u userService) LoginUser(loginUser *req.UsersRequestAPI) (*resp.CommonResponse, error) {

	errHash := util.PasswordEncrypt(loginUser)
	if errHash != nil {
		return resp.SomethingWentWrong(), errors.New("failed to hash password: " + errHash.Error())
	}

	userDB, errGetUser := u.userRepo.FindUserByUsername(loginUser.Username)
	if errGetUser != nil {
		if errGetUser == sql.ErrNoRows {
			return resp.UserNotFound(), errors.New("user not found")
		}
		return resp.SomethingWentWrong(), errGetUser
	}

	if userDB.HashedPassword != loginUser.Password {
		return resp.WrongPassword(), errors.New("wrong password")
	}

	return resp.StatusOK("login success", ""), nil
}

func (u *userService) RegisterUser(regisUser *req.UsersRequestAPI) (*resp.CommonResponse, error) {

	if regisUser.Password != regisUser.RePassword {
		return resp.PasswordNotMatch(), errors.New("password not match")
	}

	errHash := util.PasswordEncrypt(regisUser)
	if errHash != nil {
		return resp.SomethingWentWrong(), errors.New("failed to hash password")
	}

	userData := model.Users{
		Username:       regisUser.Username,
		Email:          regisUser.Email,
		HashedPassword: regisUser.Password,
	}

	if errCreateUser := u.userRepo.CreateUser(&userData); errCreateUser != nil {
		return resp.SomethingWentWrong(), errCreateUser
	}

	return resp.Created("user created", userData), nil
}
