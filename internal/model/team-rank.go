package model

type TeamRankModel struct {
	Id           int32   `gorm:"primaryKey;autoIncrement"`
	FaceitId     string  `gorm:"unique;index"`
	ActualPoints float32 `gorm:"not null"`
	OldPoints    float32 `gorm:"not null"`
	Matches      int     `gorm:"not null"`
	LeaguePoints float32 `gorm:"null"`
}

func (TeamRankModel) TableName() string {
	return "teams.team_rank"
}
