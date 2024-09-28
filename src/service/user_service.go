package service

import (
	"apifoodweb/config"
	"apifoodweb/config/database/model"
	repository "apifoodweb/config/database/repos"
	"apifoodweb/src/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(regisUser models.UsersRequestAPI) error {

	if regisUser.Password != regisUser.RePassword {
		return errors.New("password not match")
	}

	hashPassword, errHash := passwordEncrypt(regisUser.Password)
	if errHash != nil {
		return errors.New("failed to hash password")
	}

	userData := model.Users{
		Username:       regisUser.Username,
		Email:          regisUser.Email,
		HashedPassword: hashPassword,
	}

	db, errDb := config.GetConnectionGormApiFoodApp()
	if errDb != nil {
		return errors.New("failed to connect to database: " + errDb.Error())
	}

	return repository.NewUserRepository(db).CreateUser(&userData)
}

func passwordEncrypt(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPassword), err
}
