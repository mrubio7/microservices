package service

import (
	"sync"

	"ibercs/internal/model"
	"ibercs/pkg/logger"

	"github.com/google/uuid"
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

func (svc *Users) GetUserByPlayerNickname(nickname string) *model.UserModel {
	var player model.PlayerModel
	if err := svc.db.Preload("Stats").Where("nickname = ?", nickname).First(&player).Error; err != nil {
		logger.Error("Player not found: %v", err)
		return nil
	}

	var user model.UserModel
	if err := svc.db.Where("faceit_id = ?", player.FaceitId).First(&user).Error; err != nil {
		logger.Error("User not found: %v", err)
		return nil
	}

	user.Player = player

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

func (svc *Users) NewSession(id int) string {
	session := model.UserSessionModel{
		UserID:    id,
		SessionID: uuid.New().String(),
	}

	err := svc.db.Save(&session).Error
	if err != nil {
		logger.Error(err.Error())
		return ""
	}

	return session.SessionID
}

func (svc *Users) DeleteSession(id int) string {
	var session model.UserSessionModel
	err := svc.db.Where("user_id = ?", id).First(&session).Error
	if err != nil {
		logger.Error(err.Error())
		return ""
	}

	err = svc.db.Delete(&session).Error
	if err != nil {
		logger.Error(err.Error())
		return ""
	}

	return session.SessionID
}

func (svc *Users) GetAllStreams() []string {
	var streams []string
	err := svc.db.Model(&model.UserModel{}).Where("twitch != ''").Pluck("twitch", &streams).Error
	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	return streams
}
