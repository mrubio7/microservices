package model

import "github.com/google/uuid"

type UserModel struct {
	ID             int         `gorm:"primaryKey;autoIncrement"`
	FaceitId       string      `gorm:"unique;index"`
	Name           string      `gorm:"not null"`
	Description    string      `gorm:"null"`
	Twitter        string      `gorm:"null"`
	Twitch         string      `gorm:"null"`
	Role           int         `gorm:"null"`
	Player         PlayerModel `gorm:"-"`
	IsProfessional bool        `gorm:"not null; default:false"`
}

type UserSessionModel struct {
	UserID    int    `gorm:"primaryKey; not null"`
	SessionID string `gorm:"not null"`
}

func GenerateSessionId() string {
	return uuid.New().String()
}
