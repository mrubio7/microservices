package model

import "time"

type LookingForTeamModel struct {
	Id           int32           `gorm:"primaryKey; autoIncrement"`
	FaceitId     string          `gorm:"unique; not null"`
	InGameRole   JSONStringArray `gorm:"type:jsonb; not null"`
	TimeTable    string          `gorm:"not null"`
	OldTeams     string          `gorm:"not null"`
	PlayingYears int32           `gorm:"not null"`
	Location     string          `gorm:"null"`
	BornDate     time.Time       `gorm:"not null"`
	Description  string          `gorm:"not null"`
	CreatedAt    int64           `gorm:"not null"`
	UpdatedAt    int64           `gorm:"not null"`
	Player       PlayerModel     `gorm:"-"`
}
