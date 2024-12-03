package model

import (
	"time"
)

type TournamentModel struct {
	Id              int32           `gorm:"primaryKey;autoIncrement"`
	FaceitId        string          `gorm:"unique;not null"`
	OrganizerId     string          `gorm:"not null"`
	Name            string          `gorm:"not null"`
	BackgroundImage string          `gorm:"null"`
	CoverImage      string          `gorm:"null"`
	Avatar          string          `gorm:"null"`
	RegisterDate    time.Time       `gorm:"not null"`
	StartDate       time.Time       `gorm:"not null"`
	Status          string          `gorm:"null"`
	JoinPolicy      string          `gorm:"not null"`
	GeoCountries    JSONStringArray `gorm:"type:jsonb;null"`
	MinLevel        int             `gorm:"not null"`
	MaxLevel        int             `gorm:"not null"`
	Type            string          `gorm:"not null"`
	TeamsId         JSONStringArray `gorm:"type:jsonb;null"`
}

func (TournamentModel) TableName() string {
	return "tournaments.tournament"
}
