package model

import (
	"time"
)

type MatchModel struct {
	ID                 int             `gorm:"primaryKey; autoIncrement"`
	FaceitId           string          `gorm:"unique; not null"`
	TeamAFaceitId      string          `gorm:"not null"`
	TeamAName          string          `gorm:"not null"`
	TeamA              TeamModel       `gorm:"-"`
	IsTeamAKnown       bool            `gorm:"not null"`
	ScoreTeamA         int32           `gorm:"not null"`
	TeamBFaceitId      string          `gorm:"not null"`
	TeamB              TeamModel       `gorm:"-"`
	TeamBName          string          `gorm:"not null"`
	IsTeamBKnown       bool            `gorm:"not null"`
	ScoreTeamB         int32           `gorm:"not null"`
	BestOf             int32           `gorm:"not null"`
	Timestamp          time.Time       `gorm:"not null"`
	Streams            JSONStringArray `gorm:"type:jsonb; null"`
	TournamentFaceitId string          `gorm:"not null"`
	TournamentName     string          `gorm:"not null"`
	Tournament         TournamentModel `gorm:"-"`
	Map                JSONStringArray `gorm:"type:jsonb; not null"`
	Status             string          `gorm:"not null"`
}
