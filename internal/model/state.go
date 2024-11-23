package model

import (
	"database/sql"
)

type StateModel struct {
	ID                  int32        `gorm:"primaryKey"`
	PlayersLastUpdate   sql.NullTime `gorm:"null"`
	TeamsLastUpdate     sql.NullTime `gorm:"null"`
	TournamentsLastFind sql.NullTime `gorm:"null"`
	MatchesLastFind     sql.NullTime `gorm:"null"`
}

func (StateModel) TableName() string {
	return "state.state"
}
