package util

import (
	"apifoodweb/api/dto/req"

	"golang.org/x/crypto/bcrypt"
)

func PasswordEncrypt(userInfoRequest *req.UsersRequestAPI) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInfoRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userInfoRequest.Password = string(hashedPassword)

	return err
}
