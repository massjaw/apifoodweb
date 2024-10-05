package service

import (
	"apifoodweb/api/dto/req"
	"apifoodweb/api/dto/resp"
	connection "apifoodweb/pkg/database"
	"apifoodweb/pkg/database/model"
	repository "apifoodweb/pkg/database/repos"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(loginUser req.UsersRequestAPI) (*resp.CommonResponse, error) {

	errHash := passwordEncrypt(&loginUser)
	if errHash != nil {
		return resp.SomethingWentWrong(), errors.New("failed to hash password: " + errHash.Error())
	}

	db, errDb := connection.GetConnectionGormApiFoodApp()
	if errDb != nil {
		return resp.FailedConnectDatabase("apifoodapp"), errors.New("failed to connect database: " + errDb.Error())
	}

	userDB, errGetUser := repository.NewUserRepository(db).FindUserByUsername(loginUser.Username)
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

func RegisterUser(regisUser req.UsersRequestAPI) (*resp.CommonResponse, error) {

	if regisUser.Password != regisUser.RePassword {
		return resp.PasswordNotMatch(), errors.New("password not match")
	}

	errHash := passwordEncrypt(&regisUser)
	if errHash != nil {
		return resp.SomethingWentWrong(), errors.New("failed to hash password")
	}

	userData := model.Users{
		Username:       regisUser.Username,
		Email:          regisUser.Email,
		HashedPassword: regisUser.Password,
	}

	db, errDb := connection.GetConnectionGormApiFoodApp()
	if errDb != nil {
		return resp.FailedConnectDatabase("apifoodapp"), errors.New("failed to connect database: " + errDb.Error())
	}

	if errCreateUser := repository.NewUserRepository(db).CreateUser(&userData); errCreateUser != nil {
		return resp.SomethingWentWrong(), errCreateUser
	}

	return resp.Created("user created", userData), nil
}

func passwordEncrypt(userInfoRequest *req.UsersRequestAPI) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInfoRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userInfoRequest.Password = string(hashedPassword)

	return err
}
