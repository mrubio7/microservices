package model

import (
	"time"
)

type TournamentModel struct {
	ID              uint32          `gorm:"primaryKey; autoIncrement"`
	FaceitId        string          `gorm:"unique; not null"`
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
	CurrentTeams    int             `gorm:"-"`
	Slots           int             `gorm:"-"`
}

type EseaDivisionModel struct {
	ID           uint32              `gorm:"primaryKey; autoIncrement"`
	FaceitId     string              `gorm:"unique; not null"`
	TournamentId string              `gorm:"not null"`
	ConferenceId string              `gorm:"not null"`
	TeamsId      JSONStringArray     `gorm:"type:jsonb;null"`
	Name         string              `gorm:"not null"`
	Standings    []EseaStandingModel `gorm:"-"`
}

type EseaStandingModel struct {
	ID             uint32 `gorm:"primaryKey; autoIncrement"`
	FaceitId       string `gorm:"not null;uniqueIndex:idx_faceit_tournament"`
	TournamentId   string `gorm:"not null;uniqueIndex:idx_faceit_tournament"`
	IsDisqualified bool   `gorm:"not null"`
	TournamentName string `gorm:"not null"`
	RankStart      int    `gorm:"not null"`
	RankEnd        int    `gorm:"not null"`
	Points         int    `gorm:"not null"`
	MatchesPlayed  int    `gorm:"not null"`
	MatchesWon     int    `gorm:"not null"`
	MatchesLost    int    `gorm:"not null"`
	MatchesTied    int    `gorm:"not null"`
	BuchholzScore  int    `gorm:"not null"`
}
