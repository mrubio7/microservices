package service

import (
	"sync"

	"ibercs/internal/model"
	"ibercs/pkg/logger"

	"gorm.io/gorm"
)

type Users struct {
	db    *gorm.DB
	mutex sync.Mutex
}

func NewUsersService(database *gorm.DB) *Users {
	return &Users{
		db: database,
	}
}

func (svc *Users) GetUserById(id string) *model.UserModel {
	var user model.UserModel
	err := svc.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	return &user
}

func (svc *Users) GetUserByFaceitId(id string) *model.UserModel {
	var user model.UserModel
	err := svc.db.Where("faceit_id = ?", id).First(&user).Error
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	return &user
}

func (svc *Users) UpdateUser(user model.UserModel) *model.UserModel {
	err := svc.db.Save(&user).Error
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	return &user
}

func (svc *Users) NewUser(user model.UserModel) *model.UserModel {
	err := svc.db.Create(&user).Error
	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	return &user
}
