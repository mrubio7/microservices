package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type EseaLeagueModel struct {
	Id           int32               `gorm:"primaryKey;autoIncrement"`
	FaceitId     string              `gorm:"unique;not null"`
	Name         string              `gorm:"not null"`
	Season       int32               `gorm:"not null"`
	Playoffs     bool                `gorm:"not null"`
	PlayoffsData JSONString          `gorm:"not null"`
	Divisions    []EseaDivisionModel `gorm:"foreignKey:EseaLeagueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (EseaLeagueModel) TableName() string {
	return "tournaments.esea_league"
}

type JSONString string

func (j JSONString) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONString) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONString value: %v", value)
	}
	return json.Unmarshal(bytes, j)
}

type EseaDivisionModel struct {
	Id                 int32               `gorm:"primaryKey;autoIncrement"`
	EseaLeagueId       int32               `gorm:"not null;index"`
	EseaLeagueFaceitId string              `gorm:"not null"`
	FaceitId           string              `gorm:"not null"`
	TeamsId            JSONStringArray     `gorm:"type:jsonb;null"`
	Name               string              `gorm:"not null"`
	Standings          []EseaStandingModel `gorm:"foreignKey:DivisionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (EseaDivisionModel) TableName() string {
	return "tournaments.esea_division"
}

type EseaStandingModel struct {
	Id             int32     `gorm:"primaryKey; autoIncrement"`
	DivisionId     string    `gorm:"not null;index"`
	TeamFaceitId   string    `gorm:"not null"`
	IsDisqualified bool      `gorm:"not null"`
	TournamentName string    `gorm:"not null"`
	RankStart      int       `gorm:"not null"`
	RankEnd        int       `gorm:"not null"`
	Points         int       `gorm:"not null"`
	MatchesPlayed  int       `gorm:"not null"`
	MatchesWon     int       `gorm:"not null"`
	MatchesLost    int       `gorm:"not null"`
	MatchesTied    int       `gorm:"not null"`
	BuchholzScore  int       `gorm:"not null"`
	Team           TeamModel `gorm:"-"`
}

func (EseaStandingModel) TableName() string {
	return "tournaments.esea_standing"
}
