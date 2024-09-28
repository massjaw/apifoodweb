package repository

import (
	"apifoodweb/config/database/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUser() ([]model.Users, error)
	FindUserByID(id string) (model.Users, error)
	FindUserDetailByID(id string) (model.UserDetail, error)
	CreateUser(user *model.Users) error
	UpdateUser(user *model.Users) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAllUser() ([]model.Users, error) {

	var users []model.Users

	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) FindUserByID(userId string) (user model.Users, errMsg error) {

	errMsg = r.db.First(&user, "id = ?", userId).Error
	if errMsg != nil {
		return user, errMsg
	}

	return user, errMsg
}

func (r *userRepository) FindUserDetailByID(userId string) (user model.UserDetail, errMsg error) {

	errMsg = r.db.First(&user, "user_id = ?", userId).Error
	if errMsg != nil {
		return user, errMsg
	}

	return user, errMsg
}

func (r *userRepository) CreateUser(user *model.Users) error {

	tx := r.db.Begin()
	defer tx.Rollback()

	errUser := tx.Create(&user).Error
	errDetailUser := tx.Create(&model.UserDetail{UserID: user.ID}).Error

	if errUser != nil || errDetailUser != nil {
		return errors.New(errUser.Error() + errDetailUser.Error())
	}

	tx.Commit()
	return nil
}

func (r *userRepository) UpdateUser(user *model.Users) error {

	return r.db.Model(&model.Users{}).Where("user = ?", user.ID).Updates(user).Error
}

func (r *userRepository) DeleteUser(id string) error {

	return r.db.Delete(&model.Users{}, "id=?", id).Error
}
